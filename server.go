package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func fetchAllTodo(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		gin.H{"status": http.StatusOK, "message": "Hello world"})
	return
}

func main() {
	router := gin.Default()
	router.GET("/ping", fetchAllTodo)
	router.Run()
}
