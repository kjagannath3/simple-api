package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	router:= gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("books/:id", getById)
	router.Run("localhost:8080")

}

type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", Quantity: 1},
	{ID: "2", Title: "Cloud Native Go", Author: "M.-L. Reimer", Quantity: 2},
	{ID: "3", Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", Quantity: 3},
	{ID: "4", Title: "Cloud Native Go", Author: "M.-L. Reimer", Quantity: 4},
	{ID: "5", Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", Quantity: 5},
	{ID: "6", Title: "Cloud Native Go", Author: "M.-L. Reimer", Quantity: 6},
}


func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err:=c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}


func getBookById (id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}


func getById (c *gin.Context) {
	id := c.Param("id")
	current_book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, current_book)
}
