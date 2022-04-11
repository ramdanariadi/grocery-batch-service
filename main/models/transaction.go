package models

type Transaction struct {
	Id       string    `bson:"id"`
	UserId   string    `bson:"userId"`
	Name     string    `bson:"name"`
	Total    int64     `bson:"total"`
	Products []Product `bson:"products"`
}
