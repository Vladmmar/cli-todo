package main

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"slices"
	"strconv"
)

func complete(ids ...string) {
	tasks, err := os.OpenFile(filepath.Join(appDataPath, "tasks.csv"), 0, 0644)
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(tasks)

	newTasks, err := os.Create(filepath.Join(appDataPath, "_tasks.csv"))
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

	err = os.Remove(filepath.Join(appDataPath, "tasks.csv"))
	if err != nil {
		panic(err)
	}
	err = os.Rename(filepath.Join(appDataPath, "_tasks.csv"), filepath.Join(appDataPath, "tasks.csv"))
	if err != nil {
		panic(err)
	}
}
