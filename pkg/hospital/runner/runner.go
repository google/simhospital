// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package runner implements the main functionality of Simulated Hospital.
package runner

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"github.com/gorilla/mux"
	"github.com/google/simhospital/pkg/clock"
	"github.com/google/simhospital/pkg/hospital"
	"github.com/google/simhospital/pkg/hospital/runner/authentication"
	"github.com/google/simhospital/pkg/logging"
	"github.com/google/simhospital/pkg/monitoring"
	"github.com/google/simhospital/pkg/rate"
	"github.com/google/simhospital/pkg/starter"
)

var (
	log = logging.ForCallerPackage()

	// dashboardAddressRegex is the regex to match dashboard address.
	dashboardAddressRegex = regexp.MustCompile(`^:\d{4}$`)
)

// EndpointAndHandler defines a Simulated Hospital endpoint and its handler.
type EndpointAndHandler struct {
	Endpoint string
	Handler  func(http.ResponseWriter, *http.Request)
}

// APIEndpointAndHandler defines a Simulated Hospital API endpoint and its handler.
type APIEndpointAndHandler struct {
	EndpointAndHandler
	HTTPMethod string
}

// Hospital wraps the hospital.Hospital and implements the run functionality.
type Hospital struct {
	hospital                     *hospital.Hospital
	pathwayRateController        *rate.Controller
	pathwayStarter               *starter.PathwayStarter
	additionalDashboardEndpoints []EndpointAndHandler
	authenticatedEndpoints       []APIEndpointAndHandler
	authenticatedAPIConfig       APIConfig
	dashboardURI                 string
	dashboardAddress             string
	dashboardStaticDir           string
	metricsAddress               string
	sleepFor                     time.Duration
	clock                        clock.Clock
	maxPathways                  int
	creatingPathways             chan bool
	processingEvents             chan bool
	processingMessages           chan bool
}

// APIConfig contains base configuration for authenticated endpoints.
type APIConfig struct {
	APIPort string
	APIKey  string
}

// Config contains optional configuration options for Simulated Hospital Runner
// used to extend the main functionality.
type Config struct {
	// AdditionalDashboardEndpoints is a slice of endpoints and their handlers.
	// The root path for these endpoints will be the Simulated Hospital dashboard address.
	AdditionalDashboardEndpoints []EndpointAndHandler
	// AuthenticatedAPIConfig is the API config for authenticated endpoints.
	AuthenticatedAPIConfig APIConfig
	// AuthenticatedEndpoints is a slice of API endpoints and their handlers.
	// The root path for these endpoints will be the API root path.
	AuthenticatedEndpoints []APIEndpointAndHandler
	// PathwayStarter is a starter of pathways through an endpoint.
	PathwayStarter *starter.PathwayStarter
	// PathwaysPerHour indicates how often new pathways are generated.
	PathwaysPerHour float64
	// MaxPathways is the number of pathways to run before stopping.
	// If negative, Simulated Hospital will keep running pathways indefinitely.
	MaxPathways int
	// DashboardURI is the base URI path at which the simulated hospital dashboard and
	// endpoints are available, e.g., /simulated-hospital/. Note that this needs to match
	// the data-path on elements in index.html.
	DashboardURI string
	// DashboardAddress is the port on which the simulated hospital is accessible.
	// The value is expected to be in the form :int, e.g. :8000.
	DashboardAddress string
	// DashboardStaticDir is the directory for static assets for the dashboard.
	DashboardStaticDir string
	// MetricsAddress is the address for the /metrics endpoint.
	MetricsAddress string
	// SleepFor represents the interval at which the queues are checked.
	SleepFor time.Duration
	// Clock is the clock for the hospital.
	Clock clock.Clock
}

func (c Config) isValid() error {
	if c.DashboardURI == "" {
		return errors.New("must provide a base URI at which the dashboard is available")
	}
	if c.DashboardAddress == "" || !dashboardAddressRegex.MatchString(c.DashboardAddress) {
		return errors.New("must provide a valid dashboard address/port on which to start the simulated hospital")
	}
	if c.DashboardStaticDir == "" {
		return errors.New("must provide a valid directory path for serving static assets")
	}
	if len(c.AuthenticatedEndpoints) != 0 && (c.AuthenticatedAPIConfig.APIKey == "" || c.AuthenticatedAPIConfig.APIPort == "") {
		return errors.New("must provide API key and port if API endpoints are configured")
	}
	return nil
}

// New creates a new Runner.
func New(h *hospital.Hospital, config Config) (*Hospital, error) {
	if err := config.isValid(); err != nil {
		return nil, err
	}

	rand.Seed(time.Now().Unix())
	return &Hospital{
		hospital:                     h,
		pathwayRateController:        rate.NewController(config.PathwaysPerHour, time.Hour),
		pathwayStarter:               config.PathwayStarter,
		additionalDashboardEndpoints: config.AdditionalDashboardEndpoints,
		authenticatedEndpoints:       config.AuthenticatedEndpoints,
		authenticatedAPIConfig:       config.AuthenticatedAPIConfig,
		dashboardURI:                 config.DashboardURI,
		dashboardAddress:             config.DashboardAddress,
		dashboardStaticDir:           config.DashboardStaticDir,
		metricsAddress:               config.MetricsAddress,
		sleepFor:                     config.SleepFor,
		clock:                        config.Clock,
		maxPathways:                  config.MaxPathways,
	}, nil
}

// Run starts the Simulated Hospital.
// It starts multiple servers including the Simulated Hospital dashboard.
// The following happens in parallel and continuously while Simulated Hospital is running:
// 1. Start pathways, which create events (e.g., patients are admitted in the
//    hospital, test results, etc.).
// 2. Run those events at the appropriate time, which generates HL7 messages.
// 3. Process HL7 messages at the appropriate time.
// The processing of pathways, events and messages can finish if Hospital.maxPathways is negative, otherwise Run() runs forever.
// When the creation of pathways finishes, we can stop processing events after all our current events
// are processed. When processing events finishes, we can stop processing messages after all our current
// messages are processed.
// If this happens, all servers are stopped and this method returns.
func (h *Hospital) Run(ctx context.Context) {
	ctxCancel, cancel := context.WithCancel(ctx)
	defer cancel()
	// The groupCtx context is cancelled when:
	// - its parent context is cancelled, or
	// - any of the functions started in the error group eg terminates with an error.
	eg, groupCtx := errgroup.WithContext(ctxCancel)

	logLocal := log.WithContext(groupCtx)

	eg.Go(func() error {
		logLocal.Infof("Starting metrics server on address %v", h.metricsAddress)
		if err := monitoring.ListenAndServeMetrics(groupCtx, h.metricsAddress); err != nil {
			return errors.Wrapf(err, "Failed to run metrics server on address %s", h.metricsAddress)
		}
		return nil
	})

	dashboard := createHTTPServer(groupCtx, h.dashboardAddress, h.setupEndpoints())
	logLocal.Infof("Starting dashboard on address %s", dashboard.Addr)
	eg.Go(func() error {
		if err := dashboard.ListenAndServe(); err != nil {
			return errors.Wrapf(err, "Failed to run dashboard for Simulated Hospital on address %v", dashboard.Addr)
		}
		return nil
	})

	if m := h.setupAuthenticatedEndpoints(); m != nil {
		auth := createHTTPServer(groupCtx, h.authenticatedAPIConfig.APIPort, m)
		logLocal.Infof("Starting authenticated endpoints on address %s", h.authenticatedAPIConfig.APIPort)
		eg.Go(func() error {
			if err := auth.ListenAndServe(); err != nil {
				return errors.Wrapf(err, "Failed to run authenticated endpoints on address %s", auth.Addr)
			}
			return nil
		})
	}

	if h.maxPathways >= 0 {
		// We use the creatingPathways, processingEvents and processingMessages channels
		// to communicate whether pathways, events and messages are still being processed.
		h.creatingPathways = make(chan bool)
		h.processingEvents = make(chan bool)
		h.processingMessages = make(chan bool)
		go func() {
			processing := true
			for processing {
				select {
				case processing = <-h.processingMessages:
				}
			}
			// Cancelling the context exits the Run() method.
			defer cancel()
			logLocal.Info("Simulated Hospital stopped processing, will exit")
		}()
	}

	// 1. Start the pathways that create the events.
	eg.Go(func() error {
		return h.startPathways(groupCtx)
	})

	// 2. Run the events.
	eg.Go(func() error {
		return h.RunEvents(groupCtx)
	})

	// 3. Process (e.g. send) the messages.
	eg.Go(func() error {
		return h.ProcessMessages(groupCtx)
	})

	if err := eg.Wait(); err != nil {
		logLocal.WithError(err).Error("Simulated Hospital exited with errors")
		return
	}
	logLocal.Info("Simulated Hospital exited")
}

// Close closes resources held by the Hospital.
// Should be called if the Hospital is no longer needed or at the program exit.
func (h *Hospital) Close() error {
	if err := h.hospital.Close(); err != nil {
		return errors.Wrap(err, "error closing hospital")
	}
	return nil
}

// createHTTPServer creates an http server on the specified port using the handler provided.
// It also starts a goroutine that closes the server as soon as the context provided is done.
func createHTTPServer(ctx context.Context, addr string, handler http.Handler) *http.Server {
	srv := &http.Server{Addr: addr, Handler: handler}
	go func() {
		<-ctx.Done()
		log.WithContext(ctx).Infof("Closing server on address %s", addr)
		srv.Close()
	}()
	return srv
}

// startPathways starts pathways that create events.
// If h.maxPathways is negative, it runs indefinitely. Otherwise, it stops when
// the number of created pathways is equal to h.maxPathways.
// When startPathways stops, it writes "false" in the h.creatingPathways channel and closes the channel.
// The delay between running consecutive pathways is derived by the rate Controller,
// based on the rate.
// If the rate is initially set to value != 0, then the first pathway
// is started immediately.
// If the rate changes, startPathways respects the new value immediately,
// by taking into account the time that has already elapsed since
// the last pathway run.
// Eg:
// - the rate changed from 1 pathway / hour to 4 pathway / hour,
//   and last pathway was started 10 mins ago -> the next pathway will start
//   in 5 mins, as the new Heartbeat value is now 15 mins.
// - the rate changed from 4 pathway / hour to 1 pathway / hour,
//   and last pathway was started 10 mins ago -> the next pathway will start
//   in 50 mins, as the new Heartbeat value is now 1h.
// - the rate was initially set to 0 pathway / hour (so no pathway was started initially)
//   and was changed to 1 pathway / hour -> the next pathway will start after 1h elapses
//   since the beginning of SH running.
//
// Returns an error if the context is Done.
func (h *Hospital) startPathways(ctx context.Context) error {
	elapsed := h.pathwayRateController.InitialElapsed()
	nCreated := 0
	logLocal := log.WithContext(ctx)
	if h.maxPathways >= 0 {
		logLocal.Infof("Number of pathways to run: %d (excluding pathways run from the Dashboard)", h.maxPathways)
	}

	for h.maxPathways < 0 || nCreated < h.maxPathways {
		start := h.clock.Now()
		delay := h.pathwayRateController.Heartbeat() - elapsed
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-h.pathwayRateController.RateChanged():
			// The rate was changed; we might need to generate a new pathway sooner.
			elapsed += h.clock.Now().Sub(start)
			continue
		case <-time.After(delay):
			elapsed = time.Duration(0)
			nCreated++
			if err := h.hospital.StartNextPathway(); err != nil {
				logLocal.WithError(err).Error("cannot start new pathway")
			}
		}
	}
	if h.creatingPathways != nil {
		h.creatingPathways <- false
		close(h.creatingPathways)
	}
	logLocal.Info("Pathway generation finished")
	return nil
}

// RunEvents runs the events as they are due.
// Returns an error if the context is Done.
func (h *Hospital) RunEvents(ctx context.Context) error {
	err := h.processItems(ctx, h.hospital.RunNextEventIfDue, h.hospital.HasEvents, h.creatingPathways, h.processingEvents, "Failed to run the due event")
	if err != nil {
		return err
	}
	log.WithContext(ctx).Info("Event processing finished")
	return nil
}

// ProcessMessages processes (e.g. sends) the HL7 messages.
// Returns an error if the context is Done.
func (h *Hospital) ProcessMessages(ctx context.Context) error {
	f := func(_ context.Context) (bool, error) {
		return h.hospital.ProcessNextMessageIfDue()
	}
	err := h.processItems(ctx, f, h.hospital.HasMessages, h.processingEvents, h.processingMessages, "Failed to process the due message")
	if err != nil {
		return err
	}
	log.WithContext(ctx).Info("Message processing finished")
	return nil
}

// processItems runs the function f while there are items to be processed.
// There are two ways we can have items to be processed:
// (a) there are actual items waiting for processing, which is captured by hasItems, or
// (b) more items can still be created, which is captured by the creatingItems channel.
// Items of the current type can be created externally, for instance, pathways create events and events create messages.
// creatingItems is a channel that signals whether items are still being created outside this method.
// After all items are processed, processItems writes "false" in the processingItems channel and closes the channel.
// Send a nil creatingItems channel to run indefinitely.
func (h *Hospital) processItems(ctx context.Context, f func(context.Context) (bool, error), hasItems func() bool, creatingItems <-chan bool, processingItems chan bool, errMsg string) error {
	stillCreating := true
	for stillCreating || hasItems() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(h.sleepFor):
			// Process everything that is due now.
			for {
				ran, err := f(ctx)
				if err != nil {
					log.WithContext(ctx).WithError(err).Error(errMsg)
				}
				if !ran {
					break
				}
			}
		}
		if creatingItems != nil {
			select {
			case stillCreating = <-creatingItems:
			}
		}
	}
	if processingItems != nil {
		processingItems <- false
		close(processingItems)
	}
	return nil
}

// setupEndpoints sets up the regular endpoints (pathway rate and pathway starter)
// plus any additional endpoints in additionalDashboardEndpoints, and returns the http.ServeMux.
// This method always returns a non-nil item.
func (h *Hospital) setupEndpoints() *http.ServeMux {
	m := http.NewServeMux()
	m.Handle(fmt.Sprintf("/%s/", h.dashboardURI), http.StripPrefix(fmt.Sprintf("/%s/", h.dashboardURI), http.FileServer(http.Dir(h.dashboardStaticDir))))
	endpoints := append([]EndpointAndHandler{
		{Endpoint: "pathwayRate", Handler: h.pathwayRateController.ServeHTTP},
		{Endpoint: "pathwayStarter", Handler: h.pathwayStarter.ServeHTTP},
	}, h.additionalDashboardEndpoints...)
	for _, e := range endpoints {
		log.WithField("root_path", h.dashboardURI).WithField("endpoint", e.Endpoint).Info("Setting up endpoint")
		m.HandleFunc(fmt.Sprintf("/%s/%s", h.dashboardURI, e.Endpoint), e.Handler)
	}
	return m
}

// setupAuthenticatedEndpoints sets up the authenticated endpoints and returns the mux.Router.
// If there are no authenticated endpoints, this method returns nil.
func (h *Hospital) setupAuthenticatedEndpoints() *mux.Router {
	if len(h.authenticatedEndpoints) == 0 {
		log.Info("No authenticated endpoints to set up")
		return nil
	}
	r := mux.NewRouter()
	for _, e := range h.authenticatedEndpoints {
		log.WithField("root_path", h.apiRootPath()).
			WithField("endpoint", e.Endpoint).
			WithField("key_is_set", h.authenticatedAPIConfig.APIKey != "").
			WithField("http_method", e.HTTPMethod).
			Info("Setting up authenticated endpoint")
		r.HandleFunc(fmt.Sprintf("/%s/%s", h.apiRootPath(), e.Endpoint), authentication.Middleware(e.Handler, h.authenticatedAPIConfig.APIKey)).Methods(e.HTTPMethod)
	}
	return r
}

// apiRootPath returns the root path at which endpoints are available by concatenating
// the dashboard URI with '/api'.
func (h *Hospital) apiRootPath() string {
	return fmt.Sprintf("%s/%s", h.dashboardURI, "api")
}
