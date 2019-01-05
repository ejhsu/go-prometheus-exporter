package main

import "github.com/prometheus/client_golang/prometheus"

// Exporter defines the metrics types
type Exporter struct {
	counter    prometheus.Counter
	counterVec prometheus.CounterVec
}

// NewExporter returns a custom exporter
func NewExporter(metricsPrefix string) *Exporter {
	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: metricsPrefix,
		Name:      "counter_metric",
		Help:      "This is a counter for number of total api calls"})

	counterVec := *prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: metricsPrefix,
		Name:      "counter_vec_metric",
		Help:      "This is a counter vec for number of all api calls"},
		[]string{"endpoint"})

	return &Exporter{
		counter:    counter,
		counterVec: counterVec,
	}
}

// Collect implements prometheus.Collector.Collect
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.counter.Collect(ch)
	e.counterVec.Collect(ch)
}

// Describe implements prometheus.Collector.Describe
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.counter.Describe(ch)
	e.counterVec.Describe(ch)
}

// IncreCounter increments counter
func (e *Exporter) IncreCounter() {
	e.counter.Inc()
}

// IncreCounterWithEndpoint increments counter
func (e *Exporter) IncreCounterWithEndpoint(endpoint string) {
	e.counterVec.With(prometheus.Labels{"endpoint": endpoint}).Inc()
}
