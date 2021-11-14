package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transport struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	CreatorName   string             `json:"creatorName"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	CreatedAt   string             `json:"createdAt"`
	Type        string             `json:"type"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Permission string `json:"permission"`
	UserName string `json:"userName"`
}

type TransportGroup struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name     string   `json:"name"`
	UnitsIds []primitive.ObjectID `json:"unitsIds"`
}