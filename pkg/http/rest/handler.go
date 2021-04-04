package rest

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tonoy30/hexgo/pkg/business"
	"github.com/tonoy30/hexgo/pkg/http/middlewares"
	"github.com/tonoy30/hexgo/pkg/repo"
	"github.com/tonoy30/hexgo/pkg/repo/mongo"
	"github.com/tonoy30/hexgo/pkg/repo/redis"
	"github.com/tonoy30/hexgo/tool"
)

func InitHandler() *mux.Router {
	repo := chooseRepo()
	service := business.NewShortenerService(repo)
	handler := NewHandler(service)

	router := mux.NewRouter()
	router.Use(middlewares.RealIP)
	router.Use(middlewares.Logger)
	router.HandleFunc("/", Index).Methods(http.MethodGet)
	router.HandleFunc("/api/shortener", handler.Get).Methods(http.MethodGet)
	router.HandleFunc("/api/shortener/{code}", handler.Post).Methods(http.MethodPost)
	return router
}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}
func chooseRepo() repo.ShortenerRepository {
	switch tool.GetConfigValue("DB:TYPE") {
	case "redis":
		redisURL := tool.GetConfigValue("DB:REDIS_URL")
		repo, err := redis.NewRedisRepository(redisURL)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case "mongo":
		mongoURL := tool.GetConfigValue("DB:MONGO_URL")
		mongoDB := tool.GetConfigValue("DB:MONGO_DB")
		mongoTimeout, _ := strconv.Atoi(tool.GetConfigValue("DB:MONGO_TIMEOUT"))
		repo, err := mongo.NewMongoRepository(mongoURL, mongoDB, mongoTimeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
