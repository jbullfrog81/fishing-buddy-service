package controller

import (
	"fmt"
	"io"
	"net/http"

	"github.com/jbullfrog81/fishing-buddy-service/internal/app/clients"
	redis "github.com/redis/go-redis/v9"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is the fishing buddy service!\n")
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

type Weather struct {
	Cache *redis.Client
}

func (wthr *Weather) GetWeather(w http.ResponseWriter, r *http.Request) {

	//TODO - get results from cache correctly
	//TODO - add context
	val, err := wthr.Cache.Get(nil, "id1234").Result()
	if err != nil {
		if err == redis.Nil {
			//Not in cache get from the weather service
			clients.Weather()

			//TODO - store results in cache
			//TODO - add context
			//TODO - add expiration to cache key
			err := wthr.Cache.Set(nil, "key", "value", 0).Err()
			if err != nil {
				panic(err)
			}
		}
		//TODO - add structured logging
		fmt.Println(err)
	}

	//Print weather data
	fmt.Print(val)

}
