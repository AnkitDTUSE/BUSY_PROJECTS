package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

type CopyJob struct {
	Source string
	Dest   string
}

var (
	MUTEX sync.Mutex
)

func Worker(
	jobs <-chan CopyJob,
	fileSize *int,
	bufSize int,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	buffer := make([]byte, bufSize*1024)

	for job := range jobs {

		jobStat, err := os.Stat(job.Source)
		if err != nil {
			fmt.Println(err)
			continue // this continue here is imp, as if I write Return here this will lead to the death of a worker (as returning means this a particular worker is now quit the worker func)
		}

		srcHandle, err2 := os.Open(job.Source)
		if err2 != nil {
			fmt.Println(err)
			continue
		}

		dstHandle, err3 := os.Create(job.Dest)
		if err3 != nil {
			srcHandle.Close() // if Dest is Faulty then close source File
			fmt.Println(err)
			continue
		}

		_, err4 := io.CopyBuffer(dstHandle, srcHandle, buffer)

		srcHandle.Close()
		dstHandle.Close()

		if err4 != nil {
			fmt.Println(err)
			continue
		}

		// err5 = os.Remove(job.Source)
		// if err5 != nil {
		// 	fmt.Println("failed to delete:", job.Source)
		// }

		MUTEX.Lock()
		(*fileSize) += int(jobStat.Size())
		MUTEX.Unlock()

	}
}

func MoveUtil(source, dest string, jobs chan<- CopyJob) {

	sourceStat, err := os.Stat(source)
	if err != nil {
		return
	}

	basePath := filepath.Base(source)

	// this is the all task we need to do in MoveUtil , create a new job every time we hit a file

	if !sourceStat.IsDir() {

		jobs <- CopyJob{
			Source: source,
			Dest:   filepath.Join(dest, basePath),
		}

		fmt.Println("File written")
		return
	}

	sourceStructure,_ := os.ReadDir(source)

	newDest := filepath.Join(dest, basePath)
	err = os.Mkdir(newDest, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, entry := range sourceStructure {

		MoveUtil(
			filepath.Join(source, entry.Name()),
			newDest,
			jobs,
		)
	}
}
















// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"path/filepath"
// 	"sync"
// )

// var (
// 	updateMutex sync.Mutex
// )

// func MoveUtil(Source, Dest string, fileSize *float64, bufSize *int) {
// 	// retrive base address of Source
// 	basePath := filepath.Base(Source)
// 	SourceStat, _ := os.Stat(Source)

// 	// check if Source is a file or dir
// 	if !SourceStat.IsDir() {

// 		// locking fileSize wo consistent values
// 		updateMutex.Lock()
// 		(*fileSize) += float64(SourceStat.Size())
// 		updateMutex.Unlock()

// 		// THE COMMENTED CODE BELOW IS OF WRITING WHOLE FILE AT ONCE

// 		// // as source is a file -> reading its data
// 		// sourceData, err := os.ReadFile(Source)

// 		// // adding fileSize

// 		// if err != nil {
// 		// 	fmt.Println("error while reading SourceFile")
// 		// 	return
// 		// }

// 		// // adding filename at the end of Dest
// 		Dest := filepath.Join(Dest, basePath) // uncomment this line as it is crucial for both (write whole file at once or write using buffer)

// 		// //writing file
// 		// err2 := os.WriteFile(Dest, sourceData, 0644)

// 		// if err2 != nil {
// 		// 	fmt.Println("error while writing file at Dest")
// 		// 	return
// 		// }

// 		// fmt.Println("File written")
// 		// return

// 		// FROM NOW ON WE WRITE FILE IN CHUNKS OF 2MB
// 		// opening file at source

// 		sourceFile, err := os.Open(Source)
// 		if err != nil {
// 			fmt.Println("Error while open file at source")
// 			return
// 		}
// 		defer sourceFile.Close() //closing file

// 		destFile, err := os.Create(Dest)
// 		if err != nil {
// 			fmt.Println("Error while creating file at Dest")
// 			return
// 		}
// 		defer destFile.Close()
// 		buffer := make([]byte, (*bufSize)*1024) // byte is an alias for uint8 (0 - 255)

// 		_, err2 := io.CopyBuffer(destFile, sourceFile, buffer) // _ is for a varible which copybuffer return as an INT8 which is eq to the size of data written

// 		if err2 != nil {
// 			fmt.Println("Error while writting at dest")
// 			return
// 		}
// 		return
// 	}

// 	// as now source is a Dir -> reading the sturture of this dir
// 	sourceStructure, _ := os.ReadDir(Source)

// 	// making required Dir at Dest
// 	Dest = filepath.Join(Dest, basePath)
// 	err := os.Mkdir(Dest, 0755)

// 	if err != nil {
// 		fmt.Println("error while making new dir at dest")
// 		return
// 	}

// 	// now iterating through all the entry of this dir and recursively calling MoveUtil until every entry reaches to just files

// 	var wg sync.WaitGroup

// 	for _, entry := range sourceStructure {
// 		new_source := filepath.Join(Source, entry.Name()) // adding each entry's name at the end of source path
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			MoveUtil(new_source, Dest, fileSize, bufSize)
// 		}() // recursively calling moveUtil
// 	}
// 	wg.Wait()

// }
