package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID          string `json:"userid,omitempty" bson:"userid,omitempty"`
	Caption         string `json:"caption,omitempty" bson:"caption,omitempty"`
	ImageURL        string `json:"imageurl,omitempty" bson:"imageurl,omitempty"`
	PostedTimestamp time.Duration `json:"name,omitempty" bson:"name,omitempty"`
}