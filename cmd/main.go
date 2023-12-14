package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jbullfrog81/fishing-buddy-service/internal/app/clients"
	"github.com/jbullfrog81/fishing-buddy-service/internal/app/controller"
	"github.com/jbullfrog81/fishing-buddy-service/internal/app/server"
	"github.com/jbullfrog81/fishing-buddy-service/internal/app/stores"
	"github.com/jmoiron/sqlx"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("STARTING SERVER",
		"time", time.Second,
	)

	var ctx = context.Background()

	//TODO - Start utilizing environment variable configs and config struct
	//cacheHost := os.Getenv("FBS_CACHE_HOST")
	//cachePort := os.Getenv("FBS_CACHE_PORT")
	//cfg := config.Config{}
	//cfg.SetConfig(cacheConfig)

	weather := controller.NewWeatherController()

	weather.Cache = clients.NewCacheClient(ctx)
	weather.Ctx = ctx

	dbcfg := mysql.NewConfig()
	dbcfg.DBName = "ifish"
	dbcfg.User = "fisherman"
	dbcfg.Passwd = "fishon"
	dbcfg.Net = "tcp"
	dbcfg.Addr = net.JoinHostPort("mysql", strconv.FormatInt(int64(3306), 10))

	db, err := sqlx.ConnectContext(ctx, "mysql", dbcfg.FormatDSN())
	if err != nil {
		logger.Error("Error creating connection to ifish database",
			"database_address", dbcfg.Addr,
			"error", err,
		)
		panic(err)
	}

	app := server.NewApp()

	st := app.Stores

	st.Ifish = stores.NewIfishStore(db)

	fishingController := controller.NewFishingController(st.Ifish, logger)

	mux := http.NewServeMux()

	// Routes
	//mux.HandleFunc("/", controller.GetRoot)
	mux.HandleFunc("/hello", controller.GetHello)
	mux.HandleFunc("/weather", weather.GetWeather)
	mux.HandleFunc("/catch", fishingController.PostCatch)

	err = http.ListenAndServe(":8080", mux)

	if errors.Is(err, http.ErrServerClosed) {
		logger.Info("Server Closed")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		logger.Info("error starting server")
		os.Exit(1)
	}
}
