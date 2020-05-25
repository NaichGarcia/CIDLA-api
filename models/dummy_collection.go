package models

//import "go.mongodb.org/mongo-driver/bson/primitive"

type Dummy_data struct {
	ID     int    `bson:"id, omitempty"    json:"id, omitempty"`
	Value  int    `bson:"value, omitempty" json:"value, omitempty"`
	Name   string `bson:"name, omitempty"  json:"name, omitempty"`
}
