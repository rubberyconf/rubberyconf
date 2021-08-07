package feature

import (
	"container/list"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rubberyconf/rubberyconf/internal/logs"
	"gopkg.in/yaml.v2"
)

type FeatureRule struct {
	Environment []string   `yaml:"environment"`
	QueryString QueryParam `yaml:"querystring"`
	Header      struct {
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
}

type FeatureDefinition struct {
	Name string `yaml:"name"`
	Meta struct {
		Description string   `yaml:"description"`
		Tags        []string `yaml:"tags"`
	} `yaml:"meta"`

	Default struct {
		Value struct {
			Data interface{} `yaml:"data"`
			Type string      `yaml:"type"`
		} `yaml:"value"`
		TTL string `yaml:"ttl"`
	} `yaml:"default"`

	Configurations []struct {
		ConfigId       string        `yaml:"id"`
		RulesBehaviour string        `yaml:"rulesBehaviour"`
		Rules          []FeatureRule `yaml:"rules"`
		Value          interface{}   `yaml:"value"`
		Rollout        struct {
			Strategy       string `yaml:"strategy"`
			EnabledForOnly string `yaml:"enabledForOnly"`
			Selector       string `yaml:"selector"`
		} `yaml:"rollout"`
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

func checkRules(r FeatureRule, vars map[string]string) (int, *list.List) {
	total := 0
	matches := list.New()
	if len(r.Environment) > 0 {
		ok := ruleEnvironment(r.Environment)
		if ok {
			total += 1
			matches.PushBack("environment")
		}
	} else if r.QueryString.Key != "" {
		ok := queryString(r.QueryString, vars)
		if ok {
			total += 1
			matches.PushBack("querystring")
		}
	}

	return total, matches
}

func (conf *FeatureDefinition) SelectRule(vars map[string]string) (interface{}, bool, string, *list.List) {

	for _, c := range conf.Configurations {
		total := 0
		totalMatches := list.New()
		for _, r := range c.Rules {
			matches, labelMatches := checkRules(r, vars)
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
