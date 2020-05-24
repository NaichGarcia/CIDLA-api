package models

//import "go.mongodb.org/mongo-driver/bson/primitive"

type Dummy_data struct {
	id     int16  `json:"id, omitempty"    bson:"id, omitempty"`
	value  int16  `json:"value, omitempty" bson:"value, omitempty"`
	name   string `json:"name, omitempty"  bson:"name, omitempty"`
}
