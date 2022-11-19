package routes

import (
	"github.com/jackaitken/go-api/lib"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	todoLists, err := lib.LoadJSON()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"lists": todoLists,
	})
}

func TodoList(c *gin.Context) {
	requestedTodoListId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	requestedTodoList, err := lib.GetTodoList(requestedTodoListId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			requestedTodoList.ListName: requestedTodoList.Todos,
		})
	}
}

func HandleGetTodo(c *gin.Context) {
	requestedTodoId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	requestedTodo, err := lib.GetTodo(requestedTodoId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"title":    requestedTodo.Title,
			"status":   requestedTodo.Status,
			"priority": requestedTodo.Priority,
			"dueDate":  requestedTodo.DueDate,
			"notes":    requestedTodo.Notes,
		})
	}
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
	newTodoList := lib.TodoList{}

	if err := c.BindJSON(&newTodoList); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := lib.AppendTodoList(newTodoList); err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "New list created",
	})
}

func NewTodo(c *gin.Context) {
	todoListId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	newTodo := lib.Todo{}

	if err := c.BindJSON(&newTodo); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = lib.AppendTodo(newTodo, todoListId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"message": "New todo was added",
		})
	}
}
