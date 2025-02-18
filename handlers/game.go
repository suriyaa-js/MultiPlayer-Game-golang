package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"suriyaa.com/cache"
	player "suriyaa.com/models"
	"suriyaa.com/services"
)

type StringResponse struct {
	Message string `json:"message"`
}

type gameModeCount struct {
	Mode  string `json:"mode"`
	Count int    `json:"count"`
}
type GameModeCounts []gameModeCount

// @Summary Fetches the modes and the count of players for the mode
// @Description GetPopularModes fetches popular multiplayer modes based on the area code.
// @Tags players
// @Accept json
// @Produce json
// @Param areaCode query string true "areaCode"
// @Success 200 {object} GameModeCounts
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/mode-statistics [get]
// GetPopularModes fetches popular multiplayer modes based on the area code.
func GetPopularModes(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	areaCode := queryParams.Get("areaCode")

	if areaCode == "" {
		fmt.Fprintf(w, "Parameter 'param' not found")
		return
	}

	response, err := services.GetPopularMode(areaCode)
	fmt.Println(response)
	if err != nil {
		http.Error(w, "No players in that location or No such locations "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// @Summary start the game by adding the player mode to that areaCode
// @Description StartGame set the Username and mode for the given areaCode(mandatory).
// @Tags players
// @Accept json
// @Produce json
// @Param player body player.Player true "Player data"
// @Success 200 {object} StringResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/start-game [post]
// StartGame store the player with the selected mode in the area.
func StartGame(w http.ResponseWriter, r *http.Request) {
	var playerData player.Player
	if err := json.NewDecoder(r.Body).Decode(&playerData); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := services.StartGame(&playerData)
	if err != nil {
		fmt.Println("Came after service startGame", playerData)
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Game Begins🔥")
}

// @Summary EndGame the game by removing the player from cache for that areaCode
// @Description EndGame the game by taking the userName and areaCode.
// @Tags players
// @Accept json
// @Produce json
// @Param player body player.Player true "Player data"
// @Success 200 {object} StringResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/end-game [post]
// StartGame store the player with the selected mode in the area.
// EndGame Will delete the player from that area.
func EndGame(w http.ResponseWriter, r *http.Request) {
	var playerData player.Player
	if err := json.NewDecoder(r.Body).Decode(&playerData); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := services.EndGame(&playerData)
	if err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Game Ends🔚")
}

// PrintAll it will log the current cached data all keys and their values
func PrintAll(w http.ResponseWriter, r *http.Request) {
	err := cache.PrintAllKeyValues(context.Background())
	if err != nil {
		http.Error(w, "failed to printall: "+err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("response is logged")

}
