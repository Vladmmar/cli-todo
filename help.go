package main

import (
	"fmt"
	"os"
)

func help() {
	helpText := `
Todo List App - Command Line Usage

todo COMMAND [options]

Available Commands:

  add <task>                Add a new task to the list.
  complete <id> [ids...]    Mark one or more tasks as completed.
                            Example: complete 1 3 5
  delete <id> [ids...]      Delete one or more tasks.
                            Example: delete 2 4 6
  list                      Show all incomplete tasks.
  list [-a --all]           Show all tasks, including completed ones.
  help                      Show this help message.

Notes:
  - Task IDs are integers (1, 2, 3, ...).
  - Commands are case-sensitive.`

	fmt.Println(helpText)
	os.Exit(0)
}
