package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jbullfrog81/fishing-buddy-service/internal/app/clients"
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
