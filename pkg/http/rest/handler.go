package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tonoy30/hexgo/pkg/http/middlewares"
)

func InitHandler() *mux.Router {
	router := mux.NewRouter()
	router.Use(middlewares.RealIP)
	router.Use(middlewares.Logger)
	router.HandleFunc("/", Index).Methods(http.MethodGet)
	router.HandleFunc("/api/shortener", Index).Methods(http.MethodGet)
	router.HandleFunc("/api/shortener/{code}", Index).Methods(http.MethodPost)
	return router
}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}
