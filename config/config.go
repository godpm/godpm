package config

import (
	"github.com/godpm/godpm/pkg/common/utils"
)

// Config configuration struct
type Config struct {
	Port              int    `yaml:"port"`
	PprofPort         int    `yaml:"pprof_port"`
	ConfigurationPath string `yaml:"configuration_path"`
	LogFile           string `yaml:"log_file"`
	PidFile           string `yaml:"pid_file"`

	processConfigs []*ProcessConfig
}

// ProcessConfig process config
type ProcessConfig struct {
	Name         string  `yaml:"name"`
	Command      string  `yaml:"command"`            // process start command
	Environment  string  `yaml:"environment"`        // process start need env
	AutoStart    bool    `yaml:"auto_start"`         // if process can auto_start
	AutoRestart  bool    `yaml:"auto_restart"`       // if process can auto_restart
	User         string  `yaml:"user"`               // which user can to start the process
	Directory    *string `yaml:"directory"`          // process directory
	RetryTimes   *int    `yaml:"retry_times"`        // when start failed, retry start, try times
	StartSecs    *int    `yaml:"start_secs"`         // when start a process, then wait `start_secs` seconds to check if process is running
	RestartPause *int    `yaml:"restart_pause_secs"` // when start failed, wait `secs` to restart again
}

var AppConfig *Config

// SetupConfiguration set configuration with path
func SetupConfiguration(path string) {
	conf := &Config{}
	utils.ReadConfig(path, conf)
	AppConfig = conf

	AppConfig.processConfigs = ReadProcessConfigs(conf.ConfigurationPath)
}

// GetAllProcesssConfig get all process config
func (conf *Config) GetAllProcesssConfig() []*ProcessConfig {
	return conf.processConfigs
}
