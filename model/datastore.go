package model

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

// readIndex converts a read-stream for an index of a price store to a searchable map
func readIndex(f io.ReadCloser) (map[string]string, int) {
	s := bufio.NewScanner(f)
	priceKeys := map[string]string{}
	keysSize := 0

	for s.Scan() {
		val := strings.Split(s.Text(), "|")
		priceKeys[val[0]] = val[1]
		keysSize++
	}
	f.Close()
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
