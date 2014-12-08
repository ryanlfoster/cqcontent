package main

import (
	curl "github.com/andelf/go-curl"
	"github.com/fatih/color"
	"fmt"
	"os"
)

func (dc *DeleteCurl) Delete() {

	// Check to see if the content package is even there
	result, pkg := dc.CheckUploaded()
	if result == false {
		color.Yellow("%s is already deleted. Moving along...", dc.Package)
		return
	}

	easy := curl.EasyInit()
	defer easy.Cleanup()

	var url string
	if pkg.Group != "" {
		url = fmt.Sprintf("http://%s:%d/crx/packmgr/service/.json/etc/packages/%s/%s?cmd=delete",
			dc.Node,
			dc.Port,
			pkg.Group,
			dc.Package)
	} else {
		url = fmt.Sprintf("http://%s:%d/crx/packmgr/service/.json/etc/packages/%s?cmd=delete",
			dc.Node,
			dc.Port,
			dc.Package)
	}

	// Set options for curl
	easy.Setopt(curl.OPT_USERNAME, dc.Username)
	easy.Setopt(curl.OPT_PASSWORD, dc.Password)
	easy.Setopt(curl.OPT_URL, url)

	// Print upload progress
	easy.Setopt(curl.OPT_NOPROGRESS, true)

	// Set -X POST
	easy.Setopt(curl.OPT_CUSTOMREQUEST, "POST")

	// Set connection timeout
	easy.Setopt(curl.OPT_CONNECTTIMEOUT, 10)

	// Get to work
	err := easy.Perform()
	Check(err)

}

func (dc *DeleteCurl) VerifyDelete() {
	result, _ := dc.CheckUploaded()
	if result == true {
		color.Red("The package %s failed to delete.", dc.Package)
		os.Exit(1)
	}
}
