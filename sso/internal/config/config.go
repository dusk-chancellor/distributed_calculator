package config

// sso config package

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server 	   Server
	Database   Database
	Redis 	   Redis
	JWT 	   JWT
}

type EnvConfig struct {
	ConfigPath string `env:"CONFIG_PATH" env-default:"configs/local.yml"`
}

type Server struct {
	Port 	int			  `yaml:"port"`
	Env 	string		  `yaml:"env"` // dev / stage / prod
	Timeout time.Duration `yaml:"timeout"`
}

type Database struct {
	Host	 string `yaml:"host"`
	Port	 string `yaml:"port"`
	User	 string `yaml:"user"`
	Password string `yaml:"password"`
	Name	 string `yaml:"name"`
	SSLMode	 string `yaml:"sslmode"`
}

type Redis struct {
	Host	 string `yaml:"host"`
	Port	 string `yaml:"port"`
	Password string `yaml:"password"`
	DB		 int    `yaml:"db"`
}

type JWT struct {
	Secret 				 string 	   `yaml:"secret"`
	AccessTokenDuration  time.Duration `yaml:"access_token_duration"`
	RefreshTokenDuration time.Duration `yaml:"refresh_token_duration"`
}
// Loads .yml config from path in .env
func LoadConfig() (*Config, error) {
	var envCfg EnvConfig
	// Read .env
	err := cleanenv.ReadEnv(&envCfg)
	if err != nil {
		return nil, err
	}

	var cfg Config
	// Read .yml
	err = cleanenv.ReadConfig(envCfg.ConfigPath, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
