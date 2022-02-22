package db_model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
    ID      primitive.ObjectID   `json:"_id" bson:"_id"`
	Text    string   `json:"text" bson:"text"`
	Done    bool     `json:"done" bson:"done"`
    CreatedBy string `json:"created_by" bson:"created_by"`
    CreatedAt time.Time `json:"created_at" bson:"created_at"`
}