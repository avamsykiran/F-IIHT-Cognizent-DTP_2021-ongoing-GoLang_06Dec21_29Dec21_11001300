package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	
	"addressbook.cts.com/models"
	"addressbook.cts.com/services"
)

//create a template handler
var tmpls = template.Must(template.ParseGlob("templates/*"))

func home(w http.ResponseWriter, r *http.Request) {
	tmpls.ExecuteTemplate(w, "Home", services.GetAllContacts())
}

func details(w http.ResponseWriter, r *http.Request) {

	cid, _ := strconv.Atoi(r.URL.Query().Get("id"))

	contact, err := services.GetContactById(cid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		tmpls.ExecuteTemplate(w, "Details", contact)
	}
}

func deleteContact(w http.ResponseWriter, r *http.Request) {

	cid, _ := strconv.Atoi(r.URL.Query().Get("id"))

	err := services.DeleteContact(cid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		tmpls.ExecuteTemplate(w, "Home", services.GetAllContacts())
	}
}

func newContact(w http.ResponseWriter, r *http.Request) {
	tmpls.ExecuteTemplate(w, "ContactForm", nil)
}

func addContact(w http.ResponseWriter, r *http.Request) {

	contactId, _ := strconv.Atoi(r.FormValue("ContactId"))

	contact := models.Contact{contactId, r.FormValue("FirstName"), r.FormValue("LastName"), r.FormValue("Mobile"), r.FormValue("AlternateMobile"), r.FormValue("MailId")}

	services.AddContact(contact)

	tmpls.ExecuteTemplate(w, "Home", services.GetAllContacts())
}

func main() {
	const port string = ":4545"

	http.HandleFunc("/", home)

	http.HandleFunc("/view", details)
	http.HandleFunc("/new", newContact)
	http.HandleFunc("/add", addContact)
	http.HandleFunc("/delete", deleteContact)

	//http.HandleFunc("/api/contacts", getAllContacts)

	fmt.Println("Server started at port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
