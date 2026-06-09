package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func MoveUtil(Source, Dest string, fileSize *float64) {
	// retrive base address of Source
	basePath := filepath.Base(Source)
	SourceStat, _ := os.Stat(Source)

	// check if Source is a file or dir
	if !SourceStat.IsDir() {
		// as source is a file -> reading its data
		sourceData, err := os.ReadFile(Source)

		// adding fileSize
		(*fileSize) += float64(SourceStat.Size())

		if err != nil {
			fmt.Println("error while reading SourceFile")
			return
		}

		// adding filename at the end of Dest
		Dest := filepath.Join(Dest, basePath)

		//writing file
		err2 := os.WriteFile(Dest, sourceData, 0644)

		if err2 != nil {
			fmt.Println("error while writing file at Dest")
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
	for _, entry := range sourceStructure {
		new_source := filepath.Join(Source, entry.Name()) // adding each entry's name at the end of source path
		MoveUtil(new_source, Dest, fileSize)              // recursively calling moveUtil
	}

}
