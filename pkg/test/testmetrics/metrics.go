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

// Package testmetrics gets information from metrics for testing.
package testmetrics

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/protobuf/testing/protocmp"
	pb "github.com/prometheus/client_model/go"
)

type retrieveMetricsFunc func() map[string]*pb.MetricFamily

// Retriever retrieves prometheus metrics and provides an API to get metric values.
// The 'initial' snapshot is taken when the Retriever is created
// and the 'final' snapshot is taken when the first metric is obtained.
type Retriever struct {
	retrieve      retrieveMetricsFunc
	snapshotDelay func()
	initial       map[string]*pb.MetricFamily
	final         map[string]*pb.MetricFamily
}

// NewRetrieverFromGatherer creates a new metrics retriever that retrieves metrics from the
// in memory default prometheus gatherer.
func NewRetrieverFromGatherer(t *testing.T) *Retriever {
	ma := &Retriever{
		retrieve: func() map[string]*pb.MetricFamily {
			families, err := prometheus.DefaultGatherer.Gather()
			if err != nil {
				t.Fatalf("prometheus.DefaultGatherer.Gather() failed with %v", err)
			}
			familyMap := make(map[string]*pb.MetricFamily)
			for _, f := range families {
				familyMap[*f.Name] = f
			}
			return familyMap
		},
	}
	ma.snapshotBefore()
	return ma
}

func (ma *Retriever) snapshotBefore() *Retriever {
	ma.initial = ma.retrieve()
	return ma
}

func (ma *Retriever) snapshotAfter() *Retriever {
	if ma.snapshotDelay != nil {
		ma.snapshotDelay()
	}
	ma.final = ma.retrieve()
	return ma
}

func (ma *Retriever) maybeSnapshotAfter() *Retriever {
	if ma.final == nil {
		ma.snapshotAfter()
	}
	return ma
}

// GetCounterValues returns the values for both the "initial" and "final" time of a counter type metric.
// If the metric is not a counter, this method fails.
// If the metric does not exist, this method returns 0s.
func (ma *Retriever) GetCounterValues(t *testing.T, metricName string, labels map[string]string) (float64, float64) {
	t.Helper()
	ma.maybeSnapshotAfter()

	final := maybeGetMetric(t, ma.final, metricName, labels)
	if final == nil {
		// Metric may not have been created or it remains unchanged, so return 0s.
		return 0, 0
	}
	if final.GetCounter() == nil {
		t.Fatalf("Can't get counter values. Metric %q is not a counter", metricName)
	}

	initial := maybeGetMetric(t, ma.initial, metricName, labels)
	if initial == nil {
		// Metric may not have been created in the 'initial' time, so return 0.
		return 0, final.GetCounter().GetValue()
	}

	return initial.GetCounter().GetValue(), final.GetCounter().GetValue()
}

// GetHistogramBuckets returns the buckets for both the "initial" and "final" time of a histogram type metric.
// If the metric is not a histogram, this method fails.
// If the metric does not exist, this method returns empty buckets.
func (ma *Retriever) GetHistogramBuckets(t *testing.T, metricName string, labels map[string]string) ([]*pb.Bucket, []*pb.Bucket) {
	t.Helper()
	ma.maybeSnapshotAfter()

	final := maybeGetMetric(t, ma.final, metricName, labels)
	if final == nil {
		// Metric may not have been created and remain unchanged, so return empty buckets.
		return []*pb.Bucket{}, []*pb.Bucket{}
	}
	if final.GetHistogram() == nil {
		t.Fatalf("Can't get histogram values. Metric %q is not a histogram", metricName)
	}

	initial := maybeGetMetric(t, ma.initial, metricName, labels)
	if initial == nil {
		// Metric may not have been created in the 'initial' time, so create same number of empty bucket as "final".
		var zeroedBuckets []*pb.Bucket
		for _, b := range final.GetHistogram().GetBucket() {
			zero := uint64(0)
			zeroedBuckets = append(zeroedBuckets, &pb.Bucket{
				CumulativeCount: &zero,
				UpperBound:      b.UpperBound,
			})
		}
		return zeroedBuckets, final.GetHistogram().GetBucket()
	}

	return initial.GetHistogram().GetBucket(), final.GetHistogram().GetBucket()
}

// IncrementBucketsCumulatively increments the given buckets according to the provided value and
// increment step. Note that incrementing a bucket will increment all the higher buckets as it is a
// cumulative counter per bucket.
func (ma *Retriever) IncrementBucketsCumulatively(buckets []*pb.Bucket, value float64, incrementStep int) []*pb.Bucket {
	var result []*pb.Bucket
	for i := range buckets {
		upperBound := buckets[i].GetUpperBound()
		cumulativeCount := buckets[i].GetCumulativeCount()
		if upperBound >= value {
			cumulativeCount += uint64(incrementStep)
		}
		result = append(result, &pb.Bucket{
			UpperBound:      &upperBound,
			CumulativeCount: &cumulativeCount,
		})
	}
	return result
}

func maybeGetMetric(t *testing.T, metrics map[string]*pb.MetricFamily, metricName string, labels map[string]string) *pb.Metric {
	family, found := metrics[metricName]
	if !found {
		t.Logf("Metric %q not found\nAvailable:\n%v", metricName, strings.Join(metricNames(metrics), "\n"))
		return nil
	}

	if labels == nil {
		labels = make(map[string]string)
	}

	var availableLabels []string
	for _, metric := range family.Metric {
		metricLabels := getLabels(metric)
		if cmp.Equal(metricLabels, labels, protocmp.Transform()) {
			return metric
		}
		availableLabels = append(availableLabels, fmt.Sprint(metricLabels))
	}
	t.Logf("Metric %q not found with labels: %v\nAvailable: %v", metricName, labels, strings.Join(availableLabels, "\n"))
	return nil
}

func getLabels(metric *pb.Metric) map[string]string {
	labels := make(map[string]string)
	for _, label := range metric.Label {
		labels[*label.Name] = *label.Value
	}
	return labels
}

func metricNames(metrics map[string]*pb.MetricFamily) []string {
	var metricNames []string
	for k := range metrics {
		metricNames = append(metricNames, k)
	}
	sort.Strings(metricNames)
	return metricNames
}
