package main

import (
	"grumblrapi/endpoints"
	"grumblrapi/main/cors"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	ENVIRONMENT, exists := os.LookupEnv("ENVIRONMENT")
	if !exists {
		panic("cannot find environment in the env file!")
	}

	r := mux.NewRouter()

	endpoints := endpoints.NewEndpointsMgr(logger)
	err := endpoints.SetupEndpoints(ENVIRONMENT, r)
	if err != nil {
		logger.Error("Cannot setup endpoints", zap.Error(err))
		return
	}

	logger.Info("Started grumblr api")
	err = http.ListenAndServe(":3200", cors.CORS(r, ENVIRONMENT))
	if err != nil {
		log.Fatal("Cannot start server")
	}
}
