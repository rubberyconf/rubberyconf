package config

type Config struct {
	Api struct {
		Port       string `yaml:"port", envconfig:"SERVER_PORT"`
		Type       string `yaml:"type", envconfig:"SERVER_TYPE"`
		SourceType string `yaml:"sourcetype", envconfig:"SOURCE_TYPE"`
	} `yaml:"api"`
	Database struct {
		Collections struct {
			Metrics string `yaml:"metrics", envconfig:"DB_COL_METRICS"`
		} `yaml:"collections"`
		Url          string `yaml:"url", envconfig:"DB_URL"`
		DatabaseName string `yaml:"databasename", envconfig:"DB_DATABASENAME"`
	} `yaml:"database"`
	Cache struct {
		Username string `yaml:"user", envconfig:"CACHE_USERNAME"`
		Password string `yaml:"pass", envconfig:"CACHE_PASSWORD"`
		Url      string `yaml:"url", envconfig:"CACHE_URL"`
	} `yaml:"cache"`
	GitServer struct {
		Username string `yaml:"user", envconfig:"GIT_USERNAME"`
		Password string `yaml:"pass", envconfig:"GIT_PASSWORD"`
		Url      string `yaml:"url", envconfig:"GIT_URL"`
		ApiToken string `yaml:"apitoken", envconfig:"GIT_APITOKEN"`
	} `yaml:"gitserver"`
	Elastic struct {
		Url   string `yaml:"url", envconfig:"ELASTIC_URL"`
		Index string `yaml:"index", envconfig:"ELASTIC_INDEX"`
	} `yaml:"elastic"`
}
