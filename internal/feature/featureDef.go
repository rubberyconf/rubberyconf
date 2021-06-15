package feature

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

type RubberyConfig struct {
	Meta struct {
		Description string   `yaml:"description"`
		Tags        []string `yaml:"tags"`
	} `yaml:"meta"`

	Default struct {
		Value struct {
			Data interface{} `yaml:"data"`
			Type interface{} `yaml:"type"`
		} `yaml:"value"`
		TTL string `yaml:"ttl"`
	} `yaml:"default"`

	Configurations []struct {
		Config struct {
			Id             string `yaml:"id"`
			RulesBehaviour string `yaml:"rulesBehaviour"`
			Rules          []struct {
				Environment []string `yaml:"environment"`
				QueryString struct {
					Key   string   `yaml:"key"`
					Value []string `yaml:"value"`
				} `yaml:"querystring"`
				Header struct {
					Key   string   `yaml:"key"`
					Value []string `yaml:"value"`
				} `yaml:"header"`
				Platform        []string `yaml:"platform"`
				Version         []string `yaml:"version"`
				Country         []string `yaml:"country"`
				City            []string `yaml:"city"`
				UserId          []string `yaml:"userId"`
				UserGroup       []string `yaml:"userGroup"`
				Experimentation struct {
					Id    string `yaml:"id"`
					Range struct {
						lowestScore  string `yaml:"lowestScore"`
						highestScore string `yaml:"highestScore"`
					} `yaml:"range"`
					Score []string `yaml:"score"`
				} `yaml:"experimentation"`
			} `yaml:"rules"`
			Value struct {
				Data interface{} `yaml:"data"`
				Type interface{} `yaml:"type"`
			} `yaml:"value"`
			TTL     string `yaml:"ttl"`
			Rollout struct {
				Strategy       string `yaml:"strategy"`
				EnabledForOnly string `yaml:"enabledForOnly"`
				Selector       string `yaml:"selector"`
			} `yaml:"rollout"`
		} `yaml:"Config"`
	} `yaml:"configurations"`
}

func (conf *RubberyConfig) Load(payload interface{}) error {

	aux := fmt.Sprintf("%v", payload)
	decoder := yaml.NewDecoder(strings.NewReader(aux))
	err := decoder.Decode(conf)
	return err
}
