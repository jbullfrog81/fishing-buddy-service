package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jbullfrog81/fishing-buddy-service/internal/app/config"
	"github.com/jbullfrog81/fishing-buddy-service/internal/app/controller"
	redis "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	sugar.Infow("STARTING SERVER",
		"time", time.Second,
	)

	cfg := config.Config{}

	//TODO - read this from a config file
	cacheConfig := config.Cache{
		Address:  "localhost:6379", // use local docker instance
		Password: "",               // no password set
		Db:       0,                // use default DB
	}

	cfg.SetConfig(cacheConfig)

	// Setup Cache Client
	cacheClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Cache.Address,
		Password: cfg.Cache.Password,
		DB:       cfg.Cache.Db,
	})
	sugar.Infow("CACHE CLIENT",
		"time", time.Second,
	)

	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.GetRoot)
	mux.HandleFunc("/hello", controller.GetHello)

	weather := &controller.Weather{Cache: cacheClient}
	mux.HandleFunc("/weather", weather.GetWeather)

	err := http.ListenAndServe(":3333", mux)

	if errors.Is(err, http.ErrServerClosed) {
		sugar.Info("Server Closed")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		sugar.Info("error starting server")
		os.Exit(1)
	}
}
