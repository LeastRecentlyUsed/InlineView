package inlineviewfiles

import (
	"fmt"
	"os"
)

// OpenReadOnlyFile is a public function to request a file to be opened returning the file struct
func OpenReadOnlyFile(name string) *os.File {
	var err error

	handle, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Failed to open file", name, "due to", err)
	} else {
		fmt.Println("Opened file", name)
	}
	return handle
}

// CloseFile is a public function to close an existing opened file.
func CloseFile(handle *os.File) {
	handle.Close()
	fmt.Println("Closed file", handle.Name())
}

// DoesFileExist returns true if the file is found at the named location
func DoesFileExist(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}
	return true
}
