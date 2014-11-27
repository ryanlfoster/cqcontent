package main

import (
	"github.com/fatih/color"
	"os"
	curl "github.com/andelf/go-curl"
	"fmt"
	"io/ioutil"
)

func (dc DownloadCurl) Download() []byte {
	var decoder *Crx
	var foundPackage *Package
	var url string

	decoder = dc.Decoder()
	for _, p := range decoder.Response.Data.Packages.Packages {
		if dc.Package == p.DownloadName {
			foundPackage = &p
			break
		}
	}

	// Throw error if the package isn't found
	if foundPackage == nil {
		color.Red("Could not locate %s on %s", dc.Package, dc.Node)
		os.Exit(1)
	}

	if foundPackage.Group != "" {
		url = fmt.Sprintf("http://%s:8080/etc/packages/%s/%s",
			dc.Node,
			foundPackage.Group,
			dc.Package)
	} else {
		url = fmt.Sprintf("http://%s:8080/etc/packages/%s",
			dc.Node,
			dc.Package)
	}

	easy := curl.EasyInit()
	defer easy.Cleanup()

	// Set options for curl
	easy.Setopt(curl.OPT_USERNAME, dc.Username)
	easy.Setopt(curl.OPT_PASSWORD, dc.Password)
	easy.Setopt(curl.OPT_URL, url)

	// Needs the callback function WriteData to write data to TempFile
	easy.Setopt(curl.OPT_WRITEFUNCTION, WriteData)

	// Store file pointer
	easy.Setopt(curl.OPT_WRITEDATA, dc.Fp)

	// Get to work
	err := easy.Perform()
	Check(err)

	// Read xml into a variable
	output, err := ioutil.ReadFile(dc.Fp.Name())
	Check(err)

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
