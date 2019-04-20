package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

var timeout = 5 * time.Second

// AddPriceStore stores a set of Price record in the database
func AddPriceStore(key string, list *[]string) (string, error) {
	storename := key + ".dat"
	filename := GetFullFilePath(storename)
	var writeBytes int

	fileHandle, err := os.Create(filename)
	if err != nil {
		fmt.Println("Failed to create file", filename)
		return "", err
	}
	defer fileHandle.Close()

	w := bufio.NewWriter(fileHandle)
	for _, rec := range *list {
		n, err := w.WriteString(rec)
		if err != nil {
			return "", err
		}
		writeBytes = writeBytes + n
	}
	w.Flush()

	return sizeAsString(writeBytes), nil
}

func sizeAsString(size int) string {
	if size < 1024 {
		return strconv.Itoa(size) + " bytes"
	}
	if size > 1024 && size < 1024*1024 {
		return strconv.Itoa(size/1024) + " Kbytes"
	}
	if size > 1024*1024 && size < 1024*1024*1024 {
		return strconv.Itoa(size/1024/1024) + "  Mbytes"
	}
	return strconv.Itoa(size/1024/1024/1024) + " Gbytes"
}
