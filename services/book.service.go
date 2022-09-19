package services

import (
	"goapi/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookService interface {
	CreateBook(*models.Book) error
	GetBook(*primitive.ObjectID) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
	DeleteBook(*primitive.ObjectID) error
}
