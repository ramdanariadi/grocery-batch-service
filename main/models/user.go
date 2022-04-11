package models

type User struct {
	Id   string `bson:"id"`
	Name string `bson:"name"`
}
