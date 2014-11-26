package main

import (
	curl "github.com/andelf/go-curl"
	"github.com/fatih/color"
	"os"
	"fmt"
)

func (uc UploadCurl) Upload() {

	result, filePath := FileExists(uc.Package)
	if result == false {
		color.Red("The provided package does not exist at the specified path")
		os.Exit(1)
	}

	easy := curl.EasyInit()
	defer easy.Cleanup()

	// Set URL String
	url := fmt.Sprintf("http://%s:8080/crx/packmgr/service/.json/?cmd=upload",
		uc.Node)

	// Set options for curl
	easy.Setopt(curl.OPT_USERNAME, uc.Username)
	easy.Setopt(curl.OPT_PASSWORD, uc.Password)
	easy.Setopt(curl.OPT_URL, url)

	// Create a new form
	form := curl.NewForm()
	form.AddFile("package", filePath)

	// Add form to Setup
	easy.Setopt(curl.OPT_HTTPPOST, form)

	// Print upload progress
	easy.Setopt(curl.OPT_NOPROGRESS, false)

	// Setup Progress
	easy.Setopt(curl.OPT_PROGRESSFUNCTION, UploadProgress)

	// Get to work
	err := easy.Perform()
	Check(err)
}
