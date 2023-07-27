package config

import (
	"bytes"
	_ "embed"
	"strings"

	"github.com/spf13/viper"
)

//go:embed config.yaml
var configYaml []byte

var Conf *config

type config struct {
	App struct {
		Env          string `mapstructure:"env"`
		LogLevel     string `mapstructure:"log_level"`
		Name         string `mapstructure:"name"`
		URL          string `mapstructure:"url"`
		AllowOrigins string `mapstructure:"allow_origins"`
		AuthSecret   string `mapstructure:"auth_secret"`
	} `mapstructure:"app"`
	Db struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Name     string `mapstructure:"name"`
		Port     string `mapstructure:"port"`
	} `mapstructure:"db"`
}

func (c config) IsLocal() bool {
	return c.App.Env == "local"
}

func (c config) IsStaging() bool {
	return c.App.Env == "staging"
}

func (c config) IsProduction() bool {
	return c.App.Env == "production"
}

func init() {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadConfig(bytes.NewBuffer(configYaml)); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(err)
	}
}
