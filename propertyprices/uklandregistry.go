package propertyprices

import (
	"InlineView/utilities"
	"encoding/csv"
	"encoding/json"
	"log"
	"sort"
	"strings"
	"time"
)

// use magic reference date values for parsing.
const csvDate = "2006-01-02 15:04"

type postCodeFormat int

const (
	incode postCodeFormat = 1 + iota
	fullcode
	outcode
)

type priceContainer struct {
	priceKey string
	priceRec priceRec
}

type priceRec struct {
	Postcode     string
	Price        string
	Date         string
	Address      string
	PropertyType string
	NewBuild     string
}

type inputRec struct {
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
// into groups (or local stores) based on their postcode.
func SplitFileIntoPostcodes(filename string) {

	pcSet, err := distinctIncodes(filename)
	if err != nil {
		log.Panic(err)
	}

	sort.Strings(pcSet)

	for _, incode := range pcSet {
		err = createIncodeStore(filename, incode)
		if err != nil {
			log.Panic(err)
		}
		break
	}

	log.Println(len(pcSet), "distinct incode stores created")
}

// createIncodeStore builds a sub-set of price records for one incode and stores the values
func createIncodeStore(filename string, storeIncode string) (err error) {
	f, err := utilities.OpenFile(filename)
	if err != nil {
		log.Println("createIncodeStore: Unable to open a file", filename)
		return
	}
	defer f.Close()

	lineReader, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Println("createIncodeStore: cannot read csv", filename)
		return
	}

	store := []string{}

	for _, line := range lineReader {
		thisIncode := determinePostcode(line[3], incode)

		if thisIncode == storeIncode {
			_, v := priceFormat(line)
			store = append(store, v)
		}
	}

	sizeMsg, err := utilities.AddPriceStore(storeIncode, &store)
	if err != nil {
		log.Println("createIncodeStore: cannot AddPriceStore for", storeIncode)
		return err
	}
	log.Println("Stored:", storeIncode, "in", sizeMsg)
	return nil

}

// priceFormat creates an InlineView specific property price record
func priceFormat(line []string) (key string, value string) {
	addr := formatAddress(line[7], line[8], line[9], line[10], line[11], line[12], line[13])

	date, err := time.Parse(csvDate, line[2])
	if err != nil {
		log.Println("priceFormat: failed to parse CSV date", err)
		date = date.AddDate(1900, 01, 01)
	}

	var rec priceContainer

	rec.priceRec.Postcode = determinePostcode(line[3], fullcode)
	rec.priceRec.Price = line[1]
	rec.priceRec.Date = date.Format("2006-01-02")
	rec.priceRec.Address = addr
	rec.priceRec.PropertyType = line[4]
	rec.priceRec.NewBuild = line[5]

	hashKey := utilities.HashDataString(rec.priceRec.Postcode + rec.priceRec.Price + rec.priceRec.Date + rec.priceRec.Address)
	rec.priceKey = hashKey
	r1, _ := json.Marshal(rec)

	return rec.priceKey, string(r1) + "\n"
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

// distinctIncodes sequentially reads the file and returns a slice of unique incodes
func distinctIncodes(filename string) ([]string, error) {
	f, err := utilities.OpenFile(filename)
	if err != nil {
		log.Println("distinctIncodes: Unable to open a file", filename)
		return nil, err
	}
	defer f.Close()

	lineReader, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Println("distinctIncodes: cannot read csv", filename)
		return nil, err
	}

	postCodes := make(map[string]bool)
	distinctList := []string{}
	var code string

	for _, line := range lineReader {
		code = determinePostcode(line[3], incode)

		// after issues using sort.SearchStrings, using a hash map to find unique values is less coding
		if _, val := postCodes[code]; !val {
			postCodes[code] = true
			distinctList = append(distinctList, code)
		}
	}
	return distinctList, err
}

// determinePostcode
// incode = splits the postcode string on a space and returns the first element
// fullcode = returns space trimmed full postcode
func determinePostcode(postcode string, codeStyle postCodeFormat) string {
	if postcode == "" {
		return "NOPOSTCODE"
	}
	if codeStyle == fullcode {
		return strings.TrimSpace(postcode)
	}
	return strings.Fields(postcode)[0]
}
