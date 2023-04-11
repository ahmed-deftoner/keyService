package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ahmed-deftoner/keyService/models"
	"github.com/ahmed-deftoner/keyService/response"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (server *Server) CreateKey(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	Key := models.Key{}
	err = json.Unmarshal(body, &Key)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	Key.Prepare()
	err = Key.ValidateKey()
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	KeyCreated, err := Key.SaveKey(server.DB)

	if err != nil {

		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, KeyCreated.Uid))
	response.JSON(w, http.StatusCreated, KeyCreated)
}

func (server *Server) CreateExchanges(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	Exchange := models.Exchanges{}
	err = json.Unmarshal(body, &Exchange)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = Exchange.ValidateExchange()
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	KeyCreated, err := Exchange.SaveExchange(server.DB)

	if err != nil {

		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, KeyCreated.Id))
	response.JSON(w, http.StatusCreated, KeyCreated)
}

func (server *Server) GetKeys(w http.ResponseWriter, r *http.Request) {

	Key := models.Key{}

	Keys, err := Key.FindAllKeys(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, Keys)
}

func (server *Server) GetExchanges(w http.ResponseWriter, r *http.Request) {

	Exchange := models.Exchanges{}

	Exchanges, err := Exchange.FindAllExchanges(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, Exchanges)
}

func (server *Server) GetKey(w http.ResponseWriter, r *http.Request) {

	kid := mux.Vars(r)["id"] //grab the id
	new_kid, err := uuid.Parse(kid)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("Invalid Key ID"))
		return
	}
	Key := models.Key{}
	KeyGotten, err := Key.FindKeyById(server.DB, new_kid)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	response.JSON(w, http.StatusOK, KeyGotten)
}
