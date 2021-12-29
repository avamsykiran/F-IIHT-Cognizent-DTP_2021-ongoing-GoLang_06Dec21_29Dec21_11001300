package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"addressbook-api.cts.com/models"
	"addressbook-api.cts.com/services"
	"github.com/gorilla/mux"
)

func getContactsReqHandler(w http.ResponseWriter, r *http.Request) {
	var contacts []models.Contact = services.GetAllContacts()
	json.NewEncoder(w).Encode(contacts)
}

func getContactByIdReqHandler(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	id, _ := strconv.Atoi(pathVars["id"])

	contact, err := services.GetContactById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(contact)
	}
}

func delContactByIdReqHandler(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	id, _ := strconv.Atoi(pathVars["id"])

	err := services.DeleteContact(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func addContactReqHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var contact models.Contact
	err := json.Unmarshal(reqBody, &contact)
	if err == nil {
		services.AddContact(contact)
		json.NewEncoder(w).Encode(contact)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
	}
}

func main() {
	fmt.Println("REST API - Contact")

	adbRouter := mux.NewRouter().StrictSlash(true)

	adbRouter.HandleFunc("/contacts", getContactsReqHandler).Methods("GET")
	adbRouter.HandleFunc("/contacts", addContactReqHandler).Methods("POST")
	adbRouter.HandleFunc("/contacts/{id}", getContactByIdReqHandler).Methods("GET")
	adbRouter.HandleFunc("/contacts/{id}", delContactByIdReqHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9999", adbRouter))
}
