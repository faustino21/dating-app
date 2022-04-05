package config

import (
	"dating_app_last/manager"
	"dating_app_last/util"
	"fmt"
	"github.com/spf13/viper"
)

type ApiConfig struct {
	Url string
}

type Manager struct {
	InfraManager   manager.InfraManager
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
}

type DbConfig struct {
	Host     string
	User     string
	Port     string
	Password string
	Name     string
}

type Config struct {
	Manager
	ApiConfig
	DbConfig
	LogLevel string
}

func (c Config) config(path, fileName string) Config {
	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigName(fileName)
	v.SetConfigType("yaml")
	v.AddConfigPath(path)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	c.DbConfig = DbConfig{
		Host:     v.GetString("datingapp.db.host"),
		User:     v.GetString("datingapp.db.user"),
		Password: v.GetString("datingapp.db.password"),
		Port:     v.GetString("datingapp.db.port"),
		Name:     v.GetString("datingapp.db.name"),
	}
	c.ApiConfig = ApiConfig{Url: v.GetString("datingapp.api.url")}

	c.LogLevel = v.GetString("datingapp.api.log_level")

	return c
}

func NewConfig(path, name string) Config {
	cfg := Config{}
	cfg = cfg.config(path, name)
	util.NewLog(cfg.LogLevel)

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	cfg.InfraManager = manager.NewInfra(dataSourceName)
	cfg.RepoManager = manager.NewRepoManager(cfg.InfraManager)
	cfg.UseCaseManager = manager.NewUseCaseManager(cfg.RepoManager)

	return cfg
}
