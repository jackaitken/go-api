package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func LoadJson() ([]TodoList, error) {
	jsonFile, err := os.ReadFile("sample/sample_data.json")

	if err != nil {
		return nil, err
	}

	var todoLists []TodoList

	err = json.Unmarshal(jsonFile, &todoLists)

	if err != nil {
		return nil, err
	}

	return todoLists, nil
}

func AppendTodoList(newList TodoList, existingLists []TodoList) error {
	existingLists = append(existingLists, newList)

	result, err := json.Marshal(existingLists)

	if err != nil {
		return err
	}

	err = os.WriteFile("sample/sample_data.json", result, 0666)

	if err != nil {
		return err
	}

	return nil
}

func GetTodo(id int) (Todo, error) {
	todoLists, err := LoadJson()

	if err != nil {
		panic(err)
	}

	for _, list := range todoLists {
		for _, todo := range list.Todos {
			if todo.Id == id {
				return todo, nil
			}
		}
	}
	return Todo{}, errors.New(fmt.Sprintf("No todo with that id was found"))
}

/*
We need a function to create a new TodoList

*/
