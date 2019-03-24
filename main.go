package main

import (
	"InlineView/propertyprices"
	"InlineView/utilities"
	"fmt"
)

var contentLocation = "http://prod.publicdata.landregistry.gov.uk.s3-website-eu-west-1.amazonaws.com/"
var pricesLatest = "pp-monthly-update-new-version.csv"
var prices2018 = "pp-2018.txt"

func main() {
	decodeUKLandRegistryData()
}

func getUKLandRegistryData() {
	fmt.Println("Starting Land Registry Prices File Retrieval...")
	file := utilities.GetFullFilePath(pricesLatest)
	utilities.FetchFileToDisk(contentLocation+pricesLatest, file)
	fmt.Println("Completed data retrieval.")
}

func decodeUKLandRegistryData() {
	fmt.Println("Starting Land Registry file split into Post-Codes")
	file := utilities.GetFullFilePath(pricesLatest)
	propertyprices.SplitFileIntoPostcodes(file)
	fmt.Println("Completed file split.")
}
