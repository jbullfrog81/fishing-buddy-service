package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jbullfrog81/fishing-buddy-service/internal/app/controller"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	sugar.Infow("STARTING SERVER",
		"time", time.Second,
	)

	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.GetRoot)
	mux.HandleFunc("/hello", controller.GetHello)
	mux.HandleFunc("/weather", controller.GetWeather)

	err := http.ListenAndServe(":3333", mux)

	if errors.Is(err, http.ErrServerClosed) {
		sugar.Info("Server Closed")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		sugar.Info("error starting server")
		os.Exit(1)
	}
}
