package main

import (
	"bufio"
	"bytes"
	"github.com/fatih/color"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"fmt"
)

// Generic error handling function
func Check(err error) {
	if err != nil {
		color.Red("ERROR: %v\n", err)
		os.Exit(1)
	}
}

func StrToInt(s string) int64 {
	num, err := strconv.ParseInt(s, 10, 0)
	Check(err)
	return num
}

// callback function for OPT_WRITEFUNCTION. See libcurl docs.
func WriteData(ptr []byte, userdata interface{}) bool {
	// Create ptr reference to tempfile to which we are writing
	writePtr := userdata.(*os.File)

	// Create a bufio reader and writer
	reader := bufio.NewReader(bytes.NewReader(ptr))
	writer := bufio.NewWriter(writePtr)

	// Create a buffer of 1024 bytes
	buf := make([]byte, 1024)

	// Loop until break
	for {
		// read 1024 bytes and return actual number of bytes read into n
		nin, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		// If we have nothing left to read, then break
		if nin == 0 {
			break
		}

		// write a buffer data from 0 to number of bytes read
		writer.Write(buf[:nin])
		if err != nil {
			panic(err)
		}

	}

	// Flush the writer after we're done writing
	err := writer.Flush()
	if err != nil {
		panic(err)
	}

	return true
}

func IsRelative(str string) bool {
	regex, err := regexp.Compile("/")
	if err != nil {
		color.Red("%v\n", err)
		os.Exit(1)
	}
	match := regex.FindString(str)

	if match == "" {
		return true
	} else {
		return false
	}
}

func FullPath(str string) string {
	if IsRelative(str) == true {
		absPath, err := filepath.Abs(str)
		Check(err)
		return absPath
	} else {
		absPath := str
		return absPath
	}
}

func RelPath(str string) string {
	if IsRelative(str) == false {
		relPath := filepath.Base(str)
		return relPath
	} else {
		relPath := str
		return relPath
	}
}

func FileExists(str string) (bool, string) {
	absPath := FullPath(str)
	if _, err := os.Stat(absPath); err == nil {
		return true, absPath
	} else {
		return false, ""
	}
}

func UploadProgress(dltotal, dlnow, ultotal, ulnow float64, _ interface{}) bool {
	fmt.Printf("Uploading %3.2f%%\r", ulnow/ultotal*100)
	return true
}

func DownloadProgress(dltotal, dlnow, ultotal, ulnow float64, _ interface{}) bool {
	fmt.Printf("Downloading %3.2f%%\r", dlnow/dltotal*100)
	return true
}
