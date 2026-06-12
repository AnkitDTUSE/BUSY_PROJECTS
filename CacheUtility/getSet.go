package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// preRequiste is to make a CSV file to hold the data of map after the program got terminated

func cacheUtil(mpp *map[KEY]VALUE) {

	// open the csv file in append and readWrite mode
	db, err := os.OpenFile("D:\\BTECH\\INTERN_BUSY\\Projects\\CacheUtility\\data.csv", os.O_APPEND|os.O_RDWR, 0644)

	if err != nil {
		fmt.Println("error in db reading:", err)
		return
	}

	defer db.Close() // close the file

	dbInfo, _ := db.Stat() // collect info about the file (Size particularly)

	Scanner := bufio.NewScanner(os.Stdin) // decalre a Scanner to scan Inputs from terminal

	// checking File size
	if dbInfo.Size() != 0 {
		// filesize is non zero
		
		// create a CSV reader on file
		csvReader := csv.NewReader(db)
		fmt.Println("\nHIT 'P' to enter persistent mode or 'C' to clear cache")

		// using the scanner for scanninf input either P or C
		Scanner.Scan()

		ch := strings.TrimSpace(Scanner.Text())

		switch ch { // switch on the basis of input if P then load file data in MPP otherwise Truncate file on C
		case "P":
			fmt.Println("Retriving Prior data")
			dbData, _ := csvReader.ReadAll()

			for _, row := range dbData {
				(*mpp)[row[0]] = row[1]
			}

		case "C":
			db.Close() // closing the previous file handle as OS prevents File Truncation in APPEND mode

			// reOpen file in W only and truncate mode
			cleanDb, err := os.OpenFile("D:\\BTECH\\INTERN_BUSY\\Projects\\CacheUtility\\data.csv", os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Error in opening file for cleaning", err)
				return
			}
			defer cleanDb.Close()

			err2 := cleanDb.Truncate(0) // clear data

			if err2 != nil {
				fmt.Println("Error while clearing the cache", err)
			}
			fmt.Println("Clearing Cache")

		default:
			fmt.Println("invalid input... retry")
		}

	}

	// reopening file here Because if user hit C then the file is closed already
	dbReopen, err := os.OpenFile("D:\\BTECH\\INTERN_BUSY\\Projects\\CacheUtility\\data.csv", os.O_APPEND|os.O_RDWR, 0644)

	if err != nil {
		fmt.Println("error in db reading:", err)
		return
	}

	defer dbReopen.Close()

	fmt.Println("\nUse this format to GET or SET data in cache: SET <KEY> <VALUE> or GET <KEY>   EXIT to exit the cache")

	for {
		fmt.Print("\n->")
		 
		// scan for input , if null then break loop
		if !Scanner.Scan() {
			break
		}
		if err := Scanner.Err(); err != nil {
			fmt.Println("error while scaning")
			continue
		}

		args := strings.Fields(Scanner.Text()) // extract args from the input
		if len(args) == 0 {
			fmt.Println("Enter something then hit enter")
			continue
		}

		switch args[0] { // switch on the baais of args
		case "SET":
			if len(args) != 3 {
				fmt.Println("format to set value is SET <KEY> <VALUE>")
				continue
			}
			(*mpp)[args[1]] = args[2]

			recordWriter := csv.NewWriter(dbReopen) // create a writer to write new record in File
			newRow := []string{args[1], args[2]}
			if err := recordWriter.Write(newRow); err != nil {
				fmt.Println("error while writing")
				return
			}
			recordWriter.Flush()
			fmt.Println("Value Setted")

		case "GET":
			val, ok := (*mpp)[args[1]]
			if !ok {
				fmt.Println("Key dont exists :( ")
				continue
			}
			fmt.Printf("key: %v Value: %v", args[1], val)

		case "EXIT":
			fmt.Println("Saving and exiting")
			return
		default:
			fmt.Println("Invalid inputs... retry")
		}
	}

}
