package utilities

import "testing"

func TestLogFileNameReturnsSameFileName(t *testing.T) {
	expectedName := LogFileName()
	res := LogFileName()

	if res != expectedName {
		t.Error("Failed to create the same log file name within the same day:", expectedName, "vs", res)
	}
}
func TestSizeAsStringShowsBytes(t *testing.T) {
	res := SizeAsString(124)
	expectedString := "124 bytes"

	if res != expectedString {
		t.Error("Failed to show size as string in Bytes:", res)
	}
}
func TestSizeAsStringShowsKilobytes(t *testing.T) {
	res := SizeAsString(1024)
	expectedString := "1 Kbytes"

	if res != expectedString {
		t.Error("Failed to show size as string in Kilobytes:", res)
	}
}
func TestSizeAsStringShowsMegabytes(t *testing.T) {
	res := SizeAsString(1024 * 1024)
	expectedString := "1 Mbytes"

	if res != expectedString {
		t.Error("Failed to show size as string in Megabytes:", res)
	}
}
func TestSizeAsStringShowsGigabytes(t *testing.T) {
	res := SizeAsString(1024 * 1024 * 1024)
	expectedString := "1 Gbytes"

	if res != expectedString {
		t.Error("Failed to show size as string in Gigabytes", res)
	}
}
func TestHashDataStringCreatesSameTestHash(t *testing.T) {
	res := HashDataString("TestHashString")
	expectedHash := "RX9f5OrTt98VhErEp6NhJ9oz7VQ="

	if res != expectedHash {
		t.Error("Failed to produce expected hash string:", res)
	}
}
