package config

import "github.com/spf13/viper"

type Config struct {
	Postgres Postgres
}

type Postgres struct {
	Host            string
	Port            string
	Login           string
	Password        string
	ConnMaxLifeTime int
	MaxOpenConn     int
	MaxIdleConn     int
}

func ReadConfig() *Config {
	var cfg Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
