package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dmitryt/otus-microservices-hw/hw01_k8s_1/internal/config"
	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
	if err != nil {
		fmt.Printf("Error occurred %s", err)
	}
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods("GET")

	fmt.Printf("Starting server at port %d \n", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r))
}
