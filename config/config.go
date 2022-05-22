package config

import (
	"github.com/godpm/godpm/pkg/common/utils"
)

// Config configuration struct
type Config struct {
	Port              int    `yaml:"port"`
	PprofPort         int    `yaml:"pprof_port"`
	ConfigurationPath string `yaml:"config_path"`
	LogFile           string `yaml:"log_file"`
	PidFile           string `yaml:"pid_file"`
}

type ProcessConfig struct {
	Name        string  `yaml:"name"`
	Command     string  `yaml:"command"`
	Environment string  `yaml:"environment"`
	AutoStart   bool    `yaml:"auto_start"`
	AutoRestart bool    `yaml:"auto_restart"`
	User        string  `yaml:"user"`
	Directory   *string `yaml:"directory"`
	RetryTimes  *int    `yaml:"retry_times"`
	StartSecs   *int    `yaml:"start_secs"`
}

var AppConfig *Config

// SetupConfiguration set configuration with path
func SetupConfiguration(path string) {
	conf := &Config{}
	utils.ReadConfig(path, conf)
	AppConfig = conf
}
