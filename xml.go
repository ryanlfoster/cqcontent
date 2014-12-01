package main

import (
	"fmt"
	curl "github.com/andelf/go-curl"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

// Implementation of Xml()
func (lc ListCurl) Xml() ([]byte, *os.File) {

	// Initialize tempfile and get pointer
	fp, err := ioutil.TempFile("", "cq")

	// Defer closing and removing of file to the end
	defer func() {
		err := fp.Close()
		Check(err)
		os.Remove(fp.Name())
	}()

	easy := curl.EasyInit()
	defer easy.Cleanup()

	// Set URL String
	url := fmt.Sprintf("http://%s:%d/crx/packmgr/service.jsp?cmd=ls",
		lc.Node,
		lc.Port)

	// Set options for curl
	easy.Setopt(curl.OPT_USERNAME, lc.Username)
	easy.Setopt(curl.OPT_PASSWORD, lc.Password)
	easy.Setopt(curl.OPT_URL, url)

	// Needs the callback function WriteData to write data to TempFile
	easy.Setopt(curl.OPT_WRITEFUNCTION, WriteData)

	// Store file pointer
	easy.Setopt(curl.OPT_WRITEDATA, fp)

	// Set connection timeout
	easy.Setopt(curl.OPT_CONNECTTIMEOUT, 10)

	// Get to work
	err = easy.Perform()
	Check(err)

	// Read xml into a variable
	output, err := ioutil.ReadFile(fp.Name())
	Check(err)

	// go-curl doesn't error on authentication failure. However, if we have a
	// zero length string, then authentication or connection probably failed.
	// Notify the user and exit with an error code if that is the case
	if string(output) == "" {
		color.Red(`
The output is a zero length string. This typically indicates authentication
failure

`)
		os.Exit(1)
	}

	return output, fp
}
