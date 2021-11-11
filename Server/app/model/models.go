package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transport struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	CreatorID   string             `json:"creatorId"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	CreatedAt   string             `json:"createdAt"`
}
