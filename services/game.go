package services

import (
	"context"
	"fmt"

	"suriyaa.com/cache"
	player "suriyaa.com/models"
)

func StartGame(player *player.Player) error {
	fmt.Println("Came to service startGame")
	return cache.AddPlayer(context.Background(), player)
}

func EndGame(player *player.Player) error {
	return cache.Delete(context.Background(), player)
}

func GetPopularMode(areaCode string) ([]cache.GameModeCount, error) {
	return cache.FindMostFrequentGameMode(context.Background(), areaCode)
}
