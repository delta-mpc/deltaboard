package config

const (
	EnvProduction = "production"
	EnvDebug      = "debug"
	EnvTest       = "test"
)

type AppConfig struct {
	Environment string `mapstructure:"environment"`

	Db struct {
		Driver           string `mapstructure:"driver"`
		ConnectionString string `mapstructure:"connection"`
	} `mapstructure:"db"`

	Log struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"log"`

	Http struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	} `mapstructure:"http"`

	Session      SessionConfig `mapstructure:"session"`
	AllowOrigins []string      `mapstructure:"allow_origins"`
}

type SessionConfig struct {
	Driver             string `mapstructure:"driver"`
	Connection         string `mapstructure:"connection"`
	SecretKey          string `mapstructure:"secret_key"`
	LoginMaxAgeSeconds int    `mapstructure:"login_max_age_seconds"`
	VaultMaxAgeSeconds int    `mapstructure:"vault_max_age_seconds"`
}
