package main

import (
	"fmt"
)

type KEY any
type VALUE any

func main() {

	fmt.Println("Cache util")
	DataStorage := make(map[KEY]VALUE)
	cacheUtil(&DataStorage)
}
