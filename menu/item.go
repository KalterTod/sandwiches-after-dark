package menu

import (
  	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MenuItem struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
    Name     string             `bson:"name,omitempty" json:"name,omitempty"`
    Desc     string             `bson:"desc,omitempty" json:"desc,omitempty"`
    Options  []string           `bson:"options,omitempty" json:"options,omitempty"`
    Active   *bool              `bson:"active,omitempty" json:"active,omitempty"`
}
