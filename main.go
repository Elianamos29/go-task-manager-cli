package main

import (
	"fmt"
)

func main() {
	tasks := []Task{
		{ID: 1, Name: "Learn Go", Done: false},
		{ID: 2, Name: "Write cli", Done: false},
	}

	fmt.Println("Your tasks:")
	for _, task := range tasks {
		fmt.Printf("%d. %s [Done: %t]\n", task.ID, task.Name, task.Done)
	}
}