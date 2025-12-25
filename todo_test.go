package main

import (
	"reflect"
	"testing"
)

func TestAddTodo(t *testing.T) {
	todos := []Todo{}
	args := []string{"todo", "add", "Buy Milk"}

	got := addTodo(todos, args)

	if len(got) != 1 {
		t.Fatalf("Expected 1 todo but got %d", len(got))
	}

	if got[0].Description != "Buy Milk" {
		t.Errorf("expected description: %q, got %q", "\"Buy Milk\"", got[0].Description)
	}
	if got[0].Completed {
		t.Errorf("expected new to be not completed")
	}
}

func TestMarkCompleted(t *testing.T) {
	todos := []Todo{
		{ID: 1, Description: "Task 1"},
		{ID: 2, Description: "Task 2"},
	}
	got := markCompleted(todos, 1)

	if !got[0].Completed {
		t.Errorf("expected todo at index 0 to be completed")
	}
}

func TestRemoveTodo(t *testing.T) {
	todos := []Todo{
		{ID: 1, Description: "Task 1"},
		{ID: 2, Description: "Task 2"},
		{ID: 3, Description: "Task 3"},
	}

	args := []string{"todo", "remove", "1"}
	got := removeTodo(todos, args, "test.json")

	want := []Todo{
		{ID: 1, Description: "Task 1"},
		{ID: 3, Description: "Task 3"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %#v, got %#v", want, got)
	}
}
