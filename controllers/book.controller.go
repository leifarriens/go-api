package controllers

import (
	"goapi/models"
	"goapi/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookController struct {
	BookService services.BookService
}

func New(bookservice services.BookService) BookController {
	return BookController{
		BookService: bookservice,
	}
}

func (bc *BookController) CreateBook(ctx *gin.Context) {
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := bc.BookService.CreateBook(&book)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (bc *BookController) GetBook(ctx *gin.Context) {
	bookid, err := primitive.ObjectIDFromHex(ctx.Param(("bookId")))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	book, err := bc.BookService.GetBook(&bookid)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func (bc *BookController) GetAllBooks(ctx *gin.Context) {
	users, err := bc.BookService.GetAllBooks()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, users)
}

func (bc *BookController) UpdateBook(ctx *gin.Context) {
	bookid, err := primitive.ObjectIDFromHex(ctx.Param(("bookId")))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	err = bc.BookService.UpdateBook(&bookid, &book)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (bc *BookController) DeleteBook(ctx *gin.Context) {
	bookid, err := primitive.ObjectIDFromHex(ctx.Param(("bookId")))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := bc.BookService.DeleteBook(&bookid); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (bc *BookController) RegisterBookRoutes(rg *gin.RouterGroup) {
	bookroute := rg.Group("/books")
	bookroute.POST("/", bc.CreateBook)
	bookroute.GET("/:bookId", bc.GetBook)
	bookroute.GET("/", bc.GetAllBooks)
	bookroute.PATCH("/:bookId", bc.UpdateBook)
	bookroute.DELETE("/:bookId", bc.DeleteBook)
}
