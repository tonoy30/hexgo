package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tonoy30/hexgo/pkg/http/rest"
)

const (
	PORT = ":8081"
)

func main() {
	fmt.Println("starting server on port " + PORT)
	router := rest.InitHandler()
	log.Fatal(http.ListenAndServe(PORT, router))
}
