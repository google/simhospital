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

// Package rate contains functionality to deal with rates.
package rate

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/simhospital/pkg/logging"
)

var log = logging.ForCallerPackage()

// Controller is a rate controller.
type Controller struct {
	rate        float64
	per         time.Duration
	rateChanged chan bool
}

// NewController creates a new Controller.
func NewController(rate float64, per time.Duration) *Controller {
	return &Controller{
		rate:        rate,
		per:         per,
		rateChanged: make(chan bool),
	}
}

// ServeHTTP handles the requests made from the slider on the control dashboard.
func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s := strconv.FormatFloat(c.rate, 'f', -1, 64)
		w.Write([]byte(s))
	case "POST":
		c.handlePost(w, r)
	case "PUT", "DELETE":
		http.Error(w, fmt.Sprintf("Method %q not implemented", r.Method), http.StatusInternalServerError)
	default:
		http.Error(w, fmt.Sprintf("Unknown method: %q", r.Method), http.StatusInternalServerError)
	}
}

func (c *Controller) handlePost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	errStr := "Failed to change the rate value"
	if err != nil {
		log.WithError(err).Warning(errStr)
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	sbody := string(body)
	prefix := "value="
	if !strings.HasPrefix(sbody, prefix) {
		log.Warningf("%s: missing prefix %q in request body: %q", errStr, prefix, sbody)
		http.Error(w, `Error extracting value: the request must be in the format "value=X"`, http.StatusInternalServerError)
		return
	}
	svalue := strings.TrimPrefix(sbody, prefix)
	f, err := strconv.ParseFloat(svalue, 64)
	if err != nil {
		log.WithError(err).Warning(errStr)
		http.Error(w, "Error parsing value to float", http.StatusInternalServerError)
		return
	}
	if f < 0 {
		msg := fmt.Sprintf("Invalid value: rate per minute must be greater that zero, but was: %v", f)
		log.Warningf("%s: %s", errStr, msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	// Sometimes we get a POST message even if the value didn't change.
	if c.rate != f {
		c.rate = f
		c.rateChanged <- true
	}
}

func (c *Controller) heartbeat() time.Duration {
	return time.Duration(float64(c.per) / c.rate)
}

// Heartbeat returns a duration between two pathways based on rate and per values.
// If the rate is set to zero, returns the maximum duration value.
func (c *Controller) Heartbeat() time.Duration {
	if c.rate == 0 {
		log.Infof("Rate set to %v / %v. Not generating pathway", c.rate, c.per)
		// Max Duration value.
		return time.Duration(1<<63 - 1)
	}
	h := c.heartbeat()
	log.Debugf("Rate set to %v / %v. Generating one pathway every %v", c.rate, c.per, h)
	return h
}

// RateChanged returns a channel, where the changes of the rate are signaled.
func (c *Controller) RateChanged() <-chan bool {
	return c.rateChanged
}

// InitialElapsed returns a value of the heartbeat, if the rate is not zero.
// Otherwise, returns zero.
func (c *Controller) InitialElapsed() time.Duration {
	if c.rate > 0 {
		return c.heartbeat()
	}
	return 0
}
