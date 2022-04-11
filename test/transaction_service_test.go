package test

import (
	"encoding/json"
	"github.com/ramdanariadi/grocery-go-report/main/helpers"
	"github.com/ramdanariadi/grocery-go-report/main/models"
	"github.com/ramdanariadi/grocery-go-report/main/services"
	"github.com/ramdanariadi/grocery-go-report/main/utils"
	"log"
	"testing"
)

func Test_add_user_collection(t *testing.T) {
	client := utils.GetClient()
	serviceImpl := services.TransactionServiceImpl{Client: client}
	transactionJson := `{"total":3300,"id":"d9c35e9f-b8df-11ec-8c8a-67dc87d2081b","userId":"c47bad3e-b8df-11ec-8c8a-1bc273e5936d","products":[{"total":2,"perUnit":1000,"price":1100,"imageUrl":null,"name":"Broccoli","weight":1500,"id":"a3c02e8c-11d2-11ec-82a8-0242ac130003"}],"Name":"admin"}`
	transaction := models.Transaction{}
	err := json.Unmarshal([]byte(transactionJson), &transaction)
	helpers.LogIfError(err)
	serviceImpl.Add(transaction)
	log.Println(transaction.Name)
}

func Test_get_all_transaction_collection(t *testing.T) {
	client := utils.GetClient()
	serviceImpl := services.TransactionServiceImpl{Client: client}
	serviceImpl.Get()
}
