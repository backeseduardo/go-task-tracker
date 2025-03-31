package task

import (
	"fmt"
	"strconv"
	"time"

	"backeseduardo.com/task-tracker/internal"
)

func Update(strid string, d string) {
	db := LoadDB()

	id, err := strconv.Atoi(strid)
	internal.Check(err)

	for i := 0; i < len(db.Tasks); i++ {
		fmt.Printf("%d - %d", i, db.Tasks[i].Id)
		if db.Tasks[i].Id == id {
			db.Tasks[i].Description = d
			db.Tasks[i].UpdatedAt = time.Now()
			db.Tasks[i].Updated = true
			fmt.Println(db.Tasks[i])
			break
		}
	}

	db.Save()

	fmt.Printf("Task %d updated\n", id)
}
