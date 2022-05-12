package utils

import (
	"os"

	"github.com/phuslu/log"
	"gopkg.in/yaml.v2"
)

// ReadConfig read config from file
func ReadConfig(path string, conf interface{}) {
	f, err := os.OpenFile(path, os.O_CREATE, 0666)
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}

	defer f.Close()

	err = yaml.NewDecoder(f).Decode(conf)
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}
}
