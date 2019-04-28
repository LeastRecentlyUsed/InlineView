package utilities

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

var timeout = 5 * time.Second

type indexEntry struct {
	key      string
	postcode string
}

var keys = indexEntry{}

// AddPriceStore stores a set of Price records in the database
func AddPriceStore(key string, list *[]string) (sizeMsg string, err error) {
	storeName := key + ".dat"
	indexName := key + ".idx"
	var writeBytes int

	readIndex(indexName)

	f, err := CreateFile(storeName)
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

func readIndex(indexName string) {
	data, err := ReadFile(indexName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Incode index file not found ", indexName)
			return
		}
	}
	// this expects a 2-part json structure {pricekey, postcode}
	err = json.Unmarshal([]byte(data), &keys)

}
