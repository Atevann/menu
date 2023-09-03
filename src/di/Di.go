package di

import (
	"github.com/BurntSushi/toml"
	"log"
	"menu/config"
)

type Di struct {
	Config *config.Config
}

func NewDi() Di {
	return Di{Config: getConfig()}
}

func getConfig() *config.Config {
	cfg := &config.Config{}
	_, err := toml.DecodeFile("config/config.toml", cfg)
	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	return cfg
}
