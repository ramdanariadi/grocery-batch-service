package main

import (
	"encoding/json"
	"github.com/ramdanariadi/grocery-go-report/main/helpers"
	"github.com/ramdanariadi/grocery-go-report/main/models"
	"github.com/ramdanariadi/grocery-go-report/main/services"
	"github.com/ramdanariadi/grocery-go-report/main/utils"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	connection, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	helpers.LogIfError(err)

	channel, err := connection.Channel()
	helpers.LogIfError(err)

	queue, err := channel.QueueDeclare(
		"transaction",
		true,
		false,
		false,
		false,
		nil,
	)
	helpers.LogIfError(err)

	messages, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	helpers.LogIfError(err)

	forever := make(chan bool)
	client := utils.GetClient()
	serviceImpl := services.TransactionServiceImpl{Client: client}
	go func() {
		for msg := range messages {
			log.Printf("%s", msg.Body)
			transactionTmp := models.Transaction{}
			err := json.Unmarshal(msg.Body, &transactionTmp)
			if err != nil {
				log.Printf("unmarshall failed")
			}
			log.Printf(transactionTmp.Name)
			go serviceImpl.Add(transactionTmp)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
