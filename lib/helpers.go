package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func LoadJSON() ([]TodoList, error) {
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

func SaveJSON(todoLists []TodoList) error {
	result, err := json.Marshal(todoLists)

	if err != nil {
		return err
	}

	err = os.WriteFile("sample/sample_data.json", result, 0666)

	if err != nil {
		return err
	}

	return nil
}

func AppendTodoList(newList TodoList) error {
	existingLists, err := LoadJSON()

	if err != nil {
		return err
	}

	existingLists = append(existingLists, newList)

	err = SaveJSON(existingLists)

	if err != nil {
		return err
	}

	return nil
}

func AppendTodo(todo Todo, todoListId int) error {
	existingLists, err := LoadJSON()

	if err != nil {
		panic(err)
	}

	for i := range existingLists {
		if existingLists[i].Id == todoListId {
			list := &existingLists[i]
			list.Todos = append(list.Todos, todo)

			if err := SaveJSON(existingLists); err != nil {
				panic(err)
			}
			return nil
		}
	}

	return errors.New("no list with that id was found")
}

func GetTodoList(id int) (TodoList, error) {
	todoLists, err := LoadJSON()

	if err != nil {
		panic(err)
	}

	for _, list := range todoLists {
		if list.Id == id {
			return list, nil
		}
	}
	return TodoList{}, errors.New("Could not find list with that id")
}

func GetTodo(id int) (Todo, error) {
	todoLists, err := LoadJSON()

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
