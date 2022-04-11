package utils

import (
	"context"
	"github.com/ramdanariadi/grocery-go-report/main/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var ctx = context.Background()
var client *mongo.Client

func GetClient() *mongo.Client {
	if client != nil {
		log.Println("client is not nil")
		return client
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost").
		SetAuth(options.Credential{Username: "root", Password: "example"})

	connect, err := mongo.Connect(ctx, clientOptions)
	helpers.LogIfError(err)

	client = connect
	return client
}
