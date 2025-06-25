package main

import (
	"encoding/csv"
	"errors"
	"os"
	"path/filepath"
	"time"
)

type Task struct {
	Id      int
	Task    string
	Created time.Time
	Done    bool
}

func initTasks() {
	if _, err := os.Stat(filepath.Join(appDataPath, "tasks.csv")); errors.Is(err, os.ErrNotExist) {
		file, createErr := os.Create(filepath.Join(appDataPath, "tasks.csv"))
		if createErr != nil {
			panic(createErr)
		}
		defer func(file *os.File) {
			err = file.Close()
			if err != nil {
				panic(err)
			}
		}(file)

		writer := csv.NewWriter(file)

		headers := []string{"ID", "Task", "Created", "Done"}
		err = writer.Write(headers)
		if err != nil {
			panic(err)
		}

		writer.Flush()
	}
}
