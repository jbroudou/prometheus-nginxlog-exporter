package relabeling

import "github.com/martin-helmich/prometheus-nginxlog-exporter/pkg/config"

// DefaultRelabelings are hardcoded relabeling configs that are always there
// and do not need to be explicitly configured
var DefaultRelabelings = []*Relabeling{
	{
		config.RelabelConfig{
			TargetLabel: "method",
			SourceValue: "request",
			Split:       1,

			WhitelistExists: true,
			WhitelistMap: map[string]interface{}{
				"GET":     nil,
				"HEAD":    nil,
				"POST":    nil,
				"PUT":     nil,
				"DELETE":  nil,
				"CONNECT": nil,
				"OPTIONS": nil,
				"TRACE":   nil,
				"PATCH":   nil,
			},
		},
	},
	{
		config.RelabelConfig{
			TargetLabel: "status",
			SourceValue: "status",
		},
	},
	{
		config.RelabelConfig{
			TargetLabel: "chain",
			SourceValue: "chain",
		},
	},
	{
		config.RelabelConfig{
			TargetLabel: "error_code",
			SourceValue: "error_code",
		},
	},
	{
		config.RelabelConfig{
			TargetLabel: "country",
			SourceValue: "geo_country",
		},
	},
	{
		config.RelabelConfig{
			TargetLabel: "chain_method",
			SourceValue: "chain_method",
		},
	},
	{
		config.RelabelConfig{
			TargetLabel: "aws_region",
			SourceValue: "aws_region",
		},
	},
}
