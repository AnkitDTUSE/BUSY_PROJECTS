package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type CopyJob struct {
	Source string
	Dest   string
}

var (
	wg sync.WaitGroup
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

	// calculating filesize andtime
	fileSize := 0
	bufSize := 1024

	fmt.Println("Enter Buffer size in KBs or hit enter for 1MB(1024KB) default buffsize")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		if val, err := strconv.Atoi(input); err == nil {
			bufSize = val
		}
	}

	// making a Job Channel of 100 jobs at max
	jobs := make(chan CopyJob, 100)

	// initially 15 workers
	workers := 15

	t1 := time.Now()

	for i := 0; i < workers; i++ { // Initializing each worker
		wg.Add(1)

		go Worker(
			jobs,
			&fileSize,
			bufSize,
			&wg,
		)
	}

	MoveUtil(Source, Dest, jobs)

	close(jobs) // unneccesary Step ,just to prevent the receivers from asking jobs from an empty job queue

	wg.Wait()

	elapsed := time.Since(t1)

	fmt.Printf(
		"\nTotal time and total size in MBs => %v %vMB %vBytes\n",
		elapsed,
		float64(fileSize)/(1024*1024),
		fileSize,
	)
}

// bufSize	TransferTime	DataSize			 					TransferRate

// 5kb 		1m4.8684673s 	1332.5364799499512MB 1397265772Bytes -> 	21MB/s

// 10kb 	8.8855202s 		1406.5087223052979MB 1474831290Bytes -> 	156MB/s

// 1MB 		25.752762s 		1945.4457368850708MB 2039947709Bytes -> 	74 MB/s

// 5MB  	5.8288255s 		1410.5730619430542MB 1479093059Bytes -> 	235MB/s

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func main() {
// 	// accepting arguements
// 	args := os.Args

// 	// retrive Source and Dest from args

// 	Source := args[1]
// 	Dest := args[2]

// 	fmt.Println(Source)
// 	fmt.Println(Dest)

// 	// retrive stats from paths
// 	_, err1 := os.Stat(Source)
// 	_, err2 := os.Stat(Dest)

// 	// checking for errors
// 	if err1 != nil {
// 		fmt.Println("error while reading Source")
// 		return
// 	}
// 	if err2 != nil {
// 		fmt.Println("error while reading Dest")
// 		return
// 	}

// 	// calculating filesize andtime
// 	fileSize := 0.0
// 	bufSize := 0
// 	fmt.Println("Enter Buffer size in KBs or hit enter for 0 buffsize")
// 	reader := bufio.NewReader(os.Stdin)
// 	input, _ := reader.ReadString('\n')
// 	input = strings.TrimSpace(input)
// 	if input != "" {
// 		if val, err := strconv.Atoi(input); err == nil {
// 			bufSize = val
// 		}
// 	}
// 	t1 := time.Now()

// 	MoveUtil(Source, Dest, &fileSize, &bufSize) // calling move func here

// 	// var wg sync.WaitGroup // now a go idea here
// 	// wg.Add(1)
// 	// go func() { // using routines -> time goes down to 205milliseconds !!!!!!!!!!!!!!!!!!!!!!!!!!!!
// 	// 	defer wg.Done()
// 	// 	MoveUtil(Source, Dest, &fileSize)
// 	// }()
// 	// wg.Wait()

// 	elapsed := time.Since(t1)

// 	fmt.Printf("Total time and total size in MBs => %v %vMbs %vBytes\n", elapsed, float64(fileSize/(1024.0*1024.0)), fileSize)

// 	// os.Remove(Source) // uncomment this to complete the move Utility

// }

// // 50KB  19s for a 1.5GB file
// // 100KB 10s
// // 1MB   1s
// // 5MB   1s
