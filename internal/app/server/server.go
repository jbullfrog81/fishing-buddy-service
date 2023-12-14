package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/jbullfrog81/fishing-buddy-service/internal/app/controller"
	"github.com/jbullfrog81/fishing-buddy-service/internal/app/stores"
)

//type Controllers struct {
//	Fbs *v0controllers.FbsController
//}

//type App struct {
//	Controllers *Controllers
//}

//type Server struct {
//	*App
//}

//func New() (*Server, error) {
//
//}

type Stores struct {
	Ifish *stores.IfishStore
}

type App struct {
	Stores *Stores
}

func NewApp() *App {
	return &App{
		Stores: &Stores{},
	}
}

//For some reason calling this does note work :thinking:

func ServerStart() error {
	fmt.Println("SERVERSTART")
	mux := http.NewServeMux()
	fmt.Println("MUX")
	mux.HandleFunc("/", controller.GetRoot)
	fmt.Println("ROOT /")
	mux.HandleFunc("/hello", controller.GetHello)
	fmt.Println("HELLO")

	err := http.ListenAndServe(":8080", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("ERROR Starting Server")
		fmt.Println(err)
		return err
	}

	return nil
}
