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

	Session             SessionConfig `mapstructure:"session"`
	VaultExpiredSeconds int64         `mapstructure:"vault_expired_seconds"`
	AllowOrigins        []string      `mapstructure:"allow_origins"`

	Email struct {
		AccessKeyId               string `mapstructure:"access_key_id"`
		AccessKeySecret           string `mapstructure:"access_key_secret"`
		AccountName               string `mapstructure:"account_name"`
		FromAlias                 string `mapstructure:"from_alias"`
		RegionId                  string `mapstructure:"region_id"`
		VerificationLink          string `mapstructure:"verification_link"`
		VerificationEffectiveTime int64  `mapstructure:"verification_effective_time"`
		EmailExpirationTime       int64  `mapstructure:"email_expiration_time"`
	} `mapstructure:"email"`
	SMS struct {
		SmsUser string `mapstructure:"sms_user"`
		SmsKey  string `mapstructure:"sms_key"`
	} `mapstructure:"sms"`
	OP struct {
		PlatformEmail         string `mapstructure:"platform_email"`
		PlatformAdministrator string `mapstructure:"platform_administrator"`
	} `json:"op"`

	Integrals struct {
		SingleIntegrals  int64 `mapstructure:"single_integrals"`
		InitialIntegrals int64 `mapstructure:"initial_integrals"`
		BaseIntegrals    int64 `mapstructure:"base_integrals"`
	} `mapstructure:"integrals"`

	Identity struct {
		Url     string `mapstructure:"url"`
		AppCode string `mapstructure:"app_code"`
	} `mapstructure:"identity"`

	FileStorageType int `mapstructure:"file_storage_type"`
}

type SessionConfig struct {
	Driver             string `mapstructure:"driver"`
	Connection         string `mapstructure:"connection"`
	SecretKey          string `mapstructure:"secret_key"`
	LoginMaxAgeSeconds int    `mapstructure:"login_max_age_seconds"`
	VaultMaxAgeSeconds int    `mapstructure:"vault_max_age_seconds"`
}
