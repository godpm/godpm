package config

import (
	"io/fs"
	"path/filepath"

	"github.com/godpm/godpm/pkg/common/utils"
	"github.com/godpm/godpm/pkg/log"
)

// ReadProcessConfigs read all process configs
func ReadProcessConfigs(path string) (processConfigs []*ProcessConfig) {

	path, _ = filepath.Abs(path)
	processConfigs = []*ProcessConfig{}
	err := filepath.Walk(path, func(p string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatal().Fatal("walk failed ", err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		pc := &ProcessConfig{}
		utils.ReadConfig(p, pc)
		processConfigs = append(processConfigs, pc)
		return nil
	})

	if err != nil {
		log.Error().Println("Read configuration failed ", err)
	}

	return
}
