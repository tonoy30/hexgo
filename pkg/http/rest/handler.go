package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InitHandler() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/", Index).Methods(http.MethodGet)
	return router
}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}
