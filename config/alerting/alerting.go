package alerting

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("dynatrace_alerting", func(r *config.Resource) {
		r.ShortGroup = "alerting"
		r.Kind = "Profile"
	})
}
