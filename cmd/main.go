package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jbullfrog81/fishing-buddy-service/internal/app/clients"
	"github.com/jbullfrog81/fishing-buddy-service/internal/app/controller"
	"go.uber.org/zap"
)

func main() {
	//ctx := context.Background()

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	sugar.Infow("STARTING SERVER",
		"time", time.Second,
	)

	var ctx = context.Background()

	//cfg := config.Config{}

	//TODO - read this from a config file
	//cacheConfig := config.Cache{
	//	Address:  "127.0.0.1:6379", // use local docker instance
	//	Password: "",               // no password set
	//	Db:       0,                // use default DB
	//}

	//cfg.SetConfig(cacheConfig)

	// Setup Cache Client
	//cacheClient := redis.NewClient(&redis.Options{
	//	Addr:     cfg.Cache.Address,
	//	Password: cfg.Cache.Password,
	//	DB:       cfg.Cache.Db,
	//})

	//cacheClient := redis.NewClient(&redis.Options{
	//	Addr: "127.0.0.1:6379",
	//})

	//sugar.Infow("CACHE CLIENT",
	//	"time", time.Second,
	//)

	//weather := &controller.WeatherController{cache: cacheClient}
	weather := controller.NewWeatherController()
	weather.Cache = clients.NewCacheClient(ctx)
	weather.Ctx = ctx

	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.GetRoot)
	mux.HandleFunc("/hello", controller.GetHello)
	//mux.HandleFunc("/weather", weather.GetWeather)
	mux.HandleFunc("/weather", weather.GetWeather)

	err := http.ListenAndServe(":8080", mux)

	if errors.Is(err, http.ErrServerClosed) {
		sugar.Info("Server Closed")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		sugar.Info("error starting server")
		os.Exit(1)
	}
}
