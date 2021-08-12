package rules

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
	FeatureTm FeatureTimer `yaml:"timer"`
}
