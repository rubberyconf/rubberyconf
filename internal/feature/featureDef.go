package feature

import (
	"container/list"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rubberyconf/rubberyconf/internal/feature/rules"
	"github.com/rubberyconf/rubberyconf/internal/logs"
	"gopkg.in/yaml.v2"
)

type FeatureDefinition struct {
	Name string `yaml:"name" json:"name"`
	Meta struct {
		Description string   `yaml:"description" json:"description"`
		Tags        []string `yaml:"tags" json:"tags,omitempty"`
	} `yaml:"meta" json:"meta"`

	Default struct {
		Value struct {
			Data interface{} `yaml:"data" json:"data"`
			Type string      `yaml:"type" json:"type"`
		} `yaml:"value" json:"value"`
		TTL string `yaml:"ttl" json:"ttl"`
	} `yaml:"default" json:"default"`

	Configurations []struct {
		ConfigId       string              `yaml:"id" json:"id"`
		RulesBehaviour string              `yaml:"rulesBehaviour" json:"rulesBehaviour"`
		Rules          []rules.FeatureRule `yaml:"rules" json:"rules,omitempty"`
		Value          interface{}         `yaml:"value" json:"value"`
		Rollout        struct {
			Strategy       string `yaml:"strategy" json:"strategy"`
			EnabledForOnly string `yaml:"enabledForOnly" json:"enabledForOnly"`
			Selector       string `yaml:"selector" json:"selector"`
		} `yaml:"rollout" json:"rollout"`
	} `yaml:"configurations" json:"configurations,omitempty"`
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

func (conf *FeatureDefinition) GetFinalValue(vars map[string]string) (interface{}, error) {

	var afterCast interface{}

	data, found, confId, matches := conf.SelectRule(vars)
	if !found {
		data = conf.Default.Value.Data
	}

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

	logs.GetLogs().WriteMessage("info", fmt.Sprintf("configuration applied: %s, matches: %v ", confId, matches), nil)

	return afterCast, nil
}

func (conf *FeatureDefinition) SelectRule(vars map[string]string) (interface{}, bool, string, *list.List) {

	ruleMast := NewRuleMaster()

	for _, c := range conf.Configurations {
		total := 0
		totalMatches := list.New()
		for _, r := range c.Rules {
			matches, labelMatches := ruleMast.CheckRules(r, vars)
			total += matches
			totalMatches.PushBackList(labelMatches)
		}

		logic := c.RulesBehaviour
		if logic == "" {
			logic = "AND"
		}
		if logic == "OR" && total > 1 {
			return c.Value, true, c.ConfigId, totalMatches

		} else if logic == "AND" && total == len(c.Rules) {
			return c.Value, true, c.ConfigId, totalMatches
		}
	}
	return nil, false, "", nil
}
