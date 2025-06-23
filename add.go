package main

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"time"
)

func add(description string) {
	var task Task
	task.Id = getNextId()
	task.Created = time.Now()
	task.Task = description
	task.Done = false
	saveTask(task)
}

func saveTask(task Task) {
	var file *os.File
	var err error
	var created bool = false
	if _, err = os.Stat("tasks.csv"); errors.Is(err, os.ErrNotExist) {
		file, err = os.Create("tasks.csv")
		created = true
		if err != nil {
			panic(err)
		}

	} else {
		file, err = os.OpenFile("tasks.csv", os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	writer := csv.NewWriter(file)

	if created {
		headers := []string{"ID", "Task", "Created", "Done"}
		err := writer.Write(headers)
		if err != nil {
			panic(err)
		}
	}

	record := []string{strconv.Itoa(task.Id), task.Task, task.Created.Format(time.RFC3339), strconv.FormatBool(task.Done)}
	err = writer.Write(record)
	if err != nil {
		panic(err)
	}
	writer.Flush()
}
