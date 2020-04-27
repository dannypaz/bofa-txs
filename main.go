// Reads a csv transaction file from BofA and converts shows you debits + total
// expenditure.
// This can be used to see an overview of your total spend and see how much you
// are spending per month
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("transactions.csv")
	if err != nil {
		fmt.Println("We received an error reading the files")
		log.Fatal(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	var debit float64

	// Read the first line which is headings... we'll just throw it out
	reader.Read()

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		if line[5] != "" {
			d, err := strconv.ParseFloat(line[5], 64)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Debit:", line[3], line[5])
			debit += d
		}
	}

	fmt.Println("Total:", debit)
}
