package options

import "github.com/TroyXia/iam/internal/pkg/server"

// FeatureOptions contains configuration items related to API server features.
type FeatureOptions struct {
	EnableProfiling bool `json:"profiling"      mapstructure:"profiling"`
	EnableMetrics   bool `json:"enable-metrics" mapstructure:"enable-metrics"`
}

// ApplyTo applies the run options to the method receiver and returns self.
func (o *FeatureOptions) ApplyTo(c *server.Config) error {
	c.EnableProfiling = o.EnableProfiling
	c.EnableMetrics = o.EnableMetrics

	return nil
}
