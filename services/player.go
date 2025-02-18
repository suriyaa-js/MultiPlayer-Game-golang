package services

import (
	"context"
	"time"

	player "suriyaa.com/models" // Adjust with your actual module name
	"suriyaa.com/storage"
)

// SavePlayer saves a player to the MongoDB collection
func SavePlayer(newPlayer *player.Player) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := storage.SavePlayer(ctx, newPlayer)
	if err != nil {
		return err
	}
	return nil
}

func GetPlayer(username string) (*player.Player, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	player, err := storage.GetPlayerByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return player, err
}

func UpdatePlayer(player *player.Player) (*player.Player, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := storage.GetPlayerByUsername(ctx, player.Username)
	if err != nil {
		return nil, err
	}

	updatedPlayer, err := storage.UpdatePlayer(ctx, player)
	if err != nil {
		return nil, err
	}
	return updatedPlayer, nil

}
