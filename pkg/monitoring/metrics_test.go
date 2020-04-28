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

package monitoring

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestToUnderscore(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "foo", want: "foo"},
		{input: "Foo", want: "foo"},
		{input: "FOO", want: "foo"},
		{input: "AuthFoo", want: "auth_foo"},
		{input: "authFoo", want: "auth_foo"},
		{input: "OAuthFoo", want: "oauth_foo"},
		{input: "myGRPCValues", want: "my_grpcvalues"},
		{input: "myGrpcValues", want: "my_grpc_values"},
		{input: "IOSLogging", want: "ioslogging"},
		{input: "GetMRNError", want: "get_mrn_error"},
		{input: "PatientRetrievalMRNError", want: "patient_retrieval_mrn_error"},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if got, want := toUnderscore(tc.input), tc.want; got != want {
				t.Errorf("toUnderscore(%v) got %v, want %v", tc.input, got, want)
			}
		})
	}
}

func TestCreateAndRegisterMetricsFromStruct_Counter(t *testing.T) {
	wantMetric := "group_first_counter"
	m := metricsFromGatherer(t)
	if m[wantMetric] {
		t.Fatalf("prometheus.DefaultGatherer.Gather() got metrics %v, want metric %q not to be included", m, wantMetric)
	}
	var testMetrics struct {
		Group struct {
			FirstCounter prometheus.Counter `help:"Help Message"`
		}
	}
	if err := CreateAndRegisterMetricsFromStruct(&testMetrics); err != nil {
		t.Fatalf("CreateAndRegisterMetricsFromStruct(%v) failed with %v", &testMetrics, err)
	}
	defer prometheus.Unregister(testMetrics.Group.FirstCounter)
	testMetrics.Group.FirstCounter.Inc()
	m = metricsFromGatherer(t)
	if !m[wantMetric] {
		t.Errorf("prometheus.DefaultGatherer.Gather() got metrics %v, want metric %q to be included", m, wantMetric)
	}
}

func TestCreateAndRegisterMetricsFromStruct_CounterVec(t *testing.T) {
	wantMetric := "group_first_vec"
	m := metricsFromGatherer(t)
	if m[wantMetric] {
		t.Fatalf("prometheus.DefaultGatherer.Gather() got metrics %v, want metric %q not to be included", m, wantMetric)
	}
	var testMetrics struct {
		Group struct {
			FirstVec *prometheus.CounterVec `help:"Help Message" labels:"label1,label2"`
		}
	}
	if err := CreateAndRegisterMetricsFromStruct(&testMetrics); err != nil {
		t.Fatalf("CreateAndRegisterMetricsFromStruct(%v) failed with %v", &testMetrics, err)
	}
	defer prometheus.Unregister(testMetrics.Group.FirstVec)
	testMetrics.Group.FirstVec.With(prometheus.Labels{
		"label1": "value1",
		"label2": "value2",
	}).Inc()
	m = metricsFromGatherer(t)
	if !m[wantMetric] {
		t.Errorf("prometheus.DefaultGatherer.Gather() got metrics %v, want metric %q to be included", m, wantMetric)
	}
}

func TestCreateAndRegisterMetricsFromStruct_CounterMultiple(t *testing.T) {
	wantMetric1, wantMetric2 := "group_first_counter", "group_second_counter"
	m := metricsFromGatherer(t)
	if m[wantMetric1] || m[wantMetric2] {
		t.Fatalf("prometheus.DefaultGatherer.Gather() got metrics %v, want metrics %q and %q not included", m, wantMetric1, wantMetric2)
	}
	var testMetrics struct {
		Group struct {
			FirstCounter  prometheus.Counter `help:"Help Message"`
			SecondCounter prometheus.Counter `help:"This is the second counter"`
		}
	}
	if err := CreateAndRegisterMetricsFromStruct(&testMetrics); err != nil {
		t.Fatalf("CreateAndRegisterMetricsFromStruct(%v) failed with %v", &testMetrics, err)
	}
	defer prometheus.Unregister(testMetrics.Group.FirstCounter)
	defer prometheus.Unregister(testMetrics.Group.SecondCounter)
	testMetrics.Group.FirstCounter.Inc()
	testMetrics.Group.SecondCounter.Inc()
	m = metricsFromGatherer(t)
	if !m[wantMetric1] || !m[wantMetric2] {
		t.Errorf("prometheus.DefaultGatherer.Gather() got metrics %v, want metrics %q and %q to be included", m, wantMetric1, wantMetric2)
	}
}

func TestCreateAndRegisterMetricsFromStruct_CounterMultipleGroups(t *testing.T) {
	wantMetric1, wantMetric2 := "group_first_counter", "another_second_counter"
	m := metricsFromGatherer(t)
	if m[wantMetric1] || m[wantMetric2] {
		t.Fatalf("prometheus.DefaultGatherer.Gather() got metrics %v, want metrics %q and %q not included", m, wantMetric1, wantMetric2)
	}
	var testMetrics struct {
		Group struct {
			FirstCounter prometheus.Counter `help:"Help Message"`
		}
		Another struct {
			SecondCounter prometheus.Counter `help:"This is the second counter"`
		}
	}
	if err := CreateAndRegisterMetricsFromStruct(&testMetrics); err != nil {
		t.Fatalf("CreateAndRegisterMetricsFromStruct(%v) failed with %v", &testMetrics, err)
	}
	defer prometheus.Unregister(testMetrics.Group.FirstCounter)
	defer prometheus.Unregister(testMetrics.Another.SecondCounter)
	testMetrics.Group.FirstCounter.Inc()
	testMetrics.Another.SecondCounter.Inc()
	m = metricsFromGatherer(t)
	if !m[wantMetric1] || !m[wantMetric2] {
		t.Errorf("prometheus.DefaultGatherer.Gather() got metrics %v, want metrics %q and %q to be included", m, wantMetric1, wantMetric2)
	}
}

func TestCreateAndRegisterMetricsFromStruct_EmptyHelp(t *testing.T) {
	var testMetrics struct {
		Group struct {
			TestCounter prometheus.Counter `help:""`
		}
	}

	if err := CreateAndRegisterMetricsFromStruct(&testMetrics); err == nil {
		t.Fatalf("CreateAndRegisterMetricsFromStruct(%v) returned nil error, want non nil error", &testMetrics)
	}
}

func TestCreateAndRegisterMetricsFromStruct_NoHelp(t *testing.T) {
	var testMetrics struct {
		Group struct {
			TestCounter prometheus.Counter
		}
	}

	if err := CreateAndRegisterMetricsFromStruct(&testMetrics); err == nil {
		t.Fatalf("CreateAndRegisterMetricsFromStruct(%v) returned nil error, want non nil error", &testMetrics)
	}
}

func TestCreateAndRegisterMetricsFromStruct_BadStruct(t *testing.T) {
	var testMetrics struct {
		TestCounter prometheus.Counter
	}

	if err := CreateAndRegisterMetricsFromStruct(&testMetrics); err == nil {
		t.Fatalf("CreateAndRegisterMetricsFromStruct(%v) returned nil error, want non nil error", &testMetrics)
	}
}

func TestCreateAndRegisterMetricsFromStruct_Gauge(t *testing.T) {
	wantMetric := "group_first_gauge"
	m := metricsFromGatherer(t)
	if m[wantMetric] {
		t.Fatalf("prometheus.DefaultGatherer.Gather() got metrics %v, want metric %q not to be included", m, wantMetric)
	}
	var testMetrics struct {
		Group struct {
			FirstGauge prometheus.Gauge `help:"Help Message"`
		}
	}
	if err := CreateAndRegisterMetricsFromStruct(&testMetrics); err != nil {
		t.Fatalf("CreateAndRegisterMetricsFromStruct(%v) failed with %v", &testMetrics, err)
	}
	defer prometheus.Unregister(testMetrics.Group.FirstGauge)
	testMetrics.Group.FirstGauge.Set(1)
	m = metricsFromGatherer(t)
	if !m[wantMetric] {
		t.Errorf("prometheus.DefaultGatherer.Gather() got metrics %v, want metric %q to be included", m, wantMetric)
	}
}

func TestCreateAndRegisterMetricsFromStruct_GaugeVec(t *testing.T) {
	wantMetric := "group_first_gauge_vec"
	m := metricsFromGatherer(t)
	if m[wantMetric] {
		t.Fatalf("prometheus.DefaultGatherer.Gather() got metrics %v, want metric %q not to be included", m, wantMetric)
	}
	var testMetrics struct {
		Group struct {
			FirstGaugeVec *prometheus.GaugeVec `help:"Help Message" labels:"label1,label2"`
		}
	}
	if err := CreateAndRegisterMetricsFromStruct(&testMetrics); err != nil {
		t.Fatalf("CreateAndRegisterMetricsFromStruct(%v) failed with %v", &testMetrics, err)
	}
	defer prometheus.Unregister(testMetrics.Group.FirstGaugeVec)
	testMetrics.Group.FirstGaugeVec.With(prometheus.Labels{
		"label1": "value1",
		"label2": "value2",
	}).Set(1)
	m = metricsFromGatherer(t)
	if !m[wantMetric] {
		t.Errorf("prometheus.DefaultGatherer.Gather() got metrics %v, want metric %q to be included", m, wantMetric)
	}
}

func TestCreateAndRegisterMetricsFromStruct_Histogram(t *testing.T) {
	wantMetric := "group_first_hist"
	m := metricsFromGatherer(t)
	if m[wantMetric] {
		t.Fatalf("prometheus.DefaultGatherer.Gather() got metrics %v, want metric %q not to be included", m, wantMetric)
	}
	var testMetrics struct {
		Group struct {
			FirstHist prometheus.Histogram `help:"Help Message" buckets:"0.01,0.1,1,2"`
		}
	}
	if err := CreateAndRegisterMetricsFromStruct(&testMetrics); err != nil {
		t.Fatalf("CreateAndRegisterMetricsFromStruct(%v) failed with %v", &testMetrics, err)
	}
	defer prometheus.Unregister(testMetrics.Group.FirstHist)
	testMetrics.Group.FirstHist.Observe(0.015)
	m = metricsFromGatherer(t)
	if !m[wantMetric] {
		t.Errorf("prometheus.DefaultGatherer.Gather() got metrics %v, want metric %q to be included", m, wantMetric)
	}
}

func TestCreateAndRegisterMetricsFromStruct_HistogramVec(t *testing.T) {
	wantMetric := "group_hist_vec"
	m := metricsFromGatherer(t)
	if m[wantMetric] {
		t.Fatalf("prometheus.DefaultGatherer.Gather() got metrics %v, want metric %q not to be included", m, wantMetric)
	}
	var testMetrics struct {
		Group struct {
			HistVec *prometheus.HistogramVec `help:"Help Message" buckets:"0.01,0.1,1,2" labels:"label1,label2"`
		}
	}
	if err := CreateAndRegisterMetricsFromStruct(&testMetrics); err != nil {
		t.Fatalf("CreateAndRegisterMetricsFromStruct(%v) failed with %v", &testMetrics, err)
	}
	defer prometheus.Unregister(testMetrics.Group.HistVec)
	testMetrics.Group.HistVec.With(prometheus.Labels{
		"label1": "value1",
		"label2": "value2",
	}).Observe(0.015)
	testMetrics.Group.HistVec.With(prometheus.Labels{"label1": "var1", "label2": "var2"}).Observe(0.1)
	m = metricsFromGatherer(t)
	if !m[wantMetric] {
		t.Fatalf("prometheus.DefaultGatherer.Gather() got metrics %v, want metric %q to be included", m, wantMetric)
	}
}

func TestCreateAndRegisterMetricsFromStruct_HistogramAndCounter(t *testing.T) {
	wantMetric1, wantMetric2 := "group_hist", "group_counter"
	m := metricsFromGatherer(t)
	if m[wantMetric1] || m[wantMetric2] {
		t.Fatalf("prometheus.DefaultGatherer.Gather() got metrics %v, want metrics %q and %q not included", m, wantMetric1, wantMetric2)
	}
	var testMetrics struct {
		Group struct {
			Hist    prometheus.Histogram `help:"Help Hist" buckets:"0.01,0.1,1,2"`
			Counter prometheus.Counter   `help:"Help Counter"`
		}
	}
	if err := CreateAndRegisterMetricsFromStruct(&testMetrics); err != nil {
		t.Fatalf("CreateAndRegisterMetricsFromStruct(%v) failed with %v", &testMetrics, err)
	}
	defer prometheus.Unregister(testMetrics.Group.Hist)
	defer prometheus.Unregister(testMetrics.Group.Counter)
	testMetrics.Group.Hist.Observe(0.015)
	testMetrics.Group.Counter.Inc()
	m = metricsFromGatherer(t)
	if !m[wantMetric1] || !m[wantMetric2] {
		t.Errorf("prometheus.DefaultGatherer.Gather() got metrics %v, want metrics %q and %q to be included", m, wantMetric1, wantMetric2)
	}
}

func TestCreateAndRegisterMetricsFromStruct_HistogramMultipleGroups(t *testing.T) {
	wantMetric1, wantMetric2, wantMetric3 := "hist_first", "hist_second", "counter_first"
	m := metricsFromGatherer(t)
	if m[wantMetric1] || m[wantMetric2] || m[wantMetric3] {
		t.Fatalf("prometheus.DefaultGatherer.Gather() got metrics %v, want metrics %q, %q and %q not included", m, wantMetric1, wantMetric2, wantMetric3)
	}
	var testMetrics struct {
		Hist struct {
			First  prometheus.Histogram `help:"Help Fist Hist" buckets:"0.01,0.1,1,2"`
			Second prometheus.Histogram `help:"Help Second Hist" buckets:"0.5,1.5,2.5"`
		}
		Counter struct {
			First prometheus.Counter `help:"Help Counter"`
		}
	}
	if err := CreateAndRegisterMetricsFromStruct(&testMetrics); err != nil {
		t.Fatalf("CreateAndRegisterMetricsFromStruct(%v) failed with %v", &testMetrics, err)
	}
	defer prometheus.Unregister(testMetrics.Hist.First)
	defer prometheus.Unregister(testMetrics.Hist.Second)
	defer prometheus.Unregister(testMetrics.Counter.First)
	testMetrics.Hist.First.Observe(0.015)
	testMetrics.Hist.Second.Observe(0.5)
	testMetrics.Counter.First.Inc()
	m = metricsFromGatherer(t)
	if !m[wantMetric1] || !m[wantMetric2] || !m[wantMetric3] {
		t.Errorf("prometheus.DefaultGatherer.Gather() got metrics %v, want metrics %q, %q and %q to be included", m, wantMetric1, wantMetric2, wantMetric3)
	}
}

func TestCreateAndRegisterMetricsFromStruct_HistogramNoBuckets(t *testing.T) {
	var testMetrics struct {
		Group struct {
			FirstHist prometheus.Histogram `help:"Help Message"`
		}
	}

	if err := CreateAndRegisterMetricsFromStruct(&testMetrics); err == nil {
		t.Fatalf("CreateAndRegisterMetricsFromStruct(%v) returned nil error, want non nil error", &testMetrics)
	}
}

func TestCreateAndRegisterMetricsFromStruct_HistogramBadBuckets(t *testing.T) {
	var testMetrics struct {
		Group struct {
			FirstHist prometheus.Histogram `help:"Help Message" buckets:"foo,bar"`
		}
	}

	if err := CreateAndRegisterMetricsFromStruct(&testMetrics); err == nil {
		t.Fatalf("CreateAndRegisterMetricsFromStruct(%v) returned nil error, want non nil error", &testMetrics)
	}
}

func TestCreateAndRegisterMetricsFromStruct_Summary(t *testing.T) {
	wantMetric := "group_first_summary"
	m := metricsFromGatherer(t)
	if m[wantMetric] {
		t.Fatalf("prometheus.DefaultGatherer.Gather() got metrics %v, want metric %q not to be included", m, wantMetric)
	}
	var testMetrics struct {
		Group struct {
			FirstSummary prometheus.Summary `help:"Help Message"`
		}
	}
	if err := CreateAndRegisterMetricsFromStruct(&testMetrics); err != nil {
		t.Fatalf("CreateAndRegisterMetricsFromStruct(%v) failed with %v", &testMetrics, err)
	}
	defer prometheus.Unregister(testMetrics.Group.FirstSummary)
	testMetrics.Group.FirstSummary.Observe(0.015)
	m = metricsFromGatherer(t)
	if !m[wantMetric] {
		t.Errorf("prometheus.DefaultGatherer.Gather() got metrics %v, want metric %q to be included", m, wantMetric)
	}
}

func TestCreateAndRegisterMetricsFromStruct_SummaryVec(t *testing.T) {
	wantMetric := "group_summary_vec"
	m := metricsFromGatherer(t)
	if m[wantMetric] {
		t.Fatalf("prometheus.DefaultGatherer.Gather() got metrics %v, want metric %q not to be included", m, wantMetric)
	}
	var testMetrics struct {
		Group struct {
			SummaryVec *prometheus.SummaryVec `help:"Help Message" labels:"label1,label2"`
		}
	}
	if err := CreateAndRegisterMetricsFromStruct(&testMetrics); err != nil {
		t.Fatalf("CreateAndRegisterMetricsFromStruct(%v) failed with %v", &testMetrics, err)
	}
	defer prometheus.Unregister(testMetrics.Group.SummaryVec)
	testMetrics.Group.SummaryVec.With(prometheus.Labels{
		"label1": "value1",
		"label2": "value2",
	}).Observe(0.015)
	testMetrics.Group.SummaryVec.With(prometheus.Labels{
		"label1": "var1",
		"label2": "var2",
	}).Observe(0.1)
	m = metricsFromGatherer(t)
	if !m[wantMetric] {
		t.Errorf("prometheus.DefaultGatherer.Gather() got metrics %v, want metric %q to be included", m, wantMetric)
	}
}

func TestCreateAndRegisterMetricsFromStruct_UnsupportedMetricType(t *testing.T) {
	var testObservers struct {
		Group struct {
			ObserverVec prometheus.ObserverVec `help:"Help Message" labels:"label1,label2"`
		}
	}
	if err := CreateAndRegisterMetricsFromStruct(&testObservers); err == nil {
		t.Errorf("CreateAndRegisterMetricsFromStruct(%v) returned nil error, want non nil error", &testObservers)
	}
}

func metricsFromGatherer(t *testing.T) map[string]bool {
	t.Helper()
	m := map[string]bool{}
	values, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		t.Fatalf("prometheus.DefaultGatherer.Gather() failed with %v", err)
	}
	for _, v := range values {
		m[*v.Name] = true
	}
	return m
}
