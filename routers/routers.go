package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "All todos",
	})
}

func GetTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "One todo",
	})
}
