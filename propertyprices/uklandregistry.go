package propertyprices

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// use magic reference date values for parsing.
const csvDate = "2006-01-02 15:04"
const noCode = "NOPOSTCODE"

type storageRec struct {
	Key   string
	Value priceRec
}

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
		_ = record
	}
	return csvErr
}

// priceFormat creates an InlineView specific property price record
func priceFormat(line []string) storageRec {
	addr := formatAddress(line[7], line[8], line[9], line[10], line[11], line[12], line[13])

	date, err := time.Parse(csvDate, line[2])
	if err != nil {
		fmt.Println("Failed to parse CSV date", err)
	}

	var rec storageRec

	if line[3] != "" {
		rec.Key = line[3]
	} else {
		rec.Key = noCode
	}
	rec.Value.Postcode = line[3]
	rec.Value.Price = line[1]
	rec.Value.Date = date.Format("2006-01-02")
	rec.Value.Address = addr
	rec.Value.PropertyType = line[4]
	rec.Value.NewBuild = line[5]

	r1, _ := json.Marshal(rec)
	fmt.Println(string(r1))
	return rec
}

func formatAddress(paon string, saon string, street string, locality string, town string, district string, county string) string {
	var address strings.Builder
	if len(strings.TrimSpace(paon)) > 0 {
		address.WriteString(paon)
	}
	if len(strings.TrimSpace(saon)) > 0 {
		address.WriteString(", " + saon)
	}
	if len(strings.TrimSpace(street)) > 0 {
		address.WriteString(" " + street)
	}
	if len(strings.TrimSpace(locality)) > 0 {
		address.WriteString(", " + locality)
	}
	if town != locality && len(strings.TrimSpace(town)) > 0 {
		address.WriteString(", " + town)
	}
	if district != town && len(strings.TrimSpace(district)) > 0 {
		address.WriteString(", " + district)
	}
	if county != district && len(strings.TrimSpace(county)) > 0 {
		address.WriteString(", " + county)
	}

	return address.String()
}
