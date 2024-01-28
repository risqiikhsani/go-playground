package inputOutput

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// User represents the structure of a user document
type User struct {
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

func InputData2(client *mongo.Client) {
	fmt.Println("Enter name:")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Enter email:")
	var email string
	fmt.Scanln(&email)

	// Get the users collection
	collection := client.Database("mydatabase").Collection("users")

	// Create a new User instance
	user := User{
		Name:  name,
		Email: email,
	}

	// Insert the user document into the collection
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data inserted successfully!")
}
