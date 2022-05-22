package utils

import (
	"os"

	"github.com/godpm/godpm/pkg/log"
	"gopkg.in/yaml.v2"
)

// ReadConfig read config from file
func ReadConfig(path string, conf interface{}) {
	f, err := os.OpenFile(path, os.O_CREATE, 0666)
	if err != nil {
		log.Fatal().Fatalf("%v", err)
	}

	defer f.Close()

	err = yaml.NewDecoder(f).Decode(conf)
	if err != nil {
		log.Fatal().Fatalf("%v", err)
	}
}
