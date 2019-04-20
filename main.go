package main

import (
	"InlineView/propertyprices"
	"InlineView/utilities"
	"fmt"
	"os"
	"strings"
)

var contentLocation = "http://prod.publicdata.landregistry.gov.uk.s3-website-eu-west-1.amazonaws.com/"
var pricesLatest = "pp-monthly-update-new-version.csv"
var prices2018 = "pp-2018.txt"

func main() {
	args := os.Args
	run := len(args) > 1

	if run && strings.ToLower(args[1]) == "a" {
		getUKLandRegistryData()
	} else if run && strings.ToLower(args[1]) == "b" {
		decodeUKLandRegistryData()
	}
}

func getUKLandRegistryData() {
	fmt.Println("Starting Land Registry Prices File Retrieval...")
	utilities.FetchFileToDisk(contentLocation+pricesLatest, pricesLatest)
	fmt.Println("Completed data retrieval.")
}

func decodeUKLandRegistryData() {
	fmt.Println("Starting Land Registry file split into Post-Codes")
	propertyprices.SplitFileIntoPostcodes(pricesLatest)
	fmt.Println("Completed file split.")
}
