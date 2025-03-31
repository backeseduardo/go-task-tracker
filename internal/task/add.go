package task

import (
	"fmt"
	"time"
)

func Add(d string) {
	db := LoadDB()

	task := Task{
		Id:          db.NextSeq(),
		Description: d,
		Status:      STATUS_TODO,
		CreatedAt:   time.Now(),
	}
	db.Tasks = append(db.Tasks, task)
	db.Save()

	fmt.Printf("Task \"%s\" added\n", d)
}
