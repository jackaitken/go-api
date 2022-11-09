package helpers

import (
	"encoding/json"
	"github.com/jackaitken/go-api/structs"
	"os"
)

func LoadJson() ([]structs.TodoList, error) {
	jsonFile, err := os.ReadFile("sample/sample_data.json")

	if err != nil {
		return nil, err
	}

	var todoLists []structs.TodoList

	err = json.Unmarshal(jsonFile, &todoLists)

	if err != nil {
		return nil, err
	}

	return todoLists, nil
}
