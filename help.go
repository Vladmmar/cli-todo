package main

import (
	"fmt"
	"os"
)

func availableOptions() {
	fmt.Println(`
Tasks is a todo list cli app

Options:
	add - add new task
	complete - mark task as completed
	...
	`)
	os.Exit(0)
}

func help() {

}
