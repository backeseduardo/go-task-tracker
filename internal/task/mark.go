package task

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"backeseduardo.com/task-tracker/internal"
)

func Mark(strid string, status Status) {
	id, err := strconv.Atoi(os.Args[2])
	internal.Check(err)
	changeTaskStatus(id, status)

	fmt.Printf("Task %d status changed\n", id)
}

func changeTaskStatus(id int, status Status) {
	db := LoadDB()

	for i := 0; i < len(db.Tasks); i++ {
		if db.Tasks[i].Id == id {
			db.Tasks[i].Status = status
			db.Tasks[i].UpdatedAt = time.Now()
			db.Tasks[i].Updated = true
		}
	}

	db.Save()
}
