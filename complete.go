package main

import (
	"encoding/csv"
	"os"
	"slices"
	"strconv"
)

func complete(ids ...string) {
	tasks, err := os.OpenFile("tasks.csv", 0, 0644)
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(tasks)

	newTasks, err := os.Create("_tasks.csv")
	writer := csv.NewWriter(newTasks)
	record, err := reader.Read()
	if err != nil {
		panic(err)
	}
	err = writer.Write(record)
	if err != nil {
		panic(err)
	}

	for {
		record, err = reader.Read()
		if record == nil {
			break
		}
		if err != nil {
			panic(err)
		}
		id, err := strconv.Atoi(record[0])
		if err != nil {
			panic(err)
		}
		if slices.Contains(ids, strconv.Itoa(id)) {
			record[3] = "true"
		}
		err = writer.Write(record)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	err = tasks.Close()
	if err != nil {
		panic(err)
	}
	err = newTasks.Close()
	if err != nil {
		panic(err)
	}

	err = os.Remove("tasks.csv")
	if err != nil {
		panic(err)
	}
	err = os.Rename("_tasks.csv", "tasks.csv")
	if err != nil {
		panic(err)
	}
}
