package services

import (
	"register/src/models/schemas"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func Register(req schemas.RegisterSchema) error {

	return nil
}
