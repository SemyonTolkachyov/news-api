package config

import (
	"github.com/gofiber/fiber/v3/log"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Name string `mapstructure:"APP_NAME" envDefault:"news_api"`

	Host         string        `mapstructure:"HOST" envDefault:"localhost"`
	Port         string        `mapstructure:"PORT" envDefault:"3000"`
	ReadTimeout  time.Duration `mapstructure:"READ_TIMEOUT" envDefault:"10s"`
	WriteTimeout time.Duration `mapstructure:"WRITE_TIMEOUT" envDefault:"10s"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     uint   `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBSslMode  string `mapstructure:"DB_SSL_MODE"`
}

func NewConfig(cfgName string) (Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(cfgName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("error reading configs: %s", err.Error())
		return Config{}, err
	}
	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Error("parse error config ", err.Error())
		return Config{}, err
	}
	return cfg, nil
}
