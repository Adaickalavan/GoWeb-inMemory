package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Person is a struct decsribing its properties
type Person struct {
	Nama       string `json:"nama"`
	Birthday   string `json:"birthday"`
	Occupation string `json:"occupation"`
}

// This variable will act as our in-memory database. Note that any additional
// data entered in a web session will be erased when the server is restarted.
var personList []Person

func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	// Convert the `personList` variable to JSON
	personListBytes, err := json.Marshal(personList)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write JSON list of persons to response
	w.Write(personListBytes)
}

func createPersonHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML form data received in the request
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Extract the field information about the person from the form info
	person := Person{}
	person.Nama = r.Form.Get("nama")
	person.Birthday = r.Form.Get("birthday")
	person.Occupation = r.Form.Get("occupation")

	// Append our existing in-memory database with a the newly received person
	personList = append(personList, person)

	//Redirect to the originating HTML page
	http.Redirect(w, r, "/", http.StatusFound)
}
