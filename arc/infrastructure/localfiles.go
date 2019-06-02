package infrastructure

import (
	"InlineView/arc/entities"
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// FileStore is the type of repository.  This type is used when executing this application from a server host or manually where no other store type
// is available.  Beware, this repo will save data files to the disk location of the executing process (i.e. location of your CLI).
type FileStore struct {
	Prices []entities.PriceRecord
	Index  map[string]string
}

// NewFileStore is the constructor that returns a FileStore object.
func NewFileStore() *FileStore {
	return &FileStore{
		Prices: []entities.PriceRecord{},
		Index:  map[string]string{},
	}
}

// ReadIndex performs a read from a local filesystem.  Retrieves the index entries associated to a price store.
func (fs *FileStore) ReadIndex(filename string) (map[string]string, error) {
	dataFile := getFullFilePath(filename + ".idx")
	d, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer d.Close()

	s := bufio.NewScanner(d)
	indexVal := map[string]string{}

	for s.Scan() {
		val := strings.Split(s.Text(), "|")
		indexVal[val[0]] = val[1]
	}
	return indexVal, err
}

// WriteIndex performs an write to the local filesystem.  Saves the index entries associated to a price store.
func (fs *FileStore) WriteIndex(filename string, index map[string]string) error {
	dataFile := getFullFilePath(filename + ".idx")
	f, err := os.Create(dataFile) // Create will truncate any existing content if the file already exists.
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for key, val := range index {
		fmt.Fprintf(w, "%s|%s\n", key, val)
	}
	w.Flush()
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
