package main

import "time"

type Task struct {
	Id      int
	Task    string
	Created time.Time
	Done    bool
}
