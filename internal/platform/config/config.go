package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Specification contains structured configuration variables.
type Specification struct {
	ActivitiesFile string `envconfig:"ACTIVITIES_FILE" default:"/resources/activities/activities.json" required:"true"`
	Output         string `envconfig:"OUTPUT" default:"standard" required:"true"`
}

// LoadEnv load config variables into baseConfig.
func LoadEnv() (*Specification, error) {
	var config Specification
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	return &config, err
}
