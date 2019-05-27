package model

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

// PriceRecord has the required elements for storing price records (hash of the ID, the ID and the actual data record)
type PriceRecord struct {
	Hash       string
	Identifier string
	Data       string
}

// AddPriceStore stores a number of Price records in the appropriate store
func AddPriceStore(storeName string, prices *[]PriceRecord) (int, error) {
	return 0, nil
}

// readIndex converts a read-stream for an index of a price store to a searchable map
func readIndex(d io.ReadCloser) (map[string]string, int) {
	s := bufio.NewScanner(d)
	priceKeys := map[string]string{}
	keysSize := 0

	for s.Scan() {
		val := strings.Split(s.Text(), "|")
		priceKeys[val[0]] = val[1]
		keysSize++
	}
	d.Close()
	return priceKeys, keysSize
}

//writeIndex converts the map of an index for a price store to a write-stream.
func writeIndex(priceKeys map[string]string) io.Writer {
	b := new(bytes.Buffer)
	w := io.Writer(b)

	for key, val := range priceKeys {
		fmt.Fprintf(w, "%s|%s\n", key, val)
	}
	return w
}

// foundIndex searches for the price key and returns the boolean result
func foundIndex(priceKeys map[string]string, key string) bool {
	if _, found := priceKeys[key]; found {
		return true
	}
	return false
}
