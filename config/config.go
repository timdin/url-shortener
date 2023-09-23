package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DB     DBConfig     `yaml:"database"`
	Server ServerConfig `yaml:"service"`
	Cache  CacheConfig  `yaml:"redis"`
}

type ServerConfig struct {
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	AcceptExpired  bool   `yaml:"accept_expired"`
	AcceptNoExpire bool   `yaml:"accept_no_expire"`
}

type DBConfig struct {
	Conn string `yaml:"conn"`
}

type CacheConfig struct {
	Conn string `yaml:"conn"`
}

func NewConfig() *Config {
	cfg := &Config{}
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
