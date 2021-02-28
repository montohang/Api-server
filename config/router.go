package config

import (
	"api_server/utils"
	"fmt"
	"log"
	"net/http"
	"os"

	logs "github.com/MaulIbra/logs_module"
	"github.com/gorilla/mux"
)

func MuxRouter() *mux.Router {
	return mux.NewRouter()
}

func ListenServe(router *mux.Router) {
	host := os.Getenv("SERVER_HOST")
	if os.Getenv("SERVER_HOST") == "" {
		host = utils.GetEnv("SERVER_HOST", "localhost")
	}

	port := os.Getenv("SERVER_PORT")
	if os.Getenv("SERVER_PORT") == "" {
		port = utils.GetEnv("SERVER_PORT", "7000")
	}

	log.Print(fmt.Sprintf("%v:%v", host, port))
	err := http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), router)
	if err != nil {
		logs.ErrorLogger.Println(err)
		panic(err)
	}
}