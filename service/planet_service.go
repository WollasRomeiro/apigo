package service

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"go-swapi-api/model"
	"go-swapi-api/repository"
)

type PlanetService struct {
	db   *sqlx.DB
	repo *repository.PlanetRepository
}

func NewPlanetService(db *sqlx.DB) *PlanetService {
	return &PlanetService{
		db:   db,
		repo: repository.NewPlanetRepository(db),
	}
}

func (s *PlanetService) GetAndSavePlanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Buscar dados da SWAPI
	resp, err := http.Get("https://swapi.dev/api/planets/" + id)
	if err != nil {
		http.Error(w, "Error fetching data from SWAPI", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "SWAPI responded with an error", resp.StatusCode)
		return
	}

	// Ler o corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}

	var planet model.Planet
	err = json.Unmarshal(body, &planet)
	if err != nil {
		http.Error(w, "Error parsing JSON response", http.StatusInternalServerError)
		return
	}

	// Salvar no banco de dados
	err = s.repo.SavePlanet(&planet)
	if err != nil {
		http.Error(w, "Error saving planet to database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
