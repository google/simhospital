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

// Package logging provides functionality for logging.
package logging

import (
	"runtime"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const basePackagePrefix = "/simulated_hospital/"

var shBaseLogger = logrus.New()

func init() {
	shBaseLogger.Formatter = WithSource(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
		ForceColors:     true,
	})
}

// SimulatedHospitalLogger embeds the Logrus logger.
type SimulatedHospitalLogger struct {
	*logrus.Entry
}

// WithError adds an error as single field (using the key defined in ErrorKey) to the Entry.
// Overridden from logrus.Entry so that we can return a *SimulatedHospitalLogger.
func (s *SimulatedHospitalLogger) WithError(err error) *SimulatedHospitalLogger {
	return &SimulatedHospitalLogger{
		Entry: s.Entry.WithError(err),
	}
}

// WithField adds a single field to the Entry.
// Overridden from logrus.Entry so that we can return a *SimulatedHospitalLogger.
func (s *SimulatedHospitalLogger) WithField(key string, value interface{}) *SimulatedHospitalLogger {
	return &SimulatedHospitalLogger{
		Entry: s.Entry.WithField(key, value),
	}
}

// ForCallerPackage builds a base Entry that prefixes the name of the package from which it was
// called and can be used instead of the standard logger.
func ForCallerPackage() *SimulatedHospitalLogger {
	_, filename, _, _ := runtime.Caller(1)
	return &SimulatedHospitalLogger{
		Entry: shBaseLogger.WithField("prefix", packageName(filename)),
	}
}

// packageName returns the name of the package to which the specified file belongs, stripping the
// basePackagePrefix, as it is common to all the packages for which this method is used in
// practice.
// If the filename is not well formed, or for some reason does not live inside the
// basePackagePrefix folder, it will return the full base path of the file instead.
func packageName(filename string) string {
	start := strings.Index(filename, basePackagePrefix)
	if start == -1 {
		start = 0
	} else {
		start += len(basePackagePrefix)
	}
	end := strings.LastIndex(filename, "/")
	if end < start {
		end = start
	}
	return filename[start:end]
}

// SetLogLevelFromString sets the log level.
// SetLogLevelFromString returns an error if the given log level is not recognized by logrus.
func SetLogLevelFromString(level string) error {
	l, err := logrus.ParseLevel(level)
	if err != nil {
		return errors.Wrapf(err, "cannot parse log level %v", level)
	}
	SetLogLevel(l)
	return nil
}

// SetLogLevel sets the log level.
func SetLogLevel(level logrus.Level) {
	shBaseLogger.Level = level
}
