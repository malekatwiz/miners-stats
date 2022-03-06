package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/status", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Service is healthy."))
	}).Methods("GET")

	log.Fatalf("failed to start sever. '%s'", http.ListenAndServe(":5000", handlers.CORS()(router)).Error())
}
