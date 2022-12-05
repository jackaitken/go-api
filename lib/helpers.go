package lib

import (
	"encoding/json"
	"errors"
	"log"
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
		log.Fatal(err)
	}

	for i := range existingLists {
		if existingLists[i].Id == todoListId {
			list := &existingLists[i]
			list.Todos = append(list.Todos, todo)

			if err := SaveJSON(existingLists); err != nil {
				log.Fatal(err)
			}

			return nil
		}
	}

	return errors.New("no list with that id was found")
}

func GetTodoList(id int) (TodoList, error) {
	todoLists, err := LoadJSON()

	if err != nil {
		log.Fatal(err)
	}

	for _, list := range todoLists {
		if list.Id == id {
			return list, nil
		}
	}
	return TodoList{}, errors.New("could not find list with that id")
}

func GetTodo(id int) (*Todo, error) {
	todoLists, err := LoadJSON()

	if err != nil {
		log.Fatal(err)
	}

	for _, list := range todoLists {
		for _, todo := range list.Todos {
			if todo.Id == id {
				todoAddress := &todo
				return todoAddress, nil
			}
		}
	}
	return &Todo{}, errors.New("no todo with that id was found")
}

func EditTodo(id int, todoTitle Todo) error {
	// Currently can only edit the todo title

	todo, err := GetTodo(id)

	if err != nil {
		return errors.New("no todo with that id was found")
	}

	todo.Title = todoTitle.Title

	return nil
}
