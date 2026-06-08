package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	source := args[1]
	dest := args[2]

	info, err := os.Stat(source)

	if err != nil {
		fmt.Println("Invalid path")
		return
	}
	if !(info.IsDir()) { // if the source is just a file then just copy paste
		CopyUtil(source, dest, "")
	} else {
		directory, err := os.ReadDir(source)
		if err != nil {
			fmt.Println("some error during reading source directory")
			return
		}
		destDets, err2 := os.Stat(dest)
		if err2 != nil {
			fmt.Println("something wrong with dest")
			return
		}
		if !(destDets.IsDir()) { //checking for source to be dir and dest to be file if yes then return
			fmt.Println("given dest path is a file path, cant move a directory to it")
			return
		}

		for _, fileName := range directory {
			str := fmt.Sprintf("%v/%v", source, fileName.Name())
			CopyUtil(str, dest, fileName.Name())
			os.Remove(str) // removing files form source
		}
		// fmt.Println("movement successful")
		err3 := os.Remove(source) // removing files form source
		if err3 != nil {
			fmt.Println("problem in file deletion")
		}
		fmt.Println("files deleted")
	}

}
