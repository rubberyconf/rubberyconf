package feature

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

type RubberyConfig struct {
	Meta struct {
		Description string `yaml:"description"`
		Tags        []string
	} `yaml:"meta"`

	Value interface{} `yaml:"value"`

	TimeToLive struct {
		Value int    `yaml:"value"`
		Unit  string `yaml:"unit"`
	} `yaml:"timeToLive"`

	Rollout string `yaml:"rollout"`

	Environment []string `yaml:"environment"`
}

func (conf *RubberyConfig) Load(payload interface{}) error {

	aux := fmt.Sprintf("%v", payload)
	decoder := yaml.NewDecoder(strings.NewReader(aux))
	err := decoder.Decode(conf)
	return err
}
