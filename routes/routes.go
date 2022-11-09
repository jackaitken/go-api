package routes

import (
	"net/http"
	"strconv"

	"github.com/jackaitken/go-api/helpers"

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

func TodoList(c *gin.Context) {
	todoLists, err := helpers.LoadJson()

	if err != nil {
		panic(err)
	}

	requestedTodoListId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	for _, list := range todoLists {
		if list.Id == requestedTodoListId {
			c.JSON(http.StatusOK, gin.H{
				list.ListName: list.Todos,
			})
		}
	}

	c.String(http.StatusNotFound,
		"No list with id: '%d' was found", requestedTodoListId)
}

func GetTodo(c *gin.Context) {
	// todoId := c.Param("id")
	//
	// todoLists, err := helpers.LoadJson()
	//
	// if err != nil {
	// 	panic(err)
	// }

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
