// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Tetragon

package kprobemetrics

import (
	"github.com/cilium/tetragon/pkg/metrics/consts"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	MergeErrors = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   consts.MetricsNamespace,
		Name:        "generic_kprobe_merge_errors_total",
		Help:        "The total number of failed attempts to merge a kprobe and kretprobe event.",
		ConstLabels: nil,
	}, []string{"curr_fn", "curr_type", "prev_fn", "prev_type"})
	MergeOkTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace:   consts.MetricsNamespace,
		Name:        "generic_kprobe_merge_ok_total",
		Help:        "The total number of successful attempts to merge a kprobe and kretprobe event.",
		ConstLabels: nil,
	})
	MergePushed = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace:   consts.MetricsNamespace,
		Name:        "generic_kprobe_merge_pushed_total",
		Help:        "The total number of pushed events for later merge.",
		ConstLabels: nil,
	})
)

func InitMetrics(registry *prometheus.Registry) {
	registry.MustRegister(MergeErrors)
	registry.MustRegister(MergeOkTotal)
	registry.MustRegister(MergePushed)
}

// Get a new handle on the mergeErrors metric for a current and previous function
// name and probe type
func GetMergeErrors(currFn, currType, prevFn, prevType string) prometheus.Counter {
	return MergeErrors.WithLabelValues(currFn, currType, prevFn, prevType)
}

// Increment the mergeErrors metric for a current and previous function
// name and probe type
func MergeErrorsInc(currFn, currType, prevFn, prevType string) {
	GetMergeErrors(currFn, currType, prevFn, prevType).Inc()
}

func MergeOkTotalInc() {
	MergeOkTotal.Inc()
}

func MergePushedInc() {
	MergePushed.Inc()
}
