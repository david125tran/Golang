package main

import (
	"bufio"
	"example.com/go-project/note"
	"example.com/go-project/todo"
	"fmt"
	"os"
	"strings"
)

// type displayer interface {
// 	Display()
// }

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	printSomething(1)
	printSomething(1.50)
	printSomething("Hello")

	//----- Get user input -----
	title, content := getNoteData()

	todoText := getTodoData()

	// ----- Todo logic -----
	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	// ----- Note Logic -----
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func printSomething(value interface{}) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer:", intVal)
		return
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float:", floatVal)
		return
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("String:", stringVal)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println("String:", value)
	// default:
	// 	fmt.Println("Other data type:", value)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Error saving the note:", err)
		return err
	}

	fmt.Println("\nData saved.")
	return nil
}

func getTodoData() string {
	return getUserInput("Todo text:")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	// Remove text
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
