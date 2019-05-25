package data

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

// IndexReader opens the local data file containing the index entries for a data store and returns a reader.
func IndexReader(fileStore string) io.ReadCloser {
	indexFile := getFullFilePath(fileStore + ".idx")

	fileHandle, err := os.Open(indexFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("File not found error:", indexFile)
		}
		log.Print("Failed to open local index file:", indexFile)
		return nil
	}
	return fileHandle
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

// getFullFilePath returns a file name with the correct OS path prefixed for this utility type
func getFullFilePath(filename string) string {
	dir, err := os.Getwd()
	if err != nil {
		dir = workingDir()
	}
	return dir + string(os.PathSeparator) + filename
}

// workingDir finds the disk location of the current module (pwd).  If no dir can be found, exit with fatal error.
func workingDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln(err)
	}
	return dir
}
