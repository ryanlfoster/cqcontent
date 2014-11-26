/*
cq package to work with cq content packages

http://curl.haxx.se/libcurl/c/libcurl-tutorial.html
https://github.com/andelf/go-curl/
*/

package cq

import (
	"fmt"
	"os"
)

// Generic exception handling in event of failure
func check(err error) {
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}

// callback function for OPT_WRITEFUNCTION. See libcurl docs.
func WriteData(ptr []byte, userdata interface{}) bool {

	file := userdata.(*os.File)
	if _, err := file.Write(ptr); err != nil {
		return false
	} else {
		return true
	}
}

// Create generic struct to hold values for variable curling
type Curl struct {
	Fp       *os.File
	Username string
	Password string
}

// Augmented struct to hold node value
type ListCurl struct {
	Curl
	Node string
}

// Augmented struct to hold content package value and node
type DownloadCurl struct {
	Curl
	Node string
	Package string
}

// Interface for list of methods that a Curler should support. See Decoder
// implementation for information on the Crx type
type Curler interface {
	Xml() []byte
	Decoder() *Crx
	List()
	Download()
}
