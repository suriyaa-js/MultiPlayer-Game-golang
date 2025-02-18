package handlers

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	player "suriyaa.com/models" // Adjust with your actual module name
	"suriyaa.com/services"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// @Summary Create a new player on a persistant DB just to store permanently
// @Description Create a new player with the provided details
// @Tags players
// @Accept json
// @Produce json
// @Param player body player.Player true "Player data"
// @Success 200 {object} player.Player
// @Failure 400 {object} ErrorResponse
// @Router /api/v1/create-player [post]
// CreatePlayer create and store a new player in persistant DB.
func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var newPlayer player.Player

	if err := json.NewDecoder(r.Body).Decode(&newPlayer); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if newPlayer.Username == "" || newPlayer.AreaCode == "" {
		http.Error(w, "Missing required fields name and preferred/primary areacode is required", http.StatusBadRequest)
		return
	}

	err := services.SavePlayer(&newPlayer)
	if err != nil {
		http.Error(w, "player already exists", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&newPlayer)
}

// @Summary update a exisitng player on a permanent DB
// @Description update a player with the provided details areaCode is mandatory
// @Tags players
// @Accept json
// @Produce json
// @Param player body player.Player true "Player data"
// @Success 200 {object} player.Player
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/update-player [post]
// UpdatePlayer will update the player's detail in persistant DB
func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	var player player.Player

	// Decode the request body into the Player struct
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Optionally, validate the input
	if player.Username == "" || player.AreaCode == "" {
		http.Error(w, "Missing required fields areaCode is needed", http.StatusBadRequest)
		return
	}

	// Save the player to MongoDB
	updatedPlayer, err := services.UpdatePlayer(&player)
	if player.Username == "" || player.AreaCode == "" {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Player not found", http.StatusNotFound)
		} else {
			http.Error(w, "Could not retrieve player", http.StatusNotFound)
		}
		return
	}
	// Respond with the player details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPlayer)

}

// @Summary get a exisitng player from a permanent DB
// @Description get a player with the provided userName
// @Tags players
// @Accept json
// @Produce json
// @Param username query string true "Username of the player"
// @Success 200 {object} player.Player
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/get-player [get]
// GetPlayer will get the player's detail from persistant DB
func GetPlayer(w http.ResponseWriter, r *http.Request) {
	// Extract the username from the URL parameters
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	// Fetch the player from the database
	player, err := services.GetPlayer(username)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Player not found", http.StatusNotFound)
		} else {
			http.Error(w, "Could not retrieve player/No such pla", http.StatusNotFound)
		}
		return
	}

	// Respond with the player details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(player)
}
