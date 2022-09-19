package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	ID   string
	Name string
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	database := os.Getenv("DATABASE")

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Successfully connected and pinged.")

	r := gin.Default()

	r.GET("/books", func(c *gin.Context) {
		coll := client.Database(database).Collection("books")

		cursor, err := coll.Find(context.TODO(), bson.M{})

		if err != nil {
			panic(err)
		}

		var results []bson.M
		if err = cursor.All(context.TODO(), &results); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"data": results,
		})
	})

	r.GET("/books/:bookId", func(c *gin.Context) {
		coll := client.Database(database).Collection("books")

		objectId, err := primitive.ObjectIDFromHex(c.Param(("bookId")))

		if err != nil {
			log.Println("Invalid id")
		}

		var result bson.M
		err = coll.FindOne(context.TODO(), bson.D{{"_id", objectId}}).Decode(&result)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	})

	r.POST("/books", func(c *gin.Context) {
		var newBook Book

		err := c.BindJSON(&newBook)

		if err != nil {
			panic(err)
		}

		coll := client.Database(database).Collection("books")
		doc := bson.D{{"name", newBook.Name}}

		result, err := coll.InsertOne(context.TODO(), doc)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	})

	r.DELETE("/books/:bookId", func(c *gin.Context) {
		coll := client.Database(database).Collection("books")

		objectId, err := primitive.ObjectIDFromHex(c.Param(("bookId")))

		if err != nil {
			log.Println("Invalid id")
		}

		result, err := coll.DeleteOne(context.TODO(), bson.D{{"_id", objectId}})

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	})

	r.Run("localhost:8080")
}
