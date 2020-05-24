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

//mas tarde podria pasar como param la colleccion a utilizar
func ConnectDB() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	
	client, err := mongo.Connect(context.TODO(), clientOptions)
	
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Conexion exitosa")
	
	collection := client.Database("dummy").Collection("dummy_test")
	
	return collection
}

type ErrorResponse struct {
	ErrorMessage string `json:"message"`
	StatusCode   int    `json:"status"`
}

//prepara el error
func GetError(err error, w http.ResponseWriter) {
	log.Fatal(err.Error())
	var response = ErrorResponse {
		ErrorMessage: err.Error(),
		StatusCode: http.StatusInternalServerError,
	}
	
	message, _ := json.Marshal(response)
	
	w.WriteHeader(response.StatusCode)
	w.Write(message)
}


