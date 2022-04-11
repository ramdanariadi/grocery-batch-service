package services

import (
	"context"
	"github.com/ramdanariadi/grocery-go-report/main/helpers"
	"github.com/ramdanariadi/grocery-go-report/main/models"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type UserService interface {
	Get()
	Add(transaction models.User)
	Update()
	Delete()
}

type UserServiceImpl struct {
	Client *mongo.Client
}

func (service UserServiceImpl) Get() {
	//TODO implement me
	panic("implement me")
}

func (service UserServiceImpl) Add(transaction models.User) {
	collection := service.Client.Database("tunasReport").Collection("users")
	result, err := collection.InsertOne(context.Background(), transaction)
	helpers.LogIfError(err)
	log.Printf("inserted id : %s", result.InsertedID)
}

func (service UserServiceImpl) Update() {
	//TODO implement me
	panic("implement me")
}

func (service UserServiceImpl) Delete() {
	//TODO implement me
	panic("implement me")
}
