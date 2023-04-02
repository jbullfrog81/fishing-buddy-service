package controller

import (
	"fmt"
	"io"
	"net/http"

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

func GetWeather(w http.ResponseWriter, r *http.Request) {
	clients.Weather()
}
