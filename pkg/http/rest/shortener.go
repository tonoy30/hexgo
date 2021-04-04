package rest

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/tonoy30/hexgo/pkg/business"
	"github.com/tonoy30/hexgo/pkg/serializeres"
	js "github.com/tonoy30/hexgo/pkg/serializeres/json"
	ms "github.com/tonoy30/hexgo/pkg/serializeres/msgpack"
	"github.com/tonoy30/hexgo/pkg/services"
)

type ShortenerHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	service services.ShortenerService
}

func NewHandler(service services.ShortenerService) ShortenerHandler {
	return &handler{
		service: service,
	}
}
func setUpResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) serializer(contentType string) serializeres.ShortenerSerializer {
	if contentType == "application/x-msgpack" {
		return &ms.Redirect{}
	}
	return &js.Redirect{}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
	redirect, err := h.service.Find(code)
	if err != nil {
		if errors.Cause(err) == business.ErrRedirectNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, redirect.URL, http.StatusMovedPermanently)

}
func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	short, err := h.serializer(contentType).Decode(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = h.service.Store(short)
	if err != nil {
		if errors.Cause(err) == business.ErrRedirectInvalid {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseBody, err := h.serializer(contentType).Encode(short)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setUpResponse(w, contentType, responseBody, http.StatusCreated)
}
