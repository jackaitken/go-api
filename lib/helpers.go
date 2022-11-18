package lib

import (
	"encoding/json"
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

/*
We need a function to create a new TodoList

*/
