package routes

import (
	"github.com/jackaitken/go-api/lib"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	todoLists, err := lib.LoadJson()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"lists": todoLists,
	})
}

func TodoList(c *gin.Context) {
	todoLists, err := lib.LoadJson()

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
			return
		}
	}

	c.String(http.StatusNotFound,
		"No list with id: '%d' was found", requestedTodoListId)
}

func GetTodo(c *gin.Context) {
	todoLists, err := lib.LoadJson()

	if err != nil {
		panic(err)
	}

	requestedTodoId, err := strconv.Atoi(c.Param("id"))

	for _, list := range todoLists {
		for _, todo := range list.Todos {
			if todo.Id == requestedTodoId {
				c.JSON(http.StatusOK, gin.H{
					"title":    todo.Title,
					"status":   todo.Status,
					"priority": todo.Priority,
					"dueDate":  todo.DueDate,
					"notes":    todo.Notes,
				})
				return
			}
		}
	}

	c.String(http.StatusNotFound,
		"No todo with id: '%d' was found", requestedTodoId)
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

func NewTodoList(c *gin.Context) {
	todoLists, err := lib.LoadJson()

	if err != nil {
		panic(err)
	}

	newTodoList := lib.TodoList{}

	if err = c.BindJSON(&newTodoList); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err = lib.AppendTodoList(newTodoList, todoLists); err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "New list created",
	})

	/*
		To create a new todolist, we should take a request body
		We need to do some error checking that we've received everything
		that we need to nothing that we don't.
	*/
}
