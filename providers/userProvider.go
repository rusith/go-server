package providers

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/rusith/go-server/models"
)

type UserProvider struct {}

func (provider UserProvider) GetUser(id string) *models.User  {
	client, err := mongo.NewClient("mongodb://foo:bar@localhost:27017")
	if err != nil {
		panic(err)
	}

	client.Database("test").Collection("users")
	return &models.User{}
}