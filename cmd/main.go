package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tonoy30/hexgo/pkg/http/rest"
	"github.com/tonoy30/hexgo/tool"
)

const (
	PORT = ":8081"
)

func init() {
	tool.SetUpConfiguration()
}
func main() {
	fmt.Println("starting server on port " + PORT)
	router := rest.InitHandler()
	log.Fatal(http.ListenAndServe(PORT, router))
}
