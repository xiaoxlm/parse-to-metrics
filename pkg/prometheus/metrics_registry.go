package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	io_prometheus_client "github.com/prometheus/client_model/go"
)

func NewMetricsRegistry(labels map[string]string, collector ...prometheus.Collector) *CustomMetricsRegistry {
	metricsRegistry := NewCustomMetricsRegistry(labels)
	metricsRegistry.MustRegister(collectors.NewGoCollector())

	for _, c := range collector {
		metricsRegistry.MustRegister(c)
	}

	return metricsRegistry
}

type CustomMetricsRegistry struct {
	*prometheus.Registry
	customLabels []*io_prometheus_client.LabelPair
}

func NewCustomMetricsRegistry(labels map[string]string) *CustomMetricsRegistry {
	c := &CustomMetricsRegistry{
		Registry: prometheus.NewRegistry(),
	}

	for k, v := range labels {
		c.customLabels = append(c.customLabels, &io_prometheus_client.LabelPair{
			Name:  &k,
			Value: &v,
		})
	}

	return c
}

func (g *CustomMetricsRegistry) Gather() ([]*io_prometheus_client.MetricFamily, error) {
	metricFamilies, err := g.Registry.Gather()

	for _, metricFamily := range metricFamilies {
		metrics := metricFamily.Metric
		for _, metric := range metrics {
			metric.Label = append(metric.Label, g.customLabels...)
		}
	}

	return metricFamilies, err
}
