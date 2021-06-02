package api

import (
	"log"
	"os"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

var onceConfig sync.Once

var cfgSingleton *Config

func GetConfiguration() *Config {
	return cfgSingleton
}

func NewConfiguration(filePath string) *Config {

	onceConfig.Do(func() {

		cfgSingleton = new(Config)
		cfgSingleton.readFile(filePath)
		cfgSingleton.readEnv()
		log.Printf("Configuration loaded: %+v", *cfgSingleton)
	})
	return cfgSingleton
}

func (config *Config) readFile(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(config)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
}

func (config *Config) readEnv() {
	err := envconfig.Process("", config)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
}
