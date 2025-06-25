package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Config struct {
	Cnt int `json:"counter"`
}

var config Config

func loadConfig() {
	if _, err := os.Stat(filepath.Join(appConfPath, "config.json")); errors.Is(err, os.ErrNotExist) {
		initConf()
	} else {
		jsonData, err2 := os.ReadFile(filepath.Join(appConfPath, "config.json"))
		if err2 != nil {
			panic(err2)
		}

		err3 := json.Unmarshal(jsonData, &config)
		if err3 != nil {
			panic(err3)
		}
	}
}

func saveConfig() {
	data, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(appConfPath, 0644)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(filepath.Join(appConfPath, "config.json"))
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
}
