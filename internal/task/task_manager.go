package task

import "fmt"

type Task struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	Status bool   `json:"status"`
}

func addTask(task string) {
	fmt.Println(task)
}
