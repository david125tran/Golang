package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Printf(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "out\\todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	// Return the error if file writing fails. Else return nil
	return os.WriteFile(fileName, json, 0644)

}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Error: Invalid input.")
	}
	return Todo{
		Text: content,
	}, nil
}
