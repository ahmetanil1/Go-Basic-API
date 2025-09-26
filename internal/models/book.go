package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id, omitempty"`
	Title       string             `json:"title"  bson:"title"`
	Author      string             `json:"author" bson:"author"`
	PublishedAt time.Time          `json:"published_at" bson:"published_at"`
	Quantity    int                `json:"quantity" bson:"quantity"`
}
