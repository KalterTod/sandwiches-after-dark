package menu

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MenuItem struct {
  ID            primitive.ObjectID      `bson:"_id"`
  name          string                  `bson:"name"`
  desc          string                  `bson:"desc"`
  options       string                  `bson:"options"`
  active        bool                    `bson:"active"`
}
