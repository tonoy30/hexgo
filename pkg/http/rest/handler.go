package rest

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/tonoy30/hexgo/pkg/business"
	"github.com/tonoy30/hexgo/pkg/http/middlewares"
	"github.com/tonoy30/hexgo/pkg/repo"
	"github.com/tonoy30/hexgo/pkg/repo/mongo"
	"github.com/tonoy30/hexgo/pkg/repo/redis"
)

func InitHandler() *mux.Router {
	repo := chooseRepo()
	service := business.NewShortenerService(repo)
	handler := NewHandler(service)

	router := mux.NewRouter()

	router.Use(middlewares.RealIP)
	router.Use(middlewares.Logger)

	router.HandleFunc("/", Index).Methods(http.MethodGet)
	router.HandleFunc("/api/shortener", handler.Post).Methods(http.MethodOptions, http.MethodPost)
	router.HandleFunc("/api/shortener/{code}", handler.Get).Methods(http.MethodOptions, http.MethodGet)
	router.Use(mux.CORSMethodMiddleware(router))

	return router
}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}
func chooseRepo() repo.ShortenerRepository {
	switch viper.GetString("DB.TYPE") {
	case "redis":
		redisURL := viper.GetString("DB.REDIS_URL")
		repo, err := redis.NewRedisRepository(redisURL)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case "mongo":
		mongoURL := viper.GetString("DB.MONGO_URL")
		mongoDB := viper.GetString("DB.MONGO_DB")
		mongoTimeout, _ := strconv.Atoi(viper.GetString("DB.MONGO_TIMEOUT"))
		repo, err := mongo.NewMongoRepository(mongoURL, mongoDB, mongoTimeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
