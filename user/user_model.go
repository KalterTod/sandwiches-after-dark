package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MenuItem struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name   string             `bson:"name,omitempty" json:"name,omitempty"`
	Email  string             `bson:"email,omitempty" json:"email,omitempty"`
	Phone  string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Role   string             `bson:"role,omitempty" json:"role,omitempty"`
	Active *bool              `bson:"active,omitempty" json:"active,omitempty"`
}
