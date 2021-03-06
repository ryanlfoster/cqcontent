package main

import (
	"fmt"
	curl "github.com/andelf/go-curl"
	"github.com/fatih/color"
	"os"
	"path/filepath"
)

func (ic *InstallCurl) CheckInstalled() bool {
	pkgFound := false
	var decoder *Crx
	decoder = ic.Decoder()
	for _, p := range decoder.Response.Data.Packages.Packages {
		if p.DownloadName == filepath.Base(ic.Package) {
			if p.LastUnpackedBy != "null" {
				pkgFound = true
				ic.Installed = true
				break
			}
		}
	}
	return pkgFound
}

func (ic *InstallCurl) Install() {
	var url string
	fileName := RelPath(ic.Package)

	fmt.Printf("Installing %s to %s\n", ic.Package, ic.Node)

	result, pkg := ic.CheckUploaded()
	if result == false {
		color.Red("Could not locate %s on %s", ic.Package, ic.Node)
		os.Exit(1)
	}

	if ic.CheckInstalled() == true {
		color.Yellow("%s is already installed. Moving along...", ic.Package)
		return
	}

	if pkg.Group != "" {
		url = fmt.Sprintf("http://%s:%d/crx/packmgr/service/.json/etc/packages/%s/%s?cmd=install",
			ic.Node,
			ic.Port,
			pkg.Group,
			ic.Package)
	} else {
		url = fmt.Sprintf("http://%s:%d/crx/packmgr/service/.json/etc/packages/%s?cmd=install",
			ic.Node,
			ic.Port,
			ic.Package)
	}

	easy := curl.EasyInit()
	defer easy.Cleanup()

	// Set options for curl
	easy.Setopt(curl.OPT_USERNAME, ic.Username)
	easy.Setopt(curl.OPT_PASSWORD, ic.Password)
	easy.Setopt(curl.OPT_URL, url)

	// Create a new form
	form := curl.NewForm()
	form.AddFile("package", fileName)

	// Print upload progress
	easy.Setopt(curl.OPT_NOPROGRESS, true)

	// Make it a post
	easy.Setopt(curl.OPT_CUSTOMREQUEST, "POST")

	if ic.AcClear == true {
		easy.Setopt(curl.OPT_POSTFIELDS, "acHandling=clear")
	}

	if ic.AcIgnore == true {
		easy.Setopt(curl.OPT_POSTFIELDS, "acHandling=ignore")
	}

	if ic.AcOverwrite == true {
		easy.Setopt(curl.OPT_POSTFIELDS, "acHandling=overwrite")
	}

	if ic.Recursive == true {
		easy.Setopt(curl.OPT_POSTFIELDS, "recursive=true")
	}

	if ic.Autosave != 0 {
		easy.Setopt(curl.OPT_POSTFIELDS,
			fmt.Sprintf("autosave=%d", ic.Autosave))
	}

	// Set connection timeout
	easy.Setopt(curl.OPT_CONNECTTIMEOUT, 10)

	// Get to work
	fmt.Printf("Installing %s. This might take a while\n", ic.Package)
	err := easy.Perform()
	Check(err)

}

func (ic *InstallCurl) VerifyInstall(count *int64) {

	if *count < ic.VerifyTimeout {
		// Verify Installation
		if ic.CheckInstalled() == false {
			*count += 1
			ic.VerifyInstall(count)
		} else {
			fmt.Printf("Installation of %s to %s succeeded\n",
				ic.Package, ic.Node)
		}
	} else {
		color.Red("The package %s did not successfully install\n", ic.Package)
		os.Exit(1)
	}
}
