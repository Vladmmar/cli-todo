package main

import (
	"fmt"
	"os"
)

func main() {
	Initialize()
	Run()
}

func Run() {
	args := os.Args
	if len(args) <= 1 {
		help()
	}
	command := args[1]
	switch command {
	case "help", "--help", "-h":
		help()
	case "add":
		if len(args) != 3 {
			fmt.Println("Enter the description for a new task\n" +
				"tasks add <description>")
		} else {
			add(args[2])
		}
	case "list":
		if len(args) >= 3 {
			if args[2] == "-a" || args[2] == "--all" {
				list(true)
			}
		} else {
			list(false)
		}
	case "complete":
		if len(args) == 2 {
			fmt.Println("Enter id of tasks to be completed\n" +
				"tasks complete <...id>")
		} else {
			complete(args[2:]...)
		}
	case "delete":
		if len(args) == 2 {
			fmt.Println("Enter id of tasks to be deleted\n" +
				"tasks delete <...id>")
		} else {
			del(args[2:]...)
		}
	}
}
