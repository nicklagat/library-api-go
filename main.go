package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = map[string]Book{}

func main() {

	router := gin.Default()

	router.GET("/books", getAllBooks)
	router.POST("/books", createBook)

	err := router.Run(":8080")
	if err != nil {
		return
	}

}

func getAllBooks(c *gin.Context) {

	var bookList []Book
	for _, book := range books {
		bookList = append(bookList, book)
	}

	c.JSON(http.StatusOK, bookList)
}

func createBook(c *gin.Context) {
	var book Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.ID = uuid.New().String()
	books[book.ID] = book

	c.JSON(http.StatusCreated, book)
}
