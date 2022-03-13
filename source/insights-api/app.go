package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	models "github.com/malekatwiz/miners-stats-insights/models"
	minerstats "github.com/malekatwiz/miners-stats-insights/services"
	"gopkg.in/yaml.v2"
)

func main() {
	cfg := loadConfigs()

	router := mux.NewRouter()
	router.HandleFunc("/api/status", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Service is healthy."))
	}).Methods("GET")
	router.HandleFunc("/api/miners/flux", func(w http.ResponseWriter, r *http.Request) {
		content, e := json.Marshal(minerstats.GetOperationWorkers(cfg.Operations[0].Wallet))
		if e == nil {
			w.WriteHeader(http.StatusOK)
			w.Write(content)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

	}).Methods("GET")
	log.Fatalf("failed to start sever. '%s'", http.ListenAndServe(":5000", handlers.CORS()(router)).Error())
}

func loadConfigs() models.Config {
	c, e := os.Open("config.yaml")
	if e != nil {
		return models.Config{}
	}
	var cfg models.Config
	defer c.Close()
	decoder := yaml.NewDecoder(c)
	_ = decoder.Decode(&cfg)
	return cfg
}
