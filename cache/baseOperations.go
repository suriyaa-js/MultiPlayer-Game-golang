package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

func flushAll(ctx context.Context) error {
	return rdb.FlushAll(ctx).Err()
}

func PrintAllKeyValues(ctx context.Context) error {
	// Get all keys
	keys, err := rdb.Keys(ctx, "*").Result()
	if err != nil {
		return err
	}

	// Iterate over keys and print their values
	for _, key := range keys {
		value, err := rdb.Get(ctx, key).Result()
		if err != nil {
			// Handle case where key does not exist
			if err == redis.Nil {
				fmt.Printf("Key: %s, Value: (key does not exist)\n", key)
			} else {
				return err
			}
		} else {
			fmt.Printf("Key: %s, Value: %s\n", key, value)
		}
		fmt.Println()
	}

	return nil
}

func deletePlayer(ctx context.Context, areaCode, playerName string) error {
	key := fmt.Sprintf("area:%s:players", areaCode)

	// Retrieve existing players
	existingData, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	// Unmarshal existing players
	var players []Player
	err = json.Unmarshal([]byte(existingData), &players)
	if err != nil {
		return err
	}

	if len(players) <= 0 {
		return fmt.Errorf("there is no such game to end")
	}

	// Filter out the player to delete
	updatedPlayers := []Player{}
	for _, player := range players {
		if player.PlayerName != playerName {
			updatedPlayers = append(updatedPlayers, player)
		}
	}

	// Marshal the updated list back to JSON
	updatedData, err := json.Marshal(updatedPlayers)
	if err != nil {
		return err
	}

	// Store the updated list back in Redis
	return rdb.Set(ctx, key, updatedData, 0).Err()
}

// Function to store updated players back in Redis
func storePlayers(ctx context.Context, key string, players []Player) error {
	// Marshal the updated list back to JSON
	updatedData, err := json.Marshal(players)
	if err != nil {
		return err
	}

	// Store the updated list in Redis
	return rdb.Set(ctx, key, updatedData, 0).Err()
}
