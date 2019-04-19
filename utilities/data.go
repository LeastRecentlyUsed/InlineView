package utilities

import (
	"fmt"
	"time"
)

var timeout = 5 * time.Second

// AddPriceRecord stores a single price record in the database
func AddPriceRecord(rec string) bool {
	fmt.Println(rec)
	return true
}
