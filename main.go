package main

import (
	"InlineView/propertyprices"
	"InlineView/utilities"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var contentLocation = "http://prod.publicdata.landregistry.gov.uk.s3-website-eu-west-1.amazonaws.com/"
var pricesLatest = "pp-monthly-update-new-version.csv"
var prices2018 = "pp-2018.txt"

func main() {
	// set logging on
	f, err := utilities.AppendFile(logFileName())
	if err != nil {
		fmt.Println("Unable to create log file in main")
	}
	defer f.Close()
	log.SetOutput(f)

	// controlling logic
	args := os.Args
	run := len(args) > 1

	if run && strings.ToLower(args[1]) == "a" {
		getUKLandRegistryData()
	} else if run && strings.ToLower(args[1]) == "b" {
		decodeUKLandRegistryData()
	}
}

func getUKLandRegistryData() {
	log.Println("Starting Land Registry Prices File Retrieval...")
	utilities.FetchFileToDisk(contentLocation+pricesLatest, pricesLatest)
	log.Println("Completed data retrieval.")
}

func decodeUKLandRegistryData() {
	log.Println("Starting Land Registry file split into Post-Codes")
	propertyprices.SplitFileIntoPostcodes(pricesLatest)
	log.Println("Completed file split.")
}

func logFileName() string {
	today := time.Now()
	logDate := today.Format("2006-01-02")
	return logDate + "-inlineview.log"
}
