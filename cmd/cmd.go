package cmd

import (
	"fmt"
	"os"

	"backeseduardo.com/task-tracker/internal/task"
)

const (
	Bold = "\033[1m%s\033[22m"
)

func printHelpMenu() {
	fmt.Printf(Bold, "Task Tracker")
	fmt.Printf("\nUsage:\n$ task-tracker <action>\n")
	fmt.Print("\n")
	fmt.Printf(Bold, "Actions\n")
	fmt.Printf(Bold, "add:")
	fmt.Printf(" task-tracker add \"My first task\"\n")
	fmt.Printf(Bold, "update:")
	fmt.Printf(" task-tracker update 1 \"My task\"\n")
	fmt.Printf(Bold, "delete:")
	fmt.Printf(" task-tracker delete 1\n")
	fmt.Printf(Bold, "list:")
	fmt.Printf(" task-tracker list\n")
	fmt.Printf(Bold, "mark-in-progress:")
	fmt.Printf(" task-tracker mark-in-progress 1\n")
	fmt.Printf(Bold, "mark-done:")
	fmt.Printf(" task-tracker mark-done 1\n")
}

func Init() {
	if len(os.Args) < 2 {
		printHelpMenu()
		os.Exit(0)
	}

	action := os.Args[1]

	switch action {
	case "add":
		if len(os.Args) != 3 {
			fmt.Println("Invalid number of arguments. Try calling ./task-tracker add \"My first task\"")
			os.Exit(2)
		}
		task.Add(os.Args[2])
	case "update":
		if len(os.Args) != 4 {
			fmt.Println("Invalid number of arguments. Try calling ./task-tracker update 1 \"My first task\"")
			os.Exit(3)
		}
		task.Update(os.Args[2], os.Args[3])
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Invalid number of arguments. Try calling ./task-tracker delete 1 or ./task-tracker delete 1 2")
			os.Exit(4)
		}
		task.Delete(os.Args[2:])
	case "list":
		if len(os.Args) > 3 {
			fmt.Println("Invalid number of arguments. Try calling ./task-tracker list or ./task-tracker list todo")
			os.Exit(5)
		}
		if len(os.Args) > 2 {
			task.List(os.Args[2])
		} else {
			task.List("")
		}
	case "mark-in-progress":
		if len(os.Args) != 3 {
			fmt.Println("Invalid number of arguments. Try calling ./task-tracker mark-in-progress 1")
			os.Exit(6)
		}
		task.Mark(os.Args[2], task.STATUS_IN_PROGRESS)
	case "mark-done":
		if len(os.Args) != 3 {
			fmt.Println("Invalid number of arguments. Try calling ./task-tracker mark-done 1")
			os.Exit(7)
		}
		task.Mark(os.Args[2], task.STATUS_DONE)
	default:
		fmt.Println("Invalid action. Try add, update, delete, mark-in-progress, mark-done or list.")
		os.Exit(1)
	}
}
