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

// logging_source_test is a test-only package for the logging package, to test
// that log entries contain the proper source of the call. It is required
// because logging functions compute the source as being the first call frame
// not within the same package.
package logging_source_test

import (
	"bytes"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/google/simhospital/pkg/logging"
)

func funcName(v reflect.Value) string {
	name := runtime.FuncForPC(v.Pointer()).Name()
	return name[strings.LastIndex(name, ".")+1:]
}

func TestLoggingSource(t *testing.T) {
	logging.SetLogLevel(logrus.DebugLevel)
	l := logging.ForCallerPackage()
	logger := l.Entry.Logger

	const want = `logging_source_test.TestLoggingSource\(.*pkg/logging/logging_source_test/logging_source_test.go:\d+`
	wantRE, err := regexp.Compile(want)
	if err != nil {
		t.Fatalf("regexp.Compile(%s) failed with %v", want, err)
	}

	// Test non-formatting functions.
	for _, fn := range []func(...interface{}){
		l.Debug,
		l.Info,
		l.Warning,
		l.Error,
	} {
		buf := new(bytes.Buffer)
		logger.Out = buf
		const msg = "foo"
		fn(msg)

		fnName := funcName(reflect.ValueOf(fn))
		got := buf.String()
		if want := msg; !strings.Contains(got, want) {
			t.Errorf("%s(%q) got %s; want containing logged message %q", fnName, msg, got, want)
		}
		if !wantRE.MatchString(got) {
			t.Errorf("%s() got %s; want matching %s", fnName, got, want)
		}
	}
}

func TestLoggingSourceFormatted(t *testing.T) {
	logging.SetLogLevel(logrus.DebugLevel)
	l := logging.ForCallerPackage()
	logger := l.Entry.Logger

	const want = `logging_source_test.TestLoggingSourceFormatted\(pkg/logging/logging_source_test/logging_source_test.go:\d+`
	wantRE, err := regexp.Compile(want)
	if err != nil {
		t.Fatalf("regexp.Compile(%s) failed with %v", want, err)
	}

	// Test formatting functions.
	for _, fn := range []func(string, ...interface{}){
		l.Debugf,
		l.Infof,
		l.Warningf,
		l.Errorf,
	} {
		buf := new(bytes.Buffer)
		logger.Out = buf
		const (
			format  = "%s %d"
			msg     = "foo"
			answer  = 42
			wantMsg = "foo 42"
		)
		fn(format, msg, answer)

		fnName := funcName(reflect.ValueOf(fn))
		got := buf.String()
		if want := wantMsg; !strings.Contains(got, want) {
			t.Errorf("%s(%q, %q, %d) got %s; want containing logged message %q", fnName, format, msg, answer, got, want)
		}
		if !wantRE.MatchString(got) {
			t.Errorf("%s() got %s; want matching %s", fnName, got, want)
		}
	}
}
