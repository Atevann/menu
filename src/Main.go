package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"net/http"
	"strconv"
)

type сonfig struct {
	ServerPort int `toml:"serverPort"`
}

func handleHealth(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

func main() {
	cfg := &сonfig{}
	_, err := toml.DecodeFile("src/config/config.toml", cfg)
	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	http.HandleFunc("/health", handleHealth)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(cfg.ServerPort), nil))
}
