package main

import (
	"log"
	"menu/src/di"
	"menu/src/http/router"
	"net/http"
	"strconv"
)

func main() {
	container := di.NewDi()

	http.Handle("/", router.Router(container))

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(container.Config.HttpServer.ListenPort), nil))
}
