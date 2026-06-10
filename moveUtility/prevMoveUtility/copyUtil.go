package main

import (
	// "bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CopyUtil(source, dest, fileName string) {

	data, err := os.ReadFile(strings.TrimSpace(source))

	if err != nil {
		fmt.Println("error while reading ", err)
	} else {
		if fileName == "" {
			destDets, _ := os.Stat(dest)
			if destDets.IsDir() { // case where source is a file and dest is a dir
				dest += filepath.Base(source)
			}
			//case where source and dest both are files
			err3 := os.WriteFile(strings.TrimSpace(dest), data, 0644)
			if err3 != nil {
				fmt.Println("error while writing ", err3)
			}
			return
		}
		// case where both are dir
		dest += fileName // adding filename to the dest directory
		err3 := os.WriteFile(strings.TrimSpace(dest), data, 0644)
		if err3 != nil {
			fmt.Println("error while writing ", err3)
		}
	}

}

// func main() {
// 	// fmt.Println("enter source path")
// 	// reader := bufio.NewReader(os.Stdin)
// 	// source, err := reader.ReadString('\n')
// 	// if err != nil {
// 	// 	fmt.Println("source read not done successfully", err)
// 	// }

// 	// fmt.Println(source)

// 	// data, err2 := os.ReadFile(strings.TrimSpace(source)) //using trimSpace as \n also added when hitting enter
// 	// if err2 != nil {
// 	// 	fmt.Println("problem in reading", err2)
// 	// }

// 	// fmt.Println("enter dest path")
// 	// dest, err3 := reader.ReadString('\n')
// 	// if err3 != nil {
// 	// 	fmt.Println("dest read not done successfully", err)
// 	// }
// 	// fmt.Println(dest)

// 	// err4 := os.WriteFile(strings.TrimSpace(dest), data, 0644)
// 	// if err4 != nil {
// 	// 	fmt.Println("error encountered while writing")
// 	// } else {
// 	// 	fmt.Println("file written :) ")
// 	// }

// 	// -> with Args
// }

// algo
// source as path
// -> read the source
// -> dest as path
// -> create a file at dest
// -> write the data of source at dest
