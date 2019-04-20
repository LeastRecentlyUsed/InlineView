package utilities

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var diskLocation = workingDir()

// FetchFileToDisk retrieves a file from a URL and saves it to a disk location
func FetchFileToDisk(url string, fileName string) error {

	content, getErr := http.Get(url)
	if getErr != nil {
		fmt.Println("Failed to fetch url", url)
		return getErr
	}
	defer content.Body.Close()

	deleteExistingFile(fileName)

	fileHandle, createErr := os.Create(fileName)
	if createErr != nil {
		fmt.Println("Failed to create file", fileName)
		return createErr
	}
	defer fileHandle.Close()

	_, copyErr := io.Copy(fileHandle, content.Body)
	if copyErr != nil {
		fmt.Println("Failed to copy URL content to File")
		return copyErr
	}
	return copyErr
}

// deleteExistingFile removes a previously created file of the same name from the disk
func deleteExistingFile(fileName string) bool {
	if doesFileExist(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			fmt.Println("Failed to delete existing file", fileName, "Error:", err)
			return false
		}
		return true // remove ok
	}
	return true // no existing file
}

// doesFileExist returns true if the file exists
func doesFileExist(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	// to return false if not a file we must negate the isDir result
	return !info.IsDir()
}

// GetFullFilePath returns a file name with the correct OS path prefixed for this utility type
func GetFullFilePath(filename string) string {
	dir, err := os.Getwd()
	if err != nil {
		dir = workingDir()
	}
	return dir + string(os.PathSeparator) + filename
}

// workingDir finds the disk location of the current module (pwd)
func workingDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return dir
}
