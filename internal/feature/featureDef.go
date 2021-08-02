package feature

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rubberyconf/rubberyconf/internal/logs"
	"gopkg.in/yaml.v2"
)

type FeatureDefinition struct {
	Name string `yaml:"name"`
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
				Type string      `yaml:"type"`
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

func (conf *FeatureDefinition) LoadFromYaml(payload interface{}) error {

	aux := fmt.Sprintf("%v", payload)
	decoder := yaml.NewDecoder(strings.NewReader(aux))
	err := decoder.Decode(conf)
	return err
}

func (conf *FeatureDefinition) LoadFromJsonBinary(b []byte) error {

	err := json.Unmarshal(b, &conf)
	return err
}

func (conf *FeatureDefinition) LoadFromString(text string) error {

	err := yaml.Unmarshal([]byte(text), &conf)
	if err != nil {
		logs.GetLogs().WriteMessage("error", "error unmarshalling yaml content to featureDefinition", nil)
		return err
	}
	return nil
}
func (conf *FeatureDefinition) ToString() (string, error) {

	b, err := yaml.Marshal(conf)
	if err != nil {
		logs.GetLogs().WriteMessage("error", "error marshalling yaml content to featureDefinition", nil)
		return "", err
	}
	sb := string(b)
	return sb, nil
}

func (conf *FeatureDefinition) GetFinalValue() (interface{}, error) {

	var afterCast interface{}
	data := conf.Default.Value.Data

	switch conf.Default.Value.Type {
	case "string":
		afterCast = data.(string)
	case "json":
		b, err := json.MarshalIndent(data, "", "   ")
		if err != nil {
			logs.GetLogs().WriteMessage("error", "error marshalling content of featureDefinition to json", err)
			return nil, err
		}
		afterCast = string(b)
	case "number":
		afterCast = data.(int)
	}

	return afterCast, nil
}
