/*
cq package to work with cq content packages

http://curl.haxx.se/libcurl/c/libcurl-tutorial.html
https://github.com/andelf/go-curl/
*/

package main

import (
	"os"
	"fmt"
)

// callback function for OPT_WRITEFUNCTION. See libcurl docs.
func WriteData(ptr []byte, userdata interface{}) bool {

	file := userdata.(*os.File)
	if _, err := file.Write(ptr); err != nil {
		return false
	} else {
		return true
	}
}

// Progress for uploading
func UploadProgress(dltotal, dlnow, ultotal, ulnow float64, _ interface{}) bool {
	fmt.Printf("Uploading %3.2f%%\r", ulnow/ultotal*100)
	return true
}

// Progress for downloading
func DownloadProgress(dltotal, dlnow, ultotal, ulnow float64, _ interface{}) bool {
	fmt.Printf("Downloading %3.2f%%\r", dlnow/dltotal*100)
	return true
}

type Curl struct {
	Username string
	Password string
}

// Create generic struct to hold values for variable curling
type CurlFp struct {
	Curl
	Fp       *os.File
}

// Augmented struct to hold node value
type ListCurl struct {
	CurlFp
	Node string
}

// Augmented struct to hold content package value and node
type DownloadCurl struct {
	ListCurl
	Package string
}

// Augmented struct to hold pkgpath and node value
type UploadCurl struct {
	ListCurl
	Package string
	Uploaded bool
}
