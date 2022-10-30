package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jackaitken/go-api/routers"
)

func main() {
	router := gin.Default()

	router.GET("/todos", routers.Home)

	router.GET("/todo/:id", routers.GetTodo)

	router.Run(":8080")
}
