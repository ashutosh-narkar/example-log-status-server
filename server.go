// Decision Log and Status server

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func logEvent(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	gr, err := gzip.NewReader(bytes.NewReader(body))
	if err != nil {
		log.Printf("Error reading gzip body: %v", err)
		http.Error(w, "can't read gzip body", http.StatusBadRequest)
		return
	}

	var events interface{}
	if err := json.NewDecoder(gr).Decode(&events); err != nil {
		log.Printf("Error decoding body: %v", err)
		http.Error(w, "error decoding body", http.StatusBadRequest)
		return
	}

	gr.Close()

	bs, err := json.MarshalIndent(events, "", " ")
	if err != nil {
		log.Printf("Error marshal indent: %v", err)
		http.Error(w, "error marshal indent", http.StatusInternalServerError)
		return
	}

	log.Printf("Log Event: %+v\n\n", string(bs))
	w.WriteHeader(http.StatusOK)
}

func logStatus(w http.ResponseWriter, r *http.Request) {
	var status interface{}

	if err := json.NewDecoder(r.Body).Decode(&status); err != nil {
		log.Printf("Error decoding body: %v", err)
		http.Error(w, "error decoding body", http.StatusBadRequest)
		return
	}

	bs, err := json.MarshalIndent(status, "", " ")
	if err != nil {
		log.Printf("Error marshal indent: %v", err)
		http.Error(w, "error marshal indent", http.StatusInternalServerError)
		return
	}

	log.Printf("Status is:\n %+v\n\n", string(bs))
	w.WriteHeader(http.StatusOK)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome to the HomePage!")
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/logs", logEvent).Methods("POST")
	router.HandleFunc("/status", logStatus).Methods("POST")
	router.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8001", router))
}
