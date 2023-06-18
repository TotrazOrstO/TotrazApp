package config

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/multierr"
	"log"
	"os"
)

const (
	defaultHttpVersion = 1
	defaultConfigFile  = ".env"

	AppLevelProd AppLevel = "prod"
	AppLevelDev  AppLevel = "dev"
	AppLevelTest AppLevel = "test"
)

type (
	Config struct {
		App      Application
		HTTP     HTTP
		Postgres Postgres
	}

	Application struct {
		AppLevel AppLevel `json:"app_level" envconfig:"APP_LEVEL"`
		Debug    bool     `json:"debug" envconfig:"APP_DEBUG"`
	}

	HTTP struct {
		HTTPVersion int    `json:"http_version" envconfig:"HTTP_HTTPVERSION"`
		Protocol    string `json:"protocol" envconfig:"HTTP_PROTOCOL"`
		Host        string `json:"host" envconfig:"HTTP_HOST"`
		Port        int    `json:"port" envconfig:"HTTP_PORT"`
		Debug       bool   `json:"debug" envconfig:"HTTP_DEBUG"`
		AuthSecret  []byte `json:"auth_secret" envconfig:"HTTP_AUTH_SECRET"`
	}

	Postgres struct {
		Driver   string `json:"driver" envconfig:"POSTGRES_DRIVER"`
		Host     string `json:"host" envconfig:"POSTGRES_HOST"`
		Port     int    `json:"port" envconfig:"POSTGRES_PORT"`
		User     string `json:"user" envconfig:"POSTGRES_USER"`
		Password string `json:"password" envconfig:"POSTGRES_PASSWORD"`
		Name     string `json:"name" envconfig:"POSTGRES_NAME"`
		SslMode  string `json:"sslmode" envconfig:"POSTGRES_SSLMODE"`
	}

	AppLevel string
)

// NewConfig парсит, валидирует и возвращает конфигурации
func NewConfig() (Config, error) {
	var cfg Config
	loadEnvIfExists()

	if err := envconfig.Process("app", &cfg.App); err != nil {
		return Config{}, fmt.Errorf("can't load env for application: %w", err)
	}

	if err := envconfig.Process("http", &cfg.HTTP); err != nil {
		return Config{}, fmt.Errorf("can't load env for application: %w", err)
	}

	if cfg.HTTP.HTTPVersion == 0 {
		cfg.HTTP.HTTPVersion = defaultHttpVersion
	}

	if err := envconfig.Process("db_postgres", &cfg.Postgres); err != nil {
		return Config{}, fmt.Errorf("can't load env for aspojc api: %w", err)
	}

	cfg.HTTP.Debug = cfg.App.Debug

	if err := cfg.Validate(); err != nil {
		return Config{}, fmt.Errorf("failed to validate: %w", err)
	}

	//if cfg.App.Debug {
	//	prettyCfg, err := json.MarshalIndent(cfg, "", "  ")
	//	if err != nil {
	//		return Config{}, fmt.Errorf("failed to prettify configs: %w", err)
	//	}
	//	logrus.Info(string(prettyCfg))
	//}

	return cfg, nil
}

func (c *Config) Validate() error {
	return multierr.Combine(
		c.App.validate(),
		c.Postgres.validate(),
		c.HTTP.validate(),
	)
}

func (a Application) validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.AppLevel, validation.In(AppLevelProd, AppLevelDev, AppLevelTest)),
	)
}

func (b HTTP) validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Port, validation.Required),
		validation.Field(&b.HTTPVersion, validation.Required, validation.In(1, 2)),
		validation.Field(&b.AuthSecret, validation.Required),
	)
}

func (p Postgres) validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
		//validation.Field(&p.Host, validation.Required),
		validation.Field(&p.User, validation.Required),
		validation.Field(&p.Password, validation.Required),
		validation.Field(&p.Driver, validation.Required),
	)
}

func loadEnvIfExists() {
	info, err := os.Stat(defaultConfigFile)
	if os.IsNotExist(err) {
		return
	}

	if !info.IsDir() {
		if err := godotenv.Load(defaultConfigFile); err != nil {
			log.Fatalf("error reading env variables from file: %s\n", err.Error())
		}
	}
}
