package routes

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jackaitken/go-api/lib"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	todoLists, err := lib.LoadJSON()

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"lists": todoLists,
	})
}

func TodoList(c *gin.Context) {
	requestedTodoListId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	requestedTodo, err := lib.GetTodo(requestedTodoId)
	fmt.Println(requestedTodo)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":       requestedTodo.Id,
			"title":    requestedTodo.Title,
			"status":   requestedTodo.Status,
			"priority": requestedTodo.Priority,
			"dueDate":  requestedTodo.DueDate,
			"notes":    requestedTodo.Notes,
		})
	}
}

func EditTodo(c *gin.Context) {
	requestedTodoId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Fatal(err)
	}

	todoTitle := lib.Todo{}

	if err := c.BindJSON(&todoTitle); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := lib.EditTodo(requestedTodoId, todoTitle); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"message": "todo title was edited",
		})
	}
	/*
		To edit a todo:
		1. Get the new todo wording (bind the body)
		2. Get the todo by id
		3. Save the new todo title

		We have:
		1. Get the todo id
		2. Save the new todo title from the request body

		Need to:
		1. Replace that todo in the right todolist

		We can probably just find the todo, and replace the
		todo title for right now.

		For the future: I need to be able to edit the other
		properties. For right now we're just going to edit the title
	*/
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
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "New list created",
	})
}

func NewTodo(c *gin.Context) {
	todoListId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Fatal(err)
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
