package main

import (
	"log"
	"menu/src/di"
	"net/http"
	"strconv"
)

func handleHealth(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

func main() {
	container := di.NewDi()

	http.HandleFunc("/health", handleHealth)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(container.Config.HttpServer.ListenPort), nil))
}
