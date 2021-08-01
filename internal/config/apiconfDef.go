package config

type Config struct {
	Api struct {
		Port       string   `yaml:"port", envconfig:"RUBBERYCONF_PORT"`
		Cache      string   `yaml:"cache", envconfig:"RUBBERYCONF_CACHE"`
		Source     string   `yaml:"source", envconfig:"RUBBERYCONF_TYPE"`
		Logs       []string `yaml:"logs"`
		DefaultTTL string   `yaml:"default_ttl", envconfig:"RUBBERYCONG_DEFAULT_TTL", json:"default_ttl"`
		Options    struct {
			LogLevel string `yaml:"loglevel", envconfig:"RUBBERYCONG_LOGLEVEL"`
		} `yaml:"options"`
	} `yaml:"api"`
	Database struct {
		Collections struct {
			Metrics   string `yaml:"metrics", envconfig:"DB_COL_METRICS"`
			SessionID string `yaml:"sessionids", envconfig:"DB_COL_SESSIONIDS"`
		} `yaml:"collections"`
		Url          string `yaml:"url", envconfig:"DB_URL"`
		DatabaseName string `yaml:"databasename", envconfig:"DB_DATABASENAME"`
	} `yaml:"database"`
	Redis struct {
		Username string `yaml:"user", envconfig:"REDIS_USERNAME"`
		Password string `yaml:"pass", envconfig:"REDIS_PASSWORD"`
		Url      string `yaml:"url", envconfig:"REDIS_URL"`
	} `yaml:"redis"`
	GitServer struct {
		Username string `yaml:"user", envconfig:"GIT_USERNAME"`
		Password string `yaml:"pass", envconfig:"GIT_PASSWORD"`
		Url      string `yaml:"url", envconfig:"GIT_URL"`
		ApiToken string `yaml:"apitoken", envconfig:"GIT_APITOKEN"`
	} `yaml:"gitserver"`
	Elastic struct {
		Url  string `yaml:"url", envconfig:"ELASTIC_URL"`
		Logs struct {
			Index string `yaml:"index", envconfig:"ELASTIC_LOGS_INDEX"`
		} `yaml:"logs"`
	} `yaml:"elastic"`
}
