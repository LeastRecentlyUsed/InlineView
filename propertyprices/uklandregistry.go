package propertyprices

import (
	"InlineView/utilities"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

// use magic reference date values for parsing.
const csvDate = "2006-01-02 15:04"
const noCode = "NOPOSTCODE"

type priceRec struct {
	Postcode     string
	Price        string
	Date         string
	Address      string
	PropertyType string
	NewBuild     string
}

type csvRec struct {
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

	pcSet, err := distinctIncodes(filename)
	if err != nil {
		return err
	}

	sort.Strings(pcSet)
	// for _, v := range pcSet {
	// 	fmt.Println(v)
	// }
	fmt.Println(len(pcSet), "distinct incodes")

	return err
}

// ReadAllRecords scans through the file sequentially and reads/formats each record
func ReadAllRecords(filename string) error {

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open file", filename)
		return err
	}
	defer f.Close()

	lineReader, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	tx := 0
	ix := 0
	for _, line := range lineReader {
		_, v := priceFormat(line)
		utilities.AddPriceRecord(v)

		if ix == 1000 {
			tx = tx + ix
			println(tx)
			ix = 0
		} else {
			ix++
		}
	}

	println("total records processed:", tx)
	return err
}

// priceFormat creates an InlineView specific property price record
func priceFormat(line []string) (key string, value string) {
	addr := formatAddress(line[7], line[8], line[9], line[10], line[11], line[12], line[13])

	date, err := time.Parse(csvDate, line[2])
	if err != nil {
		fmt.Println("Failed to parse CSV date", err)
	}

	var rec priceRec

	if line[3] != "" {
		rec.Postcode = line[3]
	} else {
		rec.Postcode = noCode
	}
	rec.Price = line[1]
	rec.Date = date.Format("2006-01-02")
	rec.Address = addr
	rec.PropertyType = line[4]
	rec.NewBuild = line[5]

	r1, _ := json.Marshal(rec)

	return rec.Postcode, string(r1)
}

func formatAddress(paon string, saon string, street string, locality string, town string, district string, county string) string {
	var address strings.Builder
	if len(strings.TrimSpace(paon)) > 0 {
		address.WriteString(paon)
	}
	if len(strings.TrimSpace(saon)) > 0 {
		address.WriteString(" " + saon)
	}
	if len(strings.TrimSpace(street)) > 0 {
		address.WriteString(" " + street)
	}
	if len(strings.TrimSpace(locality)) > 0 {
		address.WriteString(" " + locality)
	}
	if town != locality && len(strings.TrimSpace(town)) > 0 {
		address.WriteString(" " + town)
	}
	if district != town && len(strings.TrimSpace(district)) > 0 {
		address.WriteString(" " + district)
	}
	if county != district && len(strings.TrimSpace(county)) > 0 {
		address.WriteString(" " + county)
	}

	return address.String()
}

// distinctPostcodes sequentially reads the file and returns a slice of unique
// postcodes
func distinctIncodes(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open file", filename)
		return nil, err
	}
	defer f.Close()

	lineReader, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	postCodes := make(map[string]bool)
	distinctList := []string{}
	var pc string

	for _, line := range lineReader {
		if line[3] != "" {
			//pc = line[3]
			pc = strings.Fields(line[3])[0]
		} else {
			pc = noCode
		}

		// after much pain with sort.SearchStrings, using a hash map to find unique values is less coding
		if _, val := postCodes[pc]; !val {
			postCodes[pc] = true
			distinctList = append(distinctList, pc)
		}
	}
	return distinctList, err
}
