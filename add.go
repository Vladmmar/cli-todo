package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

func add(description string) {
	id := getNextId()
	timestamp := time.Now()
	fmt.Println(timestamp)
	record := []string{strconv.Itoa(id), description, timestamp.String(), "false"}

	saveRecord(record)
}

func saveRecord(record []string) {
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

	err = writer.Write(record)
	if err != nil {
		panic(err)
	}
	writer.Flush()
}
