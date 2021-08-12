package rules

type KeyValue struct {
	Key   string   `yaml:"key"`
	Value []string `yaml:"value"`
}

type FeatureRule struct {
	Environment     []string `yaml:"environment"`
	QueryString     KeyValue `yaml:"querystring"`
	Header          KeyValue `yaml:"header"`
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
	FeatureTm FeatureTimer `yaml:"timer"`
}
