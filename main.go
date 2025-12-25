package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type Todo struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

const TodoFileName = "todos.json"

func main() {
	initToDoFile()

	todos := readFromFile(TodoFileName)

	command := os.Args[1]

	if len(os.Args) < 2 {
		fmt.Println("Usage: todo add <todo description>")
		return
	}

	switch command {
	case "add":
		todos = addTodo(todos, os.Args)
		writeToFile(TodoFileName, todos)
	case "complete":
		completeTodo(todos, os.Args, TodoFileName)
	case "list":
		list(todos)
	case "remove":
		removeTodo(todos, os.Args, TodoFileName)
	default:
		color.New(color.FgRed).Printf("Unknown command: %s", command)
	}
}
