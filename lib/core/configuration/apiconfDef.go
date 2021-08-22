package configuration

type Config struct {
	Api struct {
		Port       string   `yaml:"port" envconfig:"RUBBERYCONF_PORT"`
		Cache      string   `yaml:"cache" envconfig:"RUBBERYCONF_CACHE"`
		Source     string   `yaml:"source" envconfig:"RUBBERYCONF_TYPE"`
		Logs       []string `yaml:"logs"`
		DefaultTTL string   `yaml:"default_ttl" envconfig:"RUBBERYCONG_DEFAULT_TTL" json:"default_ttl"`
		Options    struct {
			LogLevel string `yaml:"loglevel" envconfig:"RUBBERYCONG_LOGLEVEL"`
		} `yaml:"options"`
	} `yaml:"api"`
	Database struct {
		Collections struct {
			Metrics   string `yaml:"metrics" envconfig:"DB_COL_METRICS"`
			SessionID string `yaml:"sessionids" envconfig:"DB_COL_SESSIONIDS"`
			Features  string `yaml:"features" envconfig:"DB_COL_FEATURES"`
		} `yaml:"collections"`
		Url          string `yaml:"url" envconfig:"DB_URL"`
		DatabaseName string `yaml:"databasename" envconfig:"DB_DATABASENAME"`
		TimeOut      string `yaml:"timeout" envconfig:"DB_TIMEOUT"`
	} `yaml:"database"`
	Redis struct {
		Username string `yaml:"user" envconfig:"REDIS_USERNAME"`
		Password string `yaml:"pass" envconfig:"REDIS_PASSWORD"`
		Url      string `yaml:"url" envconfig:"REDIS_URL"`
		TimeOut  string `yaml:"timeout" envconfig:"REDIS_TIMEOUT"`
	} `yaml:"redis"`
	GitServer struct {
		Username      string `yaml:"user" envconfig:"GIT_USERNAME"`
		Email         string `yaml:"email" envconfig:"GIT_EMAIL"`
		Password      string `yaml:"pass" envconfig:"GIT_PASSWORD"`
		Url           string `yaml:"url" envconfig:"GIT_URL"`
		Repo          string `yaml:"repo" envconfig:"GIT_REPO"`
		Organization  string `yaml:"organization" envconfig:"GIT_ORGANIZATION"`
		ApiToken      string `yaml:"apitoken" envconfig:"GIT_APITOKEN"`
		BranchDefault string `yaml:"branchdefault" envconfig:"GIT_BRANCHDEFAULT"`
		TimeOut       string `yaml:"timeout" envconfig:"GIT_TIMEOUT"`
	} `yaml:"gitserver"`
	Elastic struct {
		Url  string `yaml:"url" envconfig:"ELASTIC_URL"`
		Logs struct {
			Index string `yaml:"index" envconfig:"ELASTIC_LOGS_INDEX"`
		} `yaml:"logs"`
		TimeOut string `yaml:"timeout" envconfig:"ELASTIC_TIMEOUT"`
	} `yaml:"elastic"`
}
