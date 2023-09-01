package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"net/http"
)

type сonfig struct {
	ServerPort string `toml:"serverPort"`
}

func handleNewsLetter(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

func main() {
	cfg := &сonfig{}
	_, err := toml.DecodeFile("src/config/config.toml", cfg)
	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	println(cfg.ServerPort)

	http.HandleFunc("/health", handleNewsLetter)
	log.Fatal(http.ListenAndServe(cfg.ServerPort, nil))
}
