package ukland

import (
	"testing"
)

func TestPriceFormatCreatesValidStoreData(t *testing.T) {
	line := []string{"{8355F008-C0A9-55C5-E053-6B04A8C0D090}", "102000", "2003-11-25 00:00", "WA5 2PG", "S", "N", "L", "39", "", "ROTHAY DRIVE", "PENKETH", "WARRINGTON", "WARRINGTON", "WARRINGTON", "A", "A"}
	expectedData := `{"pricekey":"hrlRgMLsL3rt9bOZ-HAUiZUZ2Q0=","pricedata":{"postcode":"WA5 2PG","price":"102000","date":"2003-11-25","address":"39 ROTHAY DRIVE PENKETH WARRINGTON","propertytype":"S","newbuild":"N"}}`
	expectedPostcode := "WA5 2PG"

	_, postcode, data := priceFormat(line)

	if postcode != expectedPostcode {
		t.Error("PriceFormat Failed to return the valid postcode:", postcode, "[expected:]", expectedPostcode)
	}
	if data != expectedData {
		t.Error("PriceFormat Failed to return the valid Line:", data, "[expected:]", expectedData)
	}
}

func TestPriceFormatIgnoresRecordWithInvalidCsvDate(t *testing.T) {
	line := []string{"{8355F008-C0A9-55C5-E053-6B04A8C0D090}", "102000", "2003-25-11 00:00", "WA5 2PG", "S", "N", "L", "39", "", "ROTHAY DRIVE", "PENKETH", "WARRINGTON", "WARRINGTON", "WARRINGTON", "A", "A"}
	expectedData := ""

	_, _, data := priceFormat(line)
	if data != expectedData {
		t.Error("PriceFormat Failed to detect an invalid CSV Date")
	}
}

func TestFormatAddressBuildsValidAddressLine(t *testing.T) {
	paon := "24"
	saon := "The Hut"
	street := "Bridges Over"
	locality := "Troubled"
	town := "Waters"
	district := "Problematic"
	County := "North Issues"

	res := formatAddress(paon, saon, street, locality, town, district, County)
	if res != "24 The Hut Bridges Over Troubled Waters Problematic North Issues" {
		t.Error("Failed to build valid address line:", res)
	}
}

func TestFormatAddressBuildsValidAddressLineWithDuplicateValues(t *testing.T) {
	paon := "24"
	saon := "The Hut"
	street := "Bridges Over"
	locality := "Troubled"
	town := "Troubled"
	district := "Waters County"
	County := "Waters County"

	res := formatAddress(paon, saon, street, locality, town, district, County)
	if res != "24 The Hut Bridges Over Troubled Waters County" {
		t.Error("Failed to build valid address line (with duplicates):", res)
	}
}

func TestCanCorrectlyDeterminePostcode(t *testing.T) {
	resNoCode := determinePostcode("", incode)
	if resNoCode != "NOPOSTCODE" {
		t.Log("Failed to determine NOPOSTCODE")
		t.Fail()
	}

	resIncode := determinePostcode("incode outcode", incode)
	if resIncode != "incode" {
		t.Log("Failed to determine INCODE:", resIncode)
		t.Fail()
	}

	resFullcode := determinePostcode("  MK17 9AU ", fullcode)
	if resFullcode != "MK17 9AU" {
		t.Log("Failed to determine FULLCODE:", resFullcode)
		t.Fail()
	}

}
