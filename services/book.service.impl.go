package services

import (
	"context"
	"errors"
	"goapi/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookServiceImpl struct {
	bookcollection *mongo.Collection
	ctx            context.Context
}

func NewBookService(bookcollection *mongo.Collection, ctx context.Context) BookService {
	return &BookServiceImpl{
		bookcollection: bookcollection,
		ctx:            ctx,
	}
}

func (b *BookServiceImpl) CreateBook(book *models.Book) error {
	_, err := b.bookcollection.InsertOne(b.ctx, book)
	return err
}

func (b *BookServiceImpl) GetBook(id *primitive.ObjectID) (*models.Book, error) {
	var book *models.Book
	filter := bson.D{{"_id", id}}
	err := b.bookcollection.FindOne(b.ctx, filter).Decode(&book)
	return book, err
}

func (b *BookServiceImpl) GetAllBooks() ([]*models.Book, error) {
	var books []*models.Book
	cursor, err := b.bookcollection.Find(b.ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &books); err != nil {
		return nil, err
	}

	cursor.Close(b.ctx)

	return books, nil
}

func (b *BookServiceImpl) DeleteBook(id *primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	result, _ := b.bookcollection.DeleteOne(b.ctx, filter)

	if result.DeletedCount != 1 {
		return errors.New("No matched document found")
	}
	return nil
}
