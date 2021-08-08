package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"sync"
)

type Config struct {
}

var (
	once sync.Once
	conf Config
)

func GetConfig() Config {
	once.Do(func() {
		_ = godotenv.Load()
		if err := envconfig.Process("", &conf); err != nil {
			log.Fatal(err)
		}
	})

	return conf
}
