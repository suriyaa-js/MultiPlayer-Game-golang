package storage

import (
	"context"
	"fmt"
	"log"

	player "suriyaa.com/models" // Update with your actual import path

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var playerCollection *mongo.Collection

// Initialize the MongoDB collection
func InitMongo(collection *mongo.Collection) {
	playerCollection = collection
}

// InsertPlayer inserts a Player record into MongoDB
func InsertPlayer(p *player.Player) error {
	playerDoc := map[string]interface{}{
		"username":  p.Username,
		"area_code": p.AreaCode,
		"mode":      p.Mode.String(),
	}

	_, err := playerCollection.InsertOne(context.Background(), playerDoc)
	return err
}

// GetPlayer retrieves a Player record from MongoDB
func GetPlayer(id string) (*player.Player, error) {
	// primitiveID, err := primitive.ObjectIDFromHex(id)
	var result struct {
		Username string `bson:"username"`
		AreaCode string `bson:"area_code"`
		Mode     int32  `bson:"mode"` // Assuming mode is stored as an integer in MongoDB
	}

	err := playerCollection.FindOne(context.Background(), map[string]interface{}{"id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}

	// Convert mode from int32 to the enum type if needed
	return &player.Player{
		Username: result.Username,
		AreaCode: result.AreaCode,
		Mode:     player.Mode(result.Mode), // Assuming you have a Mode type in your protobuf
	}, nil
}

func SavePlayer(ctx context.Context, newPlayer *player.Player) error {
	fmt.Println("came to DB layer")
	fmt.Println(newPlayer.Username)
	var existingPlayer player.Player
	// bson.M{"username": newPlayer.Username}

	err := playerCollection.FindOne(ctx, map[string]interface{}{"username": newPlayer.Username}).Decode(&existingPlayer)

	if err == nil {
		// If no error, it means the player exists
		return fmt.Errorf("username %s already exists", newPlayer.Username)
	}

	if err != mongo.ErrNoDocuments {
		// If there's an error that isn't "not found", return it
		return err
	}

	// If no player exists, proceed to insert the new player
	_, err = playerCollection.InsertOne(ctx, newPlayer)
	return err

}

func GetPlayerByUsername(ctx context.Context, username string) (*player.Player, error) {
	var result struct {
		Username string `bson:"username"`
		AreaCode string `bson:"area_code"`
		Mode     int32  `bson:"mode"` // Assuming mode is stored as an integer in MongoDB
	}

	err := playerCollection.FindOne(ctx, map[string]interface{}{"username": username}).Decode(&result)
	if err != nil {
		return nil, err
	}

	// Convert mode from int32 to the enum type if needed
	return &player.Player{
		Username: result.Username,
		AreaCode: result.AreaCode,
		Mode:     player.Mode(result.Mode), // Assuming you have a Mode type in your protobuf
	}, nil
}

func UpdatePlayer(ctx context.Context, updatedPlayer *player.Player) (*player.Player, error) {
	filter := bson.M{"username": updatedPlayer.Username}
	update := bson.M{
		"$set": bson.M{
			"mode":     updatedPlayer.Mode,
			"areacode": updatedPlayer.AreaCode,
		},
	}

	_, err := playerCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return nil, err
	}

	log.Println("User updated successfully")
	return updatedPlayer, nil
}
