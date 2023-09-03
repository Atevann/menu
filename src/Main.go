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
	//DI initialization
	container := di.NewDi()

	println(container.Config.HttpServer.ListenPort)

	http.HandleFunc("/health", handleHealth)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(container.Config.HttpServer.ListenPort), nil))
}
