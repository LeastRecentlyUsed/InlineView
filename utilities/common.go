package utilities

import (
	"log"
	"runtime"
	"strconv"
	"time"
)

// LogFileName formats and returns the name of the log file to use in this execution
func LogFileName() string {
	today := time.Now()
	logDate := today.Format("2006-01-02")
	return logDate + "-inlineview.log"
}

// LogMemStats writes basic memory and gc stats to the log file.
func LogMemStats() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	log.Print("Total Memory Allocated: ", SizeAsString(int(mem.TotalAlloc)), ", Number of GCs: ", mem.NumGC)
}

// SizeAsString returns a more human readable size description from a number of bytes
func SizeAsString(size int) string {
	if size < 1024 {
		return strconv.Itoa(size) + " bytes"
	}
	if size > 1024 && size < 1024*1024 {
		return strconv.Itoa(size/1024) + " Kbytes"
	}
	if size > 1024*1024 && size < 1024*1024*1024 {
		return strconv.Itoa(size/1024/1024) + " Mbytes"
	}
	return strconv.Itoa(size/1024/1024/1024) + " Gbytes"
}
