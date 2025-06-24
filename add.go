package main

import (
	"encoding/csv"
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
	file, err = os.OpenFile("tasks.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	writer := csv.NewWriter(file)

	record := []string{strconv.Itoa(task.Id), task.Task, task.Created.Format(time.RFC3339), strconv.FormatBool(task.Done)}
	err = writer.Write(record)
	if err != nil {
		panic(err)
	}
	writer.Flush()
}
