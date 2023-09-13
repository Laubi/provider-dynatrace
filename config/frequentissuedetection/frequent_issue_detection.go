package frequentissuedetection

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("dynatrace_frequent_issues", func(r *config.Resource) {
		r.ShortGroup = "frequent-issue-detection"
		r.Kind = "FrequentIssueDetection"
	})
}
