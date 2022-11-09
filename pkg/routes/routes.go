package routes

import (
	"github.com/jackaitken/go-api/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	todoLists, err := helpers.LoadJson()

	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"lists": todoLists,
	})
}

func GetTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "One todo",
	})
}

func EditTodo(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"message": "Todo edited",
	})
}

func DeleteTodo(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"message": "Todo deleted",
	})
}

func NewTodo(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"message": "Todo created",
	})
}
