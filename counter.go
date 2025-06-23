package main

func getNextId() int {
	config.Cnt++
	saveConfig()

	return config.Cnt
}
