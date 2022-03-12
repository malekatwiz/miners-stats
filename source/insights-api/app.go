package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	minerstats "github.com/malekatwiz/miners-stats-insights/services"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/status", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Service is healthy."))
	}).Methods("GET")
	router.HandleFunc("/api/miners/flux", func(w http.ResponseWriter, r *http.Request) {
		content, e := json.Marshal(minerstats.GetOperationWorkers("flux"))
		if e == nil {
			w.WriteHeader(http.StatusOK)
			w.Write(content)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

	}).Methods("GET")
	log.Fatalf("failed to start sever. '%s'", http.ListenAndServe(":5000", handlers.CORS()(router)).Error())
}
