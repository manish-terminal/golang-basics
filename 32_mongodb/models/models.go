package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Netflix struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Movie   string             `bson:"movie,omitempty" json:"movie,omitempty"`
	Watched bool               `bson:"watched,omitempty" json:"watched,omitempty"`
}

