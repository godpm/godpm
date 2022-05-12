package config

import (
	"github.com/godpm/godpm/pkg/common/utils"
)

// Config configuration struct
type Config struct {
	Port              int    `yaml:"port"`
	PprofPort         int    `yaml:"pprof_port"`
	ConfigurationPath string `yaml:"config_path"`
}

// SetupConfiguration set configuration with path
func SetupConfiguration(path string) {
	conf := &Config{}
	utils.ReadConfig(path, conf)
}
