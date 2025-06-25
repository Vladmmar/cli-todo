package main

import (
	"os"
	"path/filepath"
)

var appConfPath string
var appDataPath string

// CheckHealth Checks the health the app
func CheckHealth() {
	sysConfPath, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	appConfPath = filepath.Join(sysConfPath, "/todo-list")
	appDataPath = filepath.Join(sysConfPath, "/todo-list")

	loadConfig()
	initTasks()
}

func initConf() {
	config = Config{}
	saveConfig()
}
