package main

import (
	"context"
	"fmt"
	"goapi/controllers"
	"goapi/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	bookservice    services.BookService
	bookcontroller controllers.BookController
	ctx            gin.Context
	bookcollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	database := os.Getenv("DATABASE")

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	mongoclient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	err = mongoclient.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb")

	bookcollection = mongoclient.Database(database).Collection(("books"))
	bookservice = services.NewBookService(bookcollection, context.TODO())
	bookcontroller = controllers.New(bookservice)
	server = gin.Default()
	server.SetTrustedProxies([]string{"localhost"})
}

func main() {
	defer mongoclient.Disconnect(context.TODO())

	basepath := server.Group("/")
	bookcontroller.RegisterBookRoutes(basepath)

	log.Fatal(server.Run("localhost:8080"))
}
