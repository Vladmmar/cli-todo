package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

func list(extended bool) {
	file, err := os.Open(filepath.Join(appDataPath, "tasks.csv"))
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	reader := csv.NewReader(file)
	var record []string
	for {
		record, err = reader.Read()
		if record == nil {
			break
		}
		if err != nil {
			panic(err)
		}

		done, parseErr := strconv.ParseBool(record[3])

		if parseErr == nil {
			createdTime, timeErr := time.Parse(time.RFC3339, record[2])
			if timeErr != nil {
				panic(timeErr)
			}
			record[2] = toReadableTime(&createdTime)
		}

		if extended {
			_, err = fmt.Fprintln(writer, strings.Join(record, "\t"))
			if err != nil {
				panic(err)
			}
		} else {
			if !done {
				_, err = fmt.Fprintln(writer, strings.Join(record[:3], "\t"))
				if err != nil {
					panic(err)
				}
			}
		}
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}
