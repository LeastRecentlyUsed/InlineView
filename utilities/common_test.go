package utilities

import (
	"math/rand"
	"testing"
	"time"
)

func TestLogFileNameReturnsSameFileName(t *testing.T) {
	expectedName := LogFileName()
	res := LogFileName()

	if res != expectedName {
		t.Error("Failed to create the same log file name within the same day:", expectedName, "vs", res)
	}
}

func BenchmarkLogFileName(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = LogFileName()
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

func BenchmarkSizeAsString(b *testing.B) {
	// seed with varying number at each execution of function
	rand.Seed(time.Now().UnixNano())
	// select random number between 1Kb and 20Mb
	sampleSize := 1024 + rand.Intn(24117248-1024)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = SizeAsString(sampleSize)
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

func BenchmarkHashDataString(b *testing.B) {
	input := "TestHashString"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = HashDataString(input)
	}
}

func TestLogMemStats(t *testing.T) {
	t.Skip("LogMemStats not yet known to have a valid test case")
}
