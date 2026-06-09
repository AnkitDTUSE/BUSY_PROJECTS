package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

var updateMutex sync.Mutex

func MoveUtil(Source, Dest string, fileSize *float64) {
	// retrive base address of Source
	basePath := filepath.Base(Source)
	SourceStat, _ := os.Stat(Source)

	// check if Source is a file or dir
	if !SourceStat.IsDir() {

		// locking fileSize wo consistent values
		updateMutex.Lock()
		(*fileSize) += float64(SourceStat.Size())
		updateMutex.Unlock()

		// THE COMMENTED CODE BELOW IS OF WRITING WHOLE FILE AT ONCE

		// // as source is a file -> reading its data
		// sourceData, err := os.ReadFile(Source)

		// // adding fileSize

		// if err != nil {
		// 	fmt.Println("error while reading SourceFile")
		// 	return
		// }

		// // adding filename at the end of Dest
		Dest := filepath.Join(Dest, basePath)

		// //writing file
		// err2 := os.WriteFile(Dest, sourceData, 0644)

		// if err2 != nil {
		// 	fmt.Println("error while writing file at Dest")
		// 	return
		// }

		// fmt.Println("File written")
		// return

		// FROM NOW ON WE WRITE FILE IN CHUNKS OF 2MB
		// opening file at source

		sourceFile, err := os.Open(Source)
		if err != nil {
			fmt.Println("Error while open file at source")
			return
		}
		defer sourceFile.Close() //closing file

		destFile, err := os.Create(Dest)
		if err != nil {
			fmt.Println("Error while creating file at Dest")
			return
		}
		defer destFile.Close()

		buffer := make([]byte, 2*1024*1024) // byte is an alias for uint8 (0 - 255)

		_, err2 := io.CopyBuffer(destFile, sourceFile, buffer) // _ is for a varible which copybuffer return as an INT8 which is eq to the size of data written

		if err2 != nil {
			fmt.Println("Error while writting at dest")
			return
		}
		fmt.Println("File written")
		return
	}

	// as now source is a Dir -> reading the sturture of this dir
	sourceStructure, _ := os.ReadDir(Source)

	// making required Dir at Dest
	Dest = filepath.Join(Dest, basePath)
	err := os.Mkdir(Dest, 0755)

	if err != nil {
		fmt.Println("error while making new dir at dest")
		return
	}

	// now iterating through all the entry of this dir and recursively calling MoveUtil until every entry reaches to just files

	var wg sync.WaitGroup

	for _, entry := range sourceStructure {
		new_source := filepath.Join(Source, entry.Name()) // adding each entry's name at the end of source path
		wg.Add(1)
		go func() {
			defer wg.Done()
			MoveUtil(new_source, Dest, fileSize)
		}() // recursively calling moveUtil
	}
	wg.Wait()

}
