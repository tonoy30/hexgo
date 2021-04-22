package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/tonoy30/hexgo/config"
	"github.com/tonoy30/hexgo/pkg/http/rest"
)

const (
	PORT = ":8081"
)

func init() {
	config.SetUpConfiguration()
}
func main() {
	fmt.Println("🚀starting server on port " + PORT)
	router := rest.InitHandler()
	log.Fatal(http.ListenAndServe(PORT, cors.Default().Handler(router)))
}
