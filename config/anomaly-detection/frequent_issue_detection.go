package anomaly_detection

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("dynatrace_frequent_issues", func(r *config.Resource) {
		r.ShortGroup = "anomaly-detection"
		r.Kind = "FrequentIssue"
	})

	p.AddResourceConfigurator("dynatrace_metric_events", func(r *config.Resource) {
		r.ShortGroup = "anomaly-detection"
		r.Kind = "MetricEvent"
	})
}
