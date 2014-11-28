package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Do work based on the json configuraiton file
func jsonWrapper(path string) {
	CheckValueLoop(path)
	JobLoop(path)
}

// Print the XML data for the given host
func xmlWrapper(
	node string,
	username string,
	password string) ([]byte, *os.File) {

	// Initialize tempfile and get pointer
	fp, err := ioutil.TempFile("", "cq")

	// Handle possible errors
	Check(err)

	// Initialize struct
	listCurl := ListCurl{
		CurlFp: CurlFp{
			Curl: Curl{Username: username, Password: password},
		Fp: fp},
	Node: node}

	// Get XML of cq package content for the given node
	output := listCurl.Xml()

	// Print the output
	fmt.Printf("%s", output)

	// Close the tempfile
	fp.Close()

	return output, fp
}

// Print XML data parsed for now
func listWrapper(node string, username string, password string) *os.File {
	// Initialize tempfile and get pointer
	fp, err := ioutil.TempFile("", "cq")

	// Handle possible errors
	Check(err)

	// Initialize struct
	listCurl := ListCurl{
		CurlFp: CurlFp{
			Curl: Curl{Username: username, Password: password},
		Fp: fp},
	Node: node}

	// Get XML of cq package content for the given node
	listCurl.List()

	return fp
}

// Download the package
func downloadWrapper(node string,
	username string,
	password string,
	pkg string) *os.File {

	// Initialize file and get pointer
	fp, err := os.OpenFile(pkg, os.O_WRONLY|os.O_CREATE, 0777)
	Check(err)
	defer func() {
        err := fp.Close()
		Check(err)
    }()

	// Initialize struct
	downloadCurl := DownloadCurl{
		ListCurl: ListCurl{
			CurlFp: CurlFp{
				Curl: Curl{Username: username, Password: password},
			Fp: fp},
		Node: node},
	Package: pkg}


	// Get XML of cq package content for the given node
	downloadCurl.Download()

	return fp

}

func uploadWrapper(node string,
	username string,
	password string,
	pkg string) *os.File {

	// Initialize tempfile and get pointer
	fp, err := ioutil.TempFile("", "cq")

	// Handle possible errors
	Check(err)

	// Initialize struct
	uploadCurl := UploadCurl{
		DownloadCurl: DownloadCurl{
			ListCurl: ListCurl{
				CurlFp: CurlFp{
					Curl: Curl{Username: username, Password: password},
				Fp: fp},
			Node: node},
		Package: pkg},
	Uploaded: false}

	// Get XML of cq package content for the given node
	uploadCurl.Upload()

	return fp

}

func installWrapper(node string,
	username string,
	password string,
	pkg string,
	autosave int64,
	recursive bool,
	acIgnore bool,
	acOverwrite bool,
	acClear bool) *os.File {

	// Initialize tempfile and get pointer
	fp, err := ioutil.TempFile("", "cq")

	// Handle possible errors
	Check(err)

	// Initialize struct
	installCurl := InstallCurl{
		UploadCurl: UploadCurl{
			DownloadCurl: DownloadCurl{
				ListCurl: ListCurl{
					CurlFp: CurlFp{
						Curl: Curl{Username: username, Password: password},
					Fp: fp},
				Node: node},
			Package: pkg},
		Uploaded: false},
	Installed: false, Autosave: autosave, Recursive: recursive, AcIgnore: acIgnore, AcOverwrite: acOverwrite, AcClear: acClear}


	// Get XML of cq package content for the given node
	installCurl.Install()

	return fp

}
