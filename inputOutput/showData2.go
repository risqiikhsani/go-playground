package inputOutput

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// User represents the structure of a user document
type UserData struct {
	ID    string `bson:"_id"`
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

func ShowData2(client *mongo.Client) {
	// Get the users collection
	collection := client.Database("mydatabase").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result....
		var user UserData
		if err := cur.Decode(&user); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %s, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
