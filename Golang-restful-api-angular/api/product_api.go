package api

import (
	"encoding/json"
	"golang-restful-api-angular/config"
	"golang-restful-api-angular/entities"
	"golang-restful-api-angular/models"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func FindAll(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModels{
			Db: db,
		}

		products, err := productModel.FindAll()
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		} else {
			respondWithJSON(w, http.StatusOK, products)
		}
	}
}

func FindById(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModels{
			Db: db,
		}
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := productModel.FindByID(id)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		} else {
			respondWithJSON(w, http.StatusOK, product)
		}
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		productModel := models.ProductModels{
			Db: db,
		}
		var product entities.Product
		product.Id = bson.NewObjectId()
		json.NewDecoder(r.Body).Decode(&product)
		err := productModel.Create(&product)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		} else {
			respondWithJSON(w, http.StatusOK, product)
		}
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModels{
			Db: db,
		}
		vars := mux.Vars(r)
		id := vars["id"]
		err := productModel.Delete(id)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		} else {
			respondWithJSON(w, http.StatusOK, entities.Product{})
		}
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	} else {
		productModel := models.ProductModels{
			Db: db,
		}
		var product entities.Product
		json.NewDecoder(r.Body).Decode(&product)
		err := productModel.Update(&product)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		} else {
			respondWithJSON(w, http.StatusOK, product)
		}
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg}) //map có key và value là string,phần tử đầu key là error,value là msg
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
