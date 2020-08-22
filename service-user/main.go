package main

import (

	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const Port = 8082

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", userHandler).Methods("GET")
	r.HandleFunc("/user", userPatchHandler).Methods("PATCH")
	log.Printf("Starting on port %d", Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(Port), r))
}
