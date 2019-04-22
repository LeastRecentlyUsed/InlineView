package utilities

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var timeout = 5 * time.Second

// AddPriceStore stores a set of Price record in the database
func AddPriceStore(key string, list *[]string) (sizeMsg string, err error) {
	storename := key + ".dat"
	var writeBytes int

	f, err := os.Create(storename)
	if err != nil {
		fmt.Println("Failed to create file", f)
		return "", err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, rec := range *list {
		n, err := w.WriteString(rec)
		if err != nil {
			return "", err
		}
		writeBytes = writeBytes + n
	}
	w.Flush()
	sizeMsg = SizeAsString(writeBytes)

	return
}
