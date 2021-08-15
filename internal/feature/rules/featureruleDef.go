package rules

type KeyValue struct {
	Key   string   `yaml:"key" json:"key"`
	Value []string `yaml:"value" json:"value"`
}

type FeatureRule struct {
	Environment []string     `yaml:"environment" json:"environment,omitempty"`
	QueryString KeyValue     `yaml:"querystring" json:"querystring"`
	Header      KeyValue     `yaml:"header" json:"header"`
	Platform    []string     `yaml:"platform" json:"platform,omitempty"`
	Version     []string     `yaml:"version" json:"version,omitempty"`
	Country     []string     `yaml:"country" json:"country,omitempty"`
	City        []string     `yaml:"city" json:"city,omitempty"`
	UserId      []string     `yaml:"userId" json:"userId,omitempty"`
	UserGroup   []string     `yaml:"userGroup" json:"userGroup,omitempty"`
	FeatureTm   FeatureTimer `yaml:"featureTimer" json:"featureTimer"`
	/*
		Experimentation struct {
			Id    string `yaml:"id"`
			Range struct {
				lowestScore  string `yaml:"lowestScore"`
				highestScore string `yaml:"highestScore"`
			} `yaml:"range"`
			Score []string `yaml:"score"`
		} `yaml:"experimentation"`*/

}
