package main

import (
	//"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	router:= gin.Default()
	router.GET("/books", getBooks)
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


