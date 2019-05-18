package ukland

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
const csvDateTime = "2006-01-02 15:04"
const csvDate = "2006-01-02"

type postCodeFormat int

const (
	incode postCodeFormat = 1 + iota
	fullcode
	outcode
)

type priceContainer struct {
	PriceKey string   `json:"pricekey"`
	PriceRec priceRec `json:"pricedata"`
}

type priceRec struct {
	Postcode     string `json:"postcode"`
	Price        string `json:"price"`
	Date         string `json:"date"`
	Address      string `json:"address"`
	PropertyType string `json:"propertytype"`
	NewBuild     string `json:"newbuild"`
}

// inputRec provided for information - this is the structure of the UKLandRegistry Prices CSV row
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

	incodeSet, err := distinctIncodes(filename)
	if err != nil {
		log.Panic(err)
	}

	sort.Strings(incodeSet)

	for _, incode := range incodeSet {
		err = createIncodeStore(filename, incode)
		if err != nil {
			log.Panic(err)
		}
		break
	}

	log.Println(len(incodeSet), "distinct incode stores created")
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

	store := []utilities.StoreData{}

	for _, line := range lineReader {
		thisIncode := determinePostcode(line[3], incode)

		if thisIncode == storeIncode {
			var aRec utilities.StoreData
			aRec.Hash, aRec.Identifier, aRec.Data = priceFormat(line)
			store = append(store, aRec)
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

		// use a map to find unique postcode (incode) values
		if _, val := postCodes[code]; !val {
			postCodes[code] = true
			distinctList = append(distinctList, code)
		}
	}
	return distinctList, err
}

// priceFormat creates an InlineView specific property price record
func priceFormat(line []string) (hash string, postcode string, data string) {
	addr := formatAddress(line[7], line[8], line[9], line[10], line[11], line[12], line[13])

	date, err := time.Parse(csvDateTime, line[2])
	if err != nil {
		log.Println("priceFormat: failed to parse CSV date", err)
		//date = date.AddDate(1900, 01, 01)
		return
	}

	var rec priceContainer

	rec.PriceRec.Postcode = determinePostcode(line[3], fullcode)
	rec.PriceRec.Price = line[1]
	rec.PriceRec.Date = date.Format(csvDate)
	rec.PriceRec.Address = addr
	rec.PriceRec.PropertyType = line[4]
	rec.PriceRec.NewBuild = line[5]

	hashKey := utilities.HashDataString(rec.PriceRec.Postcode + rec.PriceRec.Price + rec.PriceRec.Date + rec.PriceRec.Address)
	rec.PriceKey = hashKey
	r1, _ := json.Marshal(rec)

	return rec.PriceKey, rec.PriceRec.Postcode, string(r1)
}

// formatAddress creates a de-duplicated compact single address line by appending all address fields that are passed in
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
