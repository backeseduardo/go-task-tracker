package task

import (
	"fmt"
	"time"
)

func List(strstatus string) {
	db := LoadDB()

	if db.Tasks == nil {
		fmt.Println("No tasks to show")
		return
	}

	fmt.Println("+-----+----+--------------------+---------------------+---------------------+")
	fmt.Println("| st  | id | description        | created at          | updated at          |")
	fmt.Println("+-----+----+--------------------+---------------------+---------------------+")
	for _, t := range db.Tasks {
		if strstatus != "" {
			s := StatusMap[strstatus]
			if t.Status != s {
				continue
			}
		}
		var st byte = ' '
		if t.Status == STATUS_IN_PROGRESS {
			st = '-'
		} else if t.Status == STATUS_DONE {
			st = 'x'
		}
		var updated string = ""
		if t.Updated {
			updated = t.UpdatedAt.Format(time.RFC822)
		}
		fmt.Printf("|  %c  | %2d | %-18s | %s | %-19s |\n", st, t.Id, t.Description, t.CreatedAt.Format(time.RFC822), updated)
	}
	fmt.Println("+-----+----+--------------------+---------------------+---------------------+")
}
