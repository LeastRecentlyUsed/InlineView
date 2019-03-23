package uklandregistry

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

// use magic reference date values for parsing.
const csvDate = "2006-01-02 15:04"

type priceRec struct {
	Postcode     string
	Price        string
	Date         string
	Address      string
	PropertyType string
	NewBuild     string
}

type csvRecord struct {
	Key          string
	Price        string
	Date         string
	Postcode     string
	PropertyType string
	NewBuild     string
	Duration     string
	Paon         string
	Saon         string
	Street       string
	Locality     string
	Town         string
	District     string
	County       string
	PPDCategory  string
	RecordStatus string
}

// SplitFileIntoPostcodes takes a downloaded UK Land Registry file and splits the property price records
// into groups based on their postcode.
func SplitFileIntoPostcodes(filename string) error {

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open file", filename)
		return err
	}
	defer f.Close()

	lineReader, csvErr := csv.NewReader(f).ReadAll()
	if csvErr != nil {
		return csvErr
	}

	for _, line := range lineReader {
		record := priceFormat(line)
		fmt.Println(record)
	}
	return csvErr
}

// priceFormat creates a JSON string for a property price record
func priceFormat(line []string) priceRec {
	addr := formatAddress(line[7], line[8], line[9], line[10], line[11], line[12], line[13])

	date, err := time.Parse(csvDate, line[2])
	if err != nil {
		panic(err)
	}

	newRecord := priceRec{
		Postcode:     line[3],
		Price:        line[1],
		Date:         date.Format("2006-01-02"),
		Address:      addr,
		PropertyType: line[4],
		NewBuild:     line[5],
	}
	return newRecord
}

func formatAddress(paon string, saon string, street string, locality string, town string, district string, county string) string {

	return address
}
