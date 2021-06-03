package config

type Config struct {
	Api struct {
		Port string `yaml:"port", envconfig:"SERVER_PORT"`
		Type string `yaml:"type", envconfig:"SERVER_TYPE"`
	} `yaml:"api"`
	Database struct {
		Username string `yaml:"user", envconfig:"DB_USERNAME"`
		Password string `yaml:"pass", envconfig:"DB_PASSWORD"`
		Url      string `yaml:"url", envconfig:"DB_URL"`
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
}
