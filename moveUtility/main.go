package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// accepting arguements
	args := os.Args

	// retrive Source and Dest from args

	Source := args[1]
	Dest := args[2]

	fmt.Println(Source)
	fmt.Println(Dest)

	// retrive stats from paths
	_, err1 := os.Stat(Source)
	_, err2 := os.Stat(Dest)

	// checking for errors
	if err1 != nil {
		fmt.Println("error while reading Source")
		return
	}
	if err2 != nil {
		fmt.Println("error while reading Dest")
		return
	}

	// calculating filesize and time
	fileSize := 0.0
	t1 := time.Now()

	MoveUtil(Source, Dest, &fileSize) // calling move func here

	// var wg sync.WaitGroup // now a go idea here
	// wg.Add(1)
	// go func() { // using routines -> time goes down to 205milliseconds !!!!!!!!!!!!!!!!!!!!!!!!!!!!
	// 	defer wg.Done()
	// 	MoveUtil(Source, Dest, &fileSize)
	// }()
	// wg.Wait()

	elapsed := time.Since(t1)

	fmt.Printf("Total time and total size in MBs => %v %vMbs %vBytes\n", elapsed, float64(fileSize/(1024.0*1024.0)), fileSize)

	// os.Remove(Source) // uncomment this to complete the move Utility
	
}
