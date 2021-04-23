package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

func ConnectDB() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:admin@rest-api.p3kvu.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

	client, getErr := mongo.Connect(context.TODO(), clientOptions)
	if getErr != nil {
		log.Fatal(getErr)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("pokemon").Collection("lists")
	return collection
}

func GetError(err error, w http.ResponseWriter) {
	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
