package main

import (
	"grumblrapi/cors"
	"grumblrapi/endpoints"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	r := mux.NewRouter()

	endpoints := endpoints.NewEndpointsMgr(logger)
	err := endpoints.SetupEndpoints(r)
	if err != nil {
		logger.Error("Cannot setup endpoints", zap.Error(err))
	}

	logger.Info("Started grumblr api")
	err = http.ListenAndServe(":3200", cors.CORS(r))
	if err != nil {
		log.Fatal("Cannot start server")
	}

}
