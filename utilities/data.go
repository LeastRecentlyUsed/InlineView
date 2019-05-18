package utilities

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var timeout = 5 * time.Second

// priceKeys is a map of "[priceKey]postcode" used to determine if there are duplicates being created.
var priceKeys = map[string]string{}
var keysMapSize int

// AddPriceStore stores a set of Price records in the price store
func AddPriceStore(storeName string, recList *[]StoreData) (sizeMsg string, err error) {
	storeFile := storeName + ".dat"
	indexFile := storeName + ".idx"
	var writeBytes, writeCount int

	readIndex(indexFile)

	f, err := AppendFile(storeFile)
	if err != nil {
		log.Println("Failed to open/create file for Appends", f)
		return "", err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, rec := range *recList {
		if !foundIndex(rec.Hash, rec.Identifier) {
			n, err := w.WriteString(rec.Data + "\n")
			if err != nil {
				return "", err
			}
			writeBytes = writeBytes + n
			writeCount++
		}
	}
	w.Flush()
	sizeMsg = SizeAsString(writeBytes)
	log.Println("added", writeCount, "records to price store")

	if len(priceKeys) > keysMapSize {
		err = writeIndex(indexFile)
		if err != nil {
			log.Println("Failed to create index file", indexFile)
			return
		}
		log.Println("added records to index file:", strconv.Itoa(len(priceKeys)-keysMapSize))
	}
	return
}

// foundIndex searches for the pricekey hash and if not found generates an index entry
func foundIndex(aHash string, aIdentifier string) bool {
	if _, found := priceKeys[aHash]; found {
		return true
	}
	priceKeys[aHash] = aIdentifier
	return false
}

// readIndex loads the associated index file for a price store.
func readIndex(indexFile string) {
	f, err := OpenFile(indexFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Incode index file not found ", indexFile)
			return
		}
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := strings.Split(scanner.Text(), "|")
		priceKeys[val[0]] = val[1]
		keysMapSize++
	}
}

// writeIndex re-creates the index file if any new entries have been created in the priceKeys map
func writeIndex(indexFile string) (err error) {
	f, err := CreateFile(indexFile)
	if err != nil {
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for i := range priceKeys {
		_, err = w.WriteString(i + "|" + priceKeys[i] + "\n")
	}
	w.Flush()
	return
}
