package models

type Product struct {
	Id          string      `bson:"id"`
	Price       int64       `bson:"price"`
	Weight      uint        `bson:"weight"`
	Category    string      `bson:"category"`
	PerUnit     int         `bson:"perUnit"`
	Description string      `bson:"description"`
	ImageUrl    interface{} `bson:"imageUrl"`
	Name        string      `bson:"name"`
	CategoryId  string      `bson:"categoryId"`
}
