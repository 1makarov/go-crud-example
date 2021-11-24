package config

import (
	"github.com/1makarov/go-crud-example/internal/db"
	"github.com/spf13/viper"
	"os"
	"time"
)

const (
	EnvLocal = "local"
	Prod     = "prod"
)

type (
	Config struct {
		Environment string
		DB          db.ConfigDB
		HTTP        HTTPConfig
		CacheTTL    time.Duration `mapstructure:"ttl"`
	}

	HTTPConfig struct {
		Host               string
		Port               string
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}
)

func Init(configFolder string) (*Config, error) {
	if err := parseConfig(configFolder); err != nil {
		return nil, err
	}

	var cfg Config

	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	return viper.UnmarshalKey("cache.ttl", &cfg.CacheTTL)
}

func setFromEnv(cfg *Config) {
	cfg.HTTP.Host = os.Getenv("HTTP_HOST")
	cfg.HTTP.Port = os.Getenv("HTTP_PORT")

	cfg.DB.Name = os.Getenv("DB_NAME")
	cfg.DB.User = os.Getenv("DB_USER")
	cfg.DB.Host = os.Getenv("DB_HOST")
	cfg.DB.Password = os.Getenv("DB_PASSWORD")

	cfg.Environment = os.Getenv("ENV")
}

func parseConfig(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}
