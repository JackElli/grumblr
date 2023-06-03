package main

import (
	"grumblrapi/endpoints"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	endpoints := endpoints.NewEndpointsMgr()
	endpoints.SetupEndpoints(r)

	err := http.ListenAndServe(":3200", r)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
