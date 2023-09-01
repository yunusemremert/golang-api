package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id      primitive.ObjectID `bson:"_id" json:"id"`
	Title   string             `json:"title,omitempty"`
	Content string             `json:"content,omitempty"`
}
