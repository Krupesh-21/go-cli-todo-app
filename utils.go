package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func addTodo(todos []Todo, args []string) []Todo {
	fail := color.New(color.FgRed).PrintfFunc()
	success := color.New(color.FgGreen).PrintfFunc()

	if len(args) < 3 {
		fail("Please provide a description for todo task!")
		return todos
	}

	description := strings.Join(args[2:], " ")
	todo := Todo{
		ID:          len(todos) + 1,
		Description: description,
	}

	todos = append(todos, todo)
	success("Added todo: %v\n", todo)
	return todos
}

func list(todos []Todo) {
	if len(todos) == 0 {
		color.New(color.FgBlue).Println("List is empty,Please add a todo to the list!")
		return
	}

	fmt.Printf("\n%s\t%s\n\n",
		color.New(color.BgGreen).Sprint("Completed"),
		color.New(color.BgMagenta).Sprint("Incomplete"))

	for _, t := range todos {
		if t.Completed {
			fmt.Printf("%d: %s\n", t.ID, color.New(color.FgGreen, color.CrossedOut).Sprint(t.Description))
		} else {
			fmt.Printf("%d: %s\n", t.ID, color.New(color.FgMagenta).Sprint(t.Description))
		}
	}
}

func writeToFile(fileName string, todos []Todo) {
	jsonData, err := json.MarshalIndent(todos, "", " ")

	if err != nil {
		panic(err)
	}

	writeErr := os.WriteFile(fileName, jsonData, 0644)

	if writeErr != nil {
		panic(writeErr)
	}

	color.GreenString("Data written to file successfully")
}

func readFromFile(fileName string) []Todo {
	fileData, err := os.ReadFile(fileName)

	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}
		}
		panic(err)
	}

	var todos []Todo
	jsonErr := json.Unmarshal(fileData, &todos)

	if jsonErr != nil {
		panic(jsonErr)
	}

	return todos
}

func completeTodo(todos []Todo, args []string, fileName string) {
	if len(args) < 3 {
		color.Red("Usage: todo complete <ID>")
		return
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		color.New(color.FgRed).Printf("Invalid ID: %s", args[2])
		return
	}

	todos = markCompleted(todos, id)
	writeToFile(fileName, todos)
	color.New(color.FgGreen, color.Bold).Printf("Marked todo %d as completed\n", id)
}

func markCompleted(todos []Todo, id int) []Todo {
	if id < 0 || id > len(todos) {
		color.Red("Invalid ID")
		return todos
	}

	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = true
		}
	}
	return todos
}

func removeTodo(todos []Todo, args []string, fileName string) []Todo {
	if len(args) < 3 {
		color.Red("Usage: todo complete <ID>")
		return todos
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		color.New(color.FgRed).Printf("Invalid ID: %s", args[2])
		return todos
	}

	if id < 0 || id > len(todos) {
		color.Red("Invalid ID")
		return todos
	}

	todos = slices.Delete(todos, id, id+1)
	writeToFile(fileName, todos)
	return todos
}

func initToDoFile() {
	if _, err := os.Stat(TodoFileName); os.IsNotExist(err) {
		initialData, _ := json.MarshalIndent([]Todo{}, "", " ")
		if writeErr := os.WriteFile(TodoFileName, initialData, 0644); writeErr != nil {
			panic("Could not create todos.json: " + writeErr.Error())
		}
	}
}
