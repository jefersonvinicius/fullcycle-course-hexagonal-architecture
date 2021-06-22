package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jefersonvinicius/fullcycle-course-hexagonal-architecture/products-application/adapters/dto"
	"github.com/jefersonvinicius/fullcycle-course-hexagonal-architecture/products-application/application"
	"github.com/urfave/negroni"
)

func MakeProductHandlers(router *mux.Router, neg *negroni.Negroni, service application.ProductServiceInterface) {
	router.Handle("/product/{id}", neg.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	router.Handle("/products", neg.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

	router.Handle("/products/{id}/enable", neg.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("GET", "OPTIONS")

	router.Handle("/products/{id}/disable", neg.With(
		negroni.Wrap(disableProduct(service)),
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

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var productDto dto.Product
		err := json.NewDecoder(request.Body).Decode(&productDto)
		if err != nil {
			writeError(writer, err.Error())
			return
		}

		product, err := service.Create(productDto.Name, productDto.Price)
		if err != nil {
			writeError(writer, err.Error())
			return
		}

		err = json.NewEncoder(writer).Encode(product)
		if err != nil {
			writeError(writer, err.Error())
			return
		}
	})
}

func enableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		id := mux.Vars(request)["id"]
		product, err := service.Get(id)
		if err != nil {
			writeError(writer, err.Error())
			return
		}
		result, err := service.Enable(product)
		if err != nil {
			writeError(writer, err.Error())
			return
		}
		json.NewEncoder(writer).Encode(result)
	})
}

func disableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		id := mux.Vars(request)["id"]
		product, err := service.Get(id)
		if err != nil {
			writeError(writer, err.Error())
			return
		}
		result, err := service.Disable(product)
		if err != nil {
			writeError(writer, err.Error())
			return
		}
		json.NewEncoder(writer).Encode(result)
	})
}

func writeError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(jsonError(message)))
}
