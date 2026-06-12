package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func cacheUtil(mpp *map[KEY]VALUE) {
	db, err := os.OpenFile("D:\\BTECH\\INTERN_BUSY\\Projects\\CacheUtility\\data.csv", os.O_APPEND|os.O_RDWR, 0644)

	if err != nil {
		fmt.Println("error in db reading:", err)
		return
	}

	defer db.Close()

	dbInfo, _ := db.Stat()

	Scanner := bufio.NewScanner(os.Stdin)

	if dbInfo.Size() != 0 {
		csvReader := csv.NewReader(db)
		fmt.Println("\nHIT 'P' to enter persistent mode or 'C' to clear cache")

		Scanner.Scan()

		ch := strings.TrimSpace(Scanner.Text())

		switch ch {
		case "P":
			fmt.Println("Retriving Prior data")
			dbData, _ := csvReader.ReadAll()

			for _, row := range dbData {
				(*mpp)[row[0]] = row[1]
			}

		case "C":
			db.Close()

			cleanDb, err := os.OpenFile("D:\\BTECH\\INTERN_BUSY\\Projects\\CacheUtility\\data.csv", os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Error in opening file for cleaning", err)
				return
			}

			err2 := cleanDb.Truncate(0)

			if err2 != nil {
				fmt.Println("Error while clearing the cache", err)
			}
			fmt.Println("Clearing Cache")

		default:
			fmt.Println("invalid input... retry")
		}

	}
	dbReopen, err := os.OpenFile("D:\\BTECH\\INTERN_BUSY\\Projects\\CacheUtility\\data.csv", os.O_APPEND|os.O_RDWR, 0644)

	if err != nil {
		fmt.Println("error in db reading:", err)
		return
	}

	defer dbReopen.Close()

	fmt.Println("\nUse this format to GET or SET data in cache: SET <KEY> <VALUE> or GET <KEY>   EXIT to exit the cache")

	for {
		fmt.Print("\n->")

		if !Scanner.Scan() {
			break
		}
		if err := Scanner.Err(); err != nil {
			fmt.Println("error while scaning")
			continue
		}

		args := strings.Fields(Scanner.Text())

		if len(args) == 0 {
			fmt.Println("Enter something then hit enter")
			continue
		}

		switch args[0] {
		case "SET":
			if len(args) != 3 {
				fmt.Println("format to set value is SET <KEY> <VALUE>")
				continue
			}
			(*mpp)[args[1]] = args[2]

			recordWriter := csv.NewWriter(dbReopen)
			newRow := []string{args[1], args[2]}
			if err := recordWriter.Write(newRow); err != nil {
				fmt.Println("error while writing")
				return
			}
			recordWriter.Flush()
			fmt.Println("Value Setted")

		case "GET":
			val,ok:=(*mpp)[args[1]]
			if !ok{
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
