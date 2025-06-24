package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

func list(extended bool) {
	file, err := os.Open("tasks.csv")
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

		if extended {
			_, err = fmt.Fprintln(writer, strings.Join(record, "\t"))
			if err != nil {
				return
			}
		} else {
			if done, parseErr := strconv.ParseBool(record[3]); parseErr == nil {
				if !done {
					createdTime, timeErr := time.Parse(time.RFC3339, record[2])
					if timeErr != nil {
						panic(timeErr)
					}
					record[2] = toReadableTime(&createdTime)
					_, err = fmt.Fprintln(writer, strings.Join(record, "\t"))
					if err != nil {
						panic(err)
					}
				}
			} else {
				_, err = fmt.Fprintln(writer, strings.Join(record, "\t"))
				if err != nil {
					panic(err)
				}
			}
		}
	}
	writer.Flush()
}
