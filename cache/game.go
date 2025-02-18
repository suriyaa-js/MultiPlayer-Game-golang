package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	player "suriyaa.com/models"
)

type Player struct {
	PlayerName string `json:"player_name"`
	GameMode   string `json:"game_mode"`
}

type GameModeCount struct {
	Mode  string `json:"mode"`
	Count int    `json:"count"`
}

func Random() {
	// newPlayer := Player{PlayerName: "Alice", GameMode: "Squad"}
	// areaCode := "234"

	// ctx := context.Background()

	// // Add the new player to the existing list
	// err := addPlayer(ctx, areaCode, newPlayer)
	// if err != nil {
	// 	log.Fatalf("Could not add player: %v", err)
	// }

	// retrievedPlayers, err := retrievePlayers(ctx, areaCode)
	// if err != nil {
	// 	log.Fatalf("Could not retrieve players: %v", err)
	// }

	// fmt.Printf("Players in area %s: %+v\n", areaCode, retrievedPlayers)
	// delete()
	// printAllKeyValues(context.Background())
	//flushAll(context.Background())
	//fmt.Println("after deleting all values")
	//printAllKeyValues(context.Background())

}

// Function to add a player to the existing list in Redis
func AddPlayer(ctx context.Context, gameData *player.Player) error {

	areaCode := gameData.AreaCode
	key := fmt.Sprintf("area:%s:players", areaCode)

	if gameData.GetUsername() == "" || gameData.GetAreaCode() == "" || !isValidGameMode(gameData.Mode) {
		return fmt.Errorf("AreaCode and Username and Mode should not be empty")
	}

	// Retrieve existing players
	players, _ := getExistingPlayers(ctx, key)

	// Check if the player already exists
	if players != nil && playerExists(players, gameData.GetUsername()) {
		fmt.Printf("Player %s is already in the list.\n", gameData.GetUsername())
		return fmt.Errorf("already you are in the game end and re-start the game") // Exit without adding the player
	}

	// Create a new player
	newPlayer := Player{
		PlayerName: gameData.GetUsername(),
		GameMode:   gameData.GetMode().String(),
	}
	players = append(players, newPlayer)

	// Store the updated list in Redis
	return storePlayers(ctx, key, players)
}

// Function to retrieve players from Redis
func retrievePlayers(ctx context.Context, areaCode string) ([]Player, error) {
	key := fmt.Sprintf("area:%s:players", areaCode)

	// Get the JSON string from Redis
	jsonData, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON back to Player slice
	var players []Player
	err = json.Unmarshal([]byte(jsonData), &players)
	if err != nil {
		return nil, err
	}

	// fmt.Println(string(jsonData))
	return players, nil
}

func Delete(ctx context.Context, gameData *player.Player) error {

	if gameData.GetUsername() == "" || gameData.GetAreaCode() == "" {
		return fmt.Errorf("AreaCode and Username should not be empty")
	}

	players, err := retrievePlayers(ctx, gameData.GetAreaCode())
	if err != nil {
		return fmt.Errorf("no players found in this location")

	}

	// Delete a player by name
	err = deletePlayer(ctx, gameData.GetAreaCode(), gameData.GetUsername())
	if err != nil {
		// log.Fatalf("Could not delete player: %v", err)
		return err
	}

	fmt.Printf("Players in area %s after deletion: %+v\n", gameData.GetAreaCode(), players)
	return nil
}

func FindMostFrequentGameMode(ctx context.Context, areaCode string) ([]GameModeCount, error) {
	if areaCode == "" {
		return nil, fmt.Errorf("areaCode should not be empty")
	}
	key := fmt.Sprintf("area:%s:players", areaCode)

	// Retrieve the JSON string from Redis
	jsonData, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if jsonData == "" {
		return nil, fmt.Errorf("no players in that location currently or no such location found")
	}

	// Unmarshal JSON back to Player slice
	var players []Player
	err = json.Unmarshal([]byte(jsonData), &players)
	if err != nil {
		return nil, err
	}

	// Count occurrences of each game mode
	modeCount := make(map[string]int)
	for _, player := range players {
		modeCount[player.GameMode]++
	}

	// Find the most frequent game mode
	var results []GameModeCount
	for mode, count := range modeCount {
		results = append(results, GameModeCount{Mode: mode, Count: count})
	}

	return results, nil
}

// Function to check if a player already exists in the list
func playerExists(players []Player, username string) bool {
	for _, player := range players {
		if player.PlayerName == username {
			return true
		}
	}
	return false
}

// Function to retrieve existing players from Redis
func getExistingPlayers(ctx context.Context, key string) ([]Player, error) {
	existingData, err := rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("came to err in Cache")
		if err == redis.Nil {
			return []Player{}, nil // Return an empty slice if the key does not exist
		}
		return nil, err
	}

	var players []Player
	err = json.Unmarshal([]byte(existingData), &players)
	if err != nil {
		return nil, err
	}

	return players, nil
}

func isValidGameMode(mode player.Mode) bool {
	if mode == player.Mode_SOLO || mode == player.Mode_DUO || mode == player.Mode_TRIO || mode == player.Mode_SQUAD {
		return true
	}
	return false
}
