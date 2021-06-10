package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jefersonvinicius/fullcycle-course-hexagonal-architecture/products-application/application"
	"github.com/urfave/negroni"
)

func MakeProductHandlers(router *mux.Router, neg *negroni.Negroni, service application.ProductServiceInterface) {
	router.Handle("/product/{id}", neg.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(request)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			writer.Write([]byte(jsonError(err.Error())))
			return
		}

		err = json.NewEncoder(writer).Encode(product)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			return
		}
	})
}
