package utilities

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var diskLocation = workingDir()

// FetchFileToDisk retrieves a file from a URL and saves it to a disk location
func FetchFileToDisk(url string, fileName string) error {

	content, err := http.Get(url)
	if err != nil {
		return err
	}
	defer content.Body.Close()

	err = deleteExistingFile(fileName)
	if err != nil {
		return err
	}

	f, err := CreateFile(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, content.Body)
	if err != nil {
		return err
	}

	return nil
}

// OpenFile created to remove repeated code.  Returns a pointer to an existing file
func OpenFile(filename string) (fileHandle *os.File, err error) {
	dataFile := getFullFilePath(filename)
	fileHandle, err = os.Open(dataFile)
	if err != nil {
		return
	}
	return
}

// CreateFile created to remove repeated code.  Returns a pointer to the new file
func CreateFile(filename string) (fileHandle *os.File, err error) {
	dataFile := getFullFilePath(filename)
	fileHandle, err = os.Create(dataFile)
	if err != nil {
		return
	}
	return
}

// AppendFile opens a file for write. May be existing file allowing appends.
func AppendFile(filename string) (fileHandle *os.File, err error) {
	dataFile := getFullFilePath(filename)
	fileHandle, err = os.OpenFile(dataFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	return
}

// ReadFile opens a file and returns a byte array of the contents.
func ReadFile(filename string) (data []byte, err error) {
	dataFile := getFullFilePath(filename)
	data, err = ioutil.ReadFile(dataFile)
	if err != nil {
		return
	}
	return
}

// deleteExistingFile removes a previously created file of the same name from the disk
func deleteExistingFile(fileName string) error {
	delFile := getFullFilePath(fileName)
	if doesFileExist(delFile) {
		err := os.Remove(delFile)
		if err != nil {
			return err
		}
	}
	return nil
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
