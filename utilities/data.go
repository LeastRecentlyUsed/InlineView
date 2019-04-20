package utilities

import (
	"fmt"
	"time"
)

var timeout = 5 * time.Second

// AddPriceStore stores a set of Price record in the database
func AddPriceStore(key string, list *[]string) bool {
	fmt.Println(key, len(*list))
	return true
}
