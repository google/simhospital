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
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/google/simhospital/pkg/logging"
)

var (
	log = logging.ForCallerPackage()

	supportedMetrics = map[string]reflect.Type{
		"prometheus.Counter":       reflect.TypeOf((*prometheus.Counter)(nil)).Elem(),
		"*prometheus.CounterVec":   reflect.TypeOf((*prometheus.CounterVec)(nil)),
		"prometheus.Gauge":         reflect.TypeOf((*prometheus.Gauge)(nil)).Elem(),
		"*prometheus.GaugeVec":     reflect.TypeOf((*prometheus.GaugeVec)(nil)),
		"prometheus.Histogram":     reflect.TypeOf((*prometheus.Histogram)(nil)).Elem(),
		"*prometheus.HistogramVec": reflect.TypeOf((*prometheus.HistogramVec)(nil)),
		"prometheus.Summary":       reflect.TypeOf((*prometheus.Summary)(nil)).Elem(),
		"*prometheus.SummaryVec":   reflect.TypeOf((*prometheus.SummaryVec)(nil)),
	}
	camelCase = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z]+[^A-Z]+|$)")
)

const (
	helpTag    = "help"    // Identifies help text for the counter or histogram.
	labelsTag  = "labels"  // Identifies set of labels used to create a CounterVec or HistogramVec.
	bucketsTag = "buckets" // Identifies set of buckets to be used for histogram.
	namespace  = ""        // Common prefix given to all counters created by this package.
)

// CreateAndRegisterMetricsFromStruct accepts a struct and dynamically
// creates metrics from the field names and associated tags provided.
func CreateAndRegisterMetricsFromStruct(c interface{}) error {
	coll := reflect.ValueOf(c).Elem()
	st := reflect.TypeOf(c).Elem()
	for i := 0; i < coll.NumField(); i++ {
		if err := createMetricsCollection(coll.Field(i), st.Field(i).Name); err != nil {
			return errors.Wrapf(err, "Cannot register metrics: cannot create metrics collection for field %s", st.Field(i).Name)
		}
	}
	return nil
}

func createMetricsCollection(group reflect.Value, groupName string) error {
	if group.Type().Kind() != reflect.Struct {
		return errors.New("Metrics configuration must be a struct")
	}

	for i := 0; i < group.NumField(); i++ {
		metricType := group.Type().Field(i)
		c, err := collector(metricType, groupName)
		if err != nil {
			return err
		}
		metric := group.Field(i)
		metric.Set(reflect.ValueOf(c))
		prometheus.MustRegister(c)
	}
	return nil
}

func collector(metricType reflect.StructField, groupName string) (prometheus.Collector, error) {
	name := metricName(groupName, metricType)
	help, err := helpText(metricType)
	if err != nil {
		return nil, err
	}

	switch metricType.Type {
	case supportedMetrics["prometheus.Counter"]:
		return newCounter(name, help), nil
	case supportedMetrics["*prometheus.CounterVec"]:
		return newCounterVec(name, help, metricType)
	case supportedMetrics["prometheus.Gauge"]:
		return newGauge(name, help), nil
	case supportedMetrics["*prometheus.GaugeVec"]:
		return newGaugeVec(name, help, metricType)
	case supportedMetrics["prometheus.Histogram"]:
		return newHistogram(name, help, metricType)
	case supportedMetrics["*prometheus.HistogramVec"]:
		return newHistogramVec(name, help, metricType)
	case supportedMetrics["prometheus.Summary"]:
		return newSummary(name, help), nil
	case supportedMetrics["*prometheus.SummaryVec"]:
		return newSummaryVec(name, help, metricType)
	default:
		return nil, fmt.Errorf("field must be one of the following types of metrics: %v", allowedMetricTypes())
	}
}

func metricName(groupName string, metricType reflect.StructField) string {
	return fmt.Sprintf("%s_%s", toUnderscore(groupName), toUnderscore(metricType.Name))
}

func helpText(metricType reflect.StructField) (string, error) {
	help, ok := metricType.Tag.Lookup(helpTag)
	if !ok {
		return "", errMissingTag(metricType.Name, helpTag)
	}
	if help == "" {
		return "", errEmptyTag(metricType.Name, helpTag)
	}
	return help, nil
}

func allowedMetricTypes() []string {
	allowed := make([]string, 0)
	for m := range supportedMetrics {
		allowed = append(allowed, m)
	}
	return allowed
}

func newCounter(name, help string) prometheus.Counter {
	return prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	})
}

func newCounterVec(name, help string, metricType reflect.StructField) (*prometheus.CounterVec, error) {
	labels, err := parseLabels(metricType)
	if err != nil {
		return nil, err
	}
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	}, labels), nil
}

func newGauge(name, help string) prometheus.Gauge {
	return prometheus.NewGauge(prometheus.GaugeOpts{
		Name: name,
		Help: help,
	})
}

func newGaugeVec(name, help string, metricType reflect.StructField) (*prometheus.GaugeVec, error) {
	labels, err := parseLabels(metricType)
	if err != nil {
		return nil, err
	}
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: name,
		Help: help,
	}, labels), nil
}

func newHistogram(name, help string, metricType reflect.StructField) (prometheus.Histogram, error) {
	buckets, err := parseBuckets(metricType)
	if err != nil {
		return nil, err
	}
	return prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
		Buckets:   buckets,
	}), nil
}

func newHistogramVec(name, help string, metricType reflect.StructField) (*prometheus.HistogramVec, error) {
	buckets, err := parseBuckets(metricType)
	if err != nil {
		return nil, err
	}
	labels, err := parseLabels(metricType)
	if err != nil {
		return nil, err
	}
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
		Buckets:   buckets,
	}, labels), nil
}

func newSummary(name, help string) prometheus.Summary {
	return prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	})
}

func newSummaryVec(name, help string, metricType reflect.StructField) (*prometheus.SummaryVec, error) {
	labels, err := parseLabels(metricType)
	if err != nil {
		return nil, err
	}
	return prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	}, labels), nil
}

func parseLabels(metricType reflect.StructField) ([]string, error) {
	labels, ok := metricType.Tag.Lookup(labelsTag)
	if !ok {
		return nil, errMissingTag(metricType.Name, labelsTag)
	}
	if labels == "" {
		return nil, errEmptyTag(metricType.Name, labelsTag)
	}
	return strings.Split(labels, ","), nil
}

func parseBuckets(metricType reflect.StructField) ([]float64, error) {
	bucketsStr, ok := metricType.Tag.Lookup(bucketsTag)
	if !ok {
		return nil, errMissingTag(metricType.Name, bucketsTag)
	}
	if bucketsStr == "" {
		return nil, errEmptyTag(metricType.Name, bucketsTag)
	}
	split := strings.Split(bucketsStr, ",")
	buckets := make([]float64, 0, len(split))
	for _, b := range split {
		f, err := strconv.ParseFloat(b, 64)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to parse buckets tag: %v", metricType.Name, err)
		}
		buckets = append(buckets, f)
	}
	return buckets, nil
}

func toUnderscore(s string) string {
	var a []string
	for _, sub := range camelCase.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return strings.ToLower(strings.Join(a, "_"))
}

func errMissingTag(metricName, tag string) error {
	return fmt.Errorf("%s: missing %s tag", metricName, tag)
}

func errEmptyTag(metricName, tag string) error {
	return fmt.Errorf("%s: empty %s tag provided", metricName, tag)
}
