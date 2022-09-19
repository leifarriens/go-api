package models

type Book struct {
	Name   string `json:"name" bson:"name"`
	Author string `json:"author" bson:"author"`
}

type BookWithId struct {
	ID     string `json:"_id" bson:"_id"`
	Name   string `json:"name" bson:"name"`
	Author string `json:"author" bson:"author"`
}
