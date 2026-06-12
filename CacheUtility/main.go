package main

import (
	"fmt"
)

// creating 2 empty interfaces to Add flexiblity in key value type of data
type KEY any
type VALUE any

func main() {

	fmt.Println("Cache util")
	
	DataStorage := make(map[KEY]VALUE) // declare a Map to store key value pairs when program is running

	cacheUtil(&DataStorage) // calling cache utility
	
}
