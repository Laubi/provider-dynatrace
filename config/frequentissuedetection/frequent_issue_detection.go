package frequentissuedetection

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("frequent_issue_detection", func(r *config.Resource) {
		r.ShortGroup = "frequent_issue_detection"
	})
}
