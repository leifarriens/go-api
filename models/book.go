package models

type Book struct {
	Bookid string `json:"id" bson:"_id"`
	Name   string `json:"name" bson:"name"`
	Author string `json:"author" bson:"author"`
}
