package model

import (
	"bufio"
	"io"
	"strings"
)

// readIndex loads the associated index for a price store.
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
