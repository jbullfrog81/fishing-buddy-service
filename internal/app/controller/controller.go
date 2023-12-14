package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jbullfrog81/fishing-buddy-service/internal/app/clients"
	"github.com/jbullfrog81/fishing-buddy-service/internal/app/requests"
	"github.com/jbullfrog81/fishing-buddy-service/internal/app/stores"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is the fishing buddy service!\n")
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

type WeatherController struct {
	Cache *redis.Client
	Ctx   context.Context
}

//NewWeatherController constructs a new WeatherController,
//ensuring that the dependencies are valid
func NewWeatherController() *WeatherController {
	return &WeatherController{}
}

func (wthr *WeatherController) GetWeather(w http.ResponseWriter, r *http.Request) {

	//weather forecast cache key
	//wther-forecast-<weather office>-<grid x>-<grid y>-<YYYY>-<MM>-<DD>-<HH>
	//wther-forecast-<weather office>-<grid x>-<grid y>-<EPOCH UTC rounded to nearest hour>
	cTime := time.Now()
	tRound := cTime.Round(time.Hour).UTC().Unix()

	wthrCacheKey := "wthr-forecast-eax-32-32-" + strconv.FormatInt(tRound, 10)
	fmt.Printf("The weather cache key is:%s\n", wthrCacheKey)

	//TODO - get results from cache correctly
	//TODO - add context
	val, err := wthr.Cache.Get(wthr.Ctx, wthrCacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Printf("Cache entry not found!\n")
			//Not in cache get from the weather service
			wthrData := clients.Weather()

			wthrDataMar, err := json.Marshal(wthrData)
			if err != nil {
				//TODO - add return and structured logging
				fmt.Println(err)
			}

			//TODO - add expiration to cache key
			err = wthr.Cache.Set(wthr.Ctx, wthrCacheKey, wthrDataMar, 0).Err()
			if err != nil {
				panic(err)
			}

			//TODO - Better engineer this return when there is an error
			val, err = wthr.Cache.Get(wthr.Ctx, wthrCacheKey).Result()
		}
		//TODO - add structured logging
		fmt.Println(err)
	}

	//Print weather data
	//fmt.Printf("Value from cache:%s", val)

	fmt.Fprintf(w, "weather is, %s\n", val)

}

type FishingController struct {
	IfishStore *stores.IfishStore
	Logger     *slog.Logger
}

//NewFishingController constructs a new FishingController ensuring that the dependencies are valid
func NewFishingController(ifishStore *stores.IfishStore, logger *slog.Logger) *FishingController {
	return &FishingController{
		IfishStore: ifishStore,
		Logger:     logger,
	}
}

// FishOn - takes in the coordinates of where a fish is caught
func (c *FishingController) PostCatch(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var req requests.PostCatchRequestBody

	//TODO - decode and validate request headers

	err := requests.ValidateRequestHeader(r.Header.Get("Content-Type"), "application/json")
	if err != nil {
		c.Logger.ErrorContext(ctx, "Invalid request header found")
		http.Error(w, fmt.Errorf("Invalid header Content-Type").Error(), http.StatusUnsupportedMediaType)
	}

	// Enforce a maximum read of 1MB from the response body and if greater then
	// Decode() will return error "http: request body too large".
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	// Decode request body
	rdc := json.NewDecoder(r.Body)
	// Do not allow unexpected fields. When unknown fields are experienced then
	// Decode() will return error "json: unknown field ..."
	rdc.DisallowUnknownFields()

	rdc.Decode(&req)
	if err != nil {
		c.Logger.ErrorContext(ctx, "Invalid request body found")
		http.Error(w, fmt.Errorf("Invalid request body").Error(), http.StatusBadRequest)
		return
	}

	c.Logger.Info("successful POST catch request",
		"latitude", req.Coordinates.Latitude,
		"longitude", req.Coordinates.Longitude,
		"fish_species_id", req.FishSpeciesId,
		"fisherman_id", req.FishermanId,
	)

	//TODO - validate request

	ok := req.FishSpeciesId.IsValid()
	if !ok {
		c.Logger.ErrorContext(ctx, "Invalid fish species id in request body")
		http.Error(w, fmt.Errorf("Invalid request body").Error(), http.StatusBadRequest)
		return
	}

	//fishSpeciesId := 1

	//fishermanId := 1

	//var coordinates = stores.Coordinates{
	//	Latitude:  1,
	//	Longitude: 1,
	//}

	err = c.IfishStore.NewCatch(ctx, req.FishSpeciesId, req.FishermanId, req.Coordinates)
	if err != nil {
		c.Logger.ErrorContext(ctx, "error with the iFish database",
			"error", err,
		)
		http.Error(w, fmt.Errorf("Server Error").Error(), http.StatusInternalServerError)
		return
	}

	//TODO: Proper response
	fmt.Fprintf(w, "ok\n")

}
