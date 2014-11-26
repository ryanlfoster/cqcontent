package cq

import (
	"fmt"
	curl "github.com/andelf/go-curl"
	"io/ioutil"
	"os"
	"github.com/fatih/color"
)

// Implementation of Xml()
func (lc ListCurl) Xml() []byte {
	easy := curl.EasyInit()
	defer easy.Cleanup()

	// Set URL String
	url := fmt.Sprintf("http://%s:8080/crx/packmgr/service.jsp?cmd=ls",
		lc.Node)

	// Set options for curl
	easy.Setopt(curl.OPT_USERNAME, lc.Username)
	easy.Setopt(curl.OPT_PASSWORD, lc.Password)
	easy.Setopt(curl.OPT_URL, url)

	// Needs the callback function WriteData to write data to TempFile
	easy.Setopt(curl.OPT_WRITEFUNCTION, WriteData)

	// Store file pointer
	easy.Setopt(curl.OPT_WRITEDATA, lc.Fp)

	// Get to work
	err := easy.Perform()
	check(err)

	// Read xml into a variable
	output, err := ioutil.ReadFile(lc.Fp.Name())
	check(err)

	// go-curl doesn't error on authentication failure. However, if we have a
	// zero length string, then authentication or connection probably failed.
	// Notify the user and exit with an error code if that is the case
	if string(output) == "" {
		color.Red(`
The output is a zero length string. This typically indicates authentication or
connection failure

`)
		os.Exit(1)
	}

	return output
}
