package main

import (
	"fmt"
	curl "github.com/andelf/go-curl"
	"github.com/fatih/color"
	"os"
)

func (uc *UploadCurl) CheckUploaded() (bool, *Package) {
	pkgFound := false
	var decoder *Crx
	decoder = uc.Decoder()
	var foundPackage *Package
	for _, p := range decoder.Response.Data.Packages.Packages {
		if p.DownloadName == RelPath(uc.Package) {
			pkgFound = true
			uc.Uploaded = true
			foundPackage = &p
			break
		}
	}
	return pkgFound, foundPackage
}

func (uc *UploadCurl) Upload() {

	result, _ := uc.CheckUploaded()
	if result == true {
		color.Yellow("%s is already uploaded. Moving along...", uc.Package)
		return
	}

	result, filePath := FileExists(uc.Package)
	if result == false {
		color.Red("The provided package does not exist at the specified path")
		os.Exit(1)
	}

	easy := curl.EasyInit()
	defer easy.Cleanup()

	// Set URL String
	url := fmt.Sprintf("http://%s:%d/crx/packmgr/service/.json/?cmd=upload",
		uc.Node, uc.Port)

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

	// Set connection timeout
	easy.Setopt(curl.OPT_CONNECTTIMEOUT, 10)

	// Get to work
	err := easy.Perform()
	Check(err)
}

func (uc *UploadCurl) VerifyUpload() {
	// Verify upload
	uploaded, _ := uc.CheckUploaded()
	if uploaded == false {
		color.Red("The package %s failed to upload.", RelPath(uc.Package))
		os.Exit(1)
	}
}
