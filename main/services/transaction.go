package services

import (
	"context"
	"github.com/ramdanariadi/grocery-go-report/main/helpers"
	"github.com/ramdanariadi/grocery-go-report/main/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type TransactionService interface {
	Get()
	Add(transaction models.Transaction)
	Update()
	Delete()
}

type TransactionServiceImpl struct {
	Client *mongo.Client
}

func NewTransactionServiceImpl(client *mongo.Client) *TransactionServiceImpl {
	return &TransactionServiceImpl{Client: client}
}

func (service TransactionServiceImpl) Get() {
	collection := service.Client.Database("tunasReport").Collection("transactions")
	cursor, err := collection.Find(context.Background(), bson.M{})
	helpers.LogIfError(err)

	for cursor.Next(context.TODO()) {
		transactionTmp := models.Transaction{}
		err := cursor.Decode(&transactionTmp)
		helpers.LogIfError(err)
		log.Printf("%s - %s", transactionTmp.Id, transactionTmp.Name)
	}
}

func (service TransactionServiceImpl) Add(transaction models.Transaction) {
	collection := service.Client.Database("tunasReport").Collection("transactions")
	result, err := collection.InsertOne(context.Background(), transaction)
	helpers.LogIfError(err)
	log.Printf("inserted id : %s", result.InsertedID)
}

func (service TransactionServiceImpl) Update() {
	//TODO implement me
	panic("implement me")
}

func (service TransactionServiceImpl) Delete() {
	//TODO implement me
	panic("implement me")
}
