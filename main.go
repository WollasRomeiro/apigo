package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go-swapi-api/service"
)

func main() {
	// Database connection setup
	db, err := sqlx.Open("postgres", "user=username dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize router and service
	r := mux.NewRouter()
	planetService := service.NewPlanetService(db)

	r.HandleFunc("/planets/{id:[0-9]+}", planetService.GetAndSavePlanet).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
