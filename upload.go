package main

import (
	curl "github.com/andelf/go-curl"
	"github.com/fatih/color"
	"os"
	"fmt"
	"path/filepath"
)

func (uc UploadCurl) CheckUploaded() bool {
	pkgFound := false
	var decoder *Crx
	decoder = uc.Decoder()
	for _, p := range decoder.Response.Data.Packages.Packages {
		if p.DownloadName == filepath.Base(uc.Package) {
			pkgFound = true
			uc.Uploaded = true
			break
		}
	}
	return pkgFound
}

// Move to install.go when working on that
//func (uc UploadCurl) CheckInstalled() bool {
//	pkgFound := false
//	var decoder *Crx
//	decoder = Decoder()
//	for _, p := range decoder.Response.Data.Packages.Packages {
//		if p.DownloadName.(string) == filepath.Base(uc.Package.(string)) {
//			if p.LastUnpackedBy != nil {
//				pkgFound = true
//				break
//			}
//		}
//	}
//	return found
//}


func (uc UploadCurl) Upload() {

	if uc.CheckUploaded() == true {
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
