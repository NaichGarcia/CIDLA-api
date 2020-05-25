package main

import(
	"context"
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"strconv"
	
	"NaichGarcia/CIDLA-api/helper"
	"NaichGarcia/CIDLA-api/models"
	
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	r := mux.NewRouter()
	
	//rutas de la api
	r.HandleFunc("/test/dummy_collection", getDummies).Methods("GET")
	r.HandleFunc("/test/dummy_collection/{id}", getDummy).Methods("GET")
	
	//designar puerto a usar
	log.Fatal(http.ListenAndServe(":8080", r))
}

//Devuelve todos los dummies
func getDummies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var dummies []models.Dummy_data
	
	collection := helper.ConnectDB()
	
	cur, err := collection.Find(context.TODO(), bson.M{})
	
	if err != nil {
		helper.GetError(err, w)
		return
	}
	
	defer cur.Close(context.TODO())
	
	for cur.Next(context.TODO()) {
		var dummy models.Dummy_data
		err := cur.Decode(&dummy)
		if err != nil {
			log.Fatal(err)
		}
		dummies = append(dummies, dummy)	
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	json.NewEncoder(w).Encode(dummies)
}

func getDummy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dummy models.Dummy_data
	var params = mux.Vars(r)
	
	//id, _ := primitive.ObjectIDFromHex(params["id"])
	id, _ := strconv.ParseInt(params["id"], 10, 16)
	fmt.Println(id)
	
	collection := helper.ConnectDB()
	
	err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&dummy)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	
	json.NewEncoder(w).Encode(dummy)
}