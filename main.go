package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jackaitken/go-api/routes"
	"log"
)

func main() {
	router := gin.Default()

	router.GET("/todolists", routes.Home)

	router.GET("/todolist/:id", routes.TodoList)

	router.GET("/todo/:id", routes.HandleGetTodo)

	router.PUT("/todo/:id", routes.EditTodo)

	router.DELETE("/todo/:id", routes.DeleteTodo)

	router.POST("/todolist", routes.NewTodoList)

	router.POST("/todolist/:id/new-todo", routes.NewTodo)

	err := router.Run(":8080")

	if err != nil {
		log.Fatal(err)
	}
}

/*
Should be able to:

1. get all todos
2. get a single todo
3. edit a todo
4. delete a todo
5. mark a todo as done
6. create a todo
7. create a todo list
8.

*/
