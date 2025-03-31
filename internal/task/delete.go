package task

import (
	"fmt"
	"strconv"

	"backeseduardo.com/task-tracker/internal"
)

func Delete(strids []string) {
	db := LoadDB()
	var ids []int

	for _, i := range strids {
		id, err := strconv.Atoi(i)
		internal.Check(err)
		ids = append(ids, id)
	}

	var tasks []Task
	for _, t := range db.Tasks {
		var isIn bool
		for _, id := range ids {
			if t.Id == id {
				isIn = true
			}
		}
		if !isIn {
			tasks = append(tasks, t)
		}
	}

	db.Tasks = tasks
	if len(db.Tasks) == 0 {
		db.Seq = 0
	}
	db.Save()

	fmt.Printf("Task %v ids deleted\n", ids)
}
