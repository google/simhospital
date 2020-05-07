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

// Package monitoring contains functionality for metrics.
package monitoring

import (
	"context"
	"net"
	"net/http"
	"net/http/pprof"

	"github.com/pkg/errors"
	"golang.org/x/net/trace"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// ListenAndServeMetrics initialises a new HTTP Server in the address specified in the parameter.
// This function blocks indefinitely and should be run on a separate goroutine.
// If the context provided is done, the HTTP Server is closed.
func ListenAndServeMetrics(ctx context.Context, address string) error {
	lis, err := startTCPListener(address)
	if err != nil {
		return errors.Wrapf(err, "cannot start TCP listener on address %s", address)
	}
	return serveMetrics(ctx, lis)
}

// serveMetrics initialises a new HTTP Server using the provided TCP listener with handlers for
// /metrics, /debug/requests, /debug/events and /debug/pprof.
func serveMetrics(ctx context.Context, lis net.Listener) error {
	s := http.NewServeMux()
	s.Handle("/metrics", promhttp.Handler())

	// Endpoints displaying RPC histograms and traces.
	s.HandleFunc("/debug/requests", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		trace.Render(w, req, false)
	})
	s.HandleFunc("/debug/events", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		trace.RenderEvents(w, req, false)
	})
	s.HandleFunc("/debug/pprof/", pprof.Index)
	s.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	s.HandleFunc("/debug/pprof/profile", pprof.Profile)
	s.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	s.HandleFunc("/debug/pprof/trace", pprof.Trace)

	logLocal := log.WithContext(ctx).WithField("metrics_listen_address", lis.Addr())
	logLocal.Info("Starting /metrics HTTP server")
	srv := &http.Server{Handler: s}
	go func() {
		<-ctx.Done()
		logLocal.Info("Closing metrics server")
		srv.Close()
	}()

	if err := srv.Serve(lis); err != nil {
		return errors.Wrap(err, "cannot serve /metrics on HTTP server")
	}
	return nil
}

// startTCPListener creates and returns TCP listener on the specified address.
// Note: ":0" does not appear to work well, looks to be related to https://github.com/golang/go/issues/22811
func startTCPListener(address string) (net.Listener, error) {
	log.WithField("listen_address", address).Debug("Starting TCP listener")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot listen on address %s", address)
	}
	log.WithField("listen_address", lis.Addr()).Info("Started TCP listener")
	return lis, nil
}
