package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"database"`
	Redis  RedisConfig  `yaml:"redis"`
	JWT    JWTConfig    `yaml:"jwt"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type DBConfig struct {
	Host            string        `yaml:"host"`
	Port            string        `yaml:"port"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	DBName          string        `yaml:"dbname"`
	MaxConns        int           `yaml:"max_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
	Name            string        `yaml:"name"`
}

type RedisConfig struct {
	Host        string        `yaml:"host"`
	Port        int           `yaml:"port"`
	Password    string        `yaml:"password"`
	DB          int           `yaml:"db"`
	PoolSize    int           `yaml:"pool_size"`
	DialTimeout time.Duration `yaml:"dial_timeout"`
}

type JWTConfig struct {
	Secret         string        `yaml:"secret"`          // 密钥
	ExpiresTime    time.Duration `yaml:"expires_time"`    // 过期时间
	RefreshExpires time.Duration `yaml:"refresh_expires"` // 刷新token过期时间
}

func ReadConfig(filename string) (*Config, error) {
	configFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(configFile, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
