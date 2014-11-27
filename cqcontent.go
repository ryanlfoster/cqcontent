package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

// Generic error handling function
func Check(err error) {
	if err != nil {
		color.Red("ERROR: %v\n", err)
		os.Exit(1)
	}
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
	fp, err := os.OpenFile(pkg, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	Check(err)

	// Handle possible errors
	Check(err)

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
	pkg string) {

	// Initialize tempfile and get pointer
	fp, err := ioutil.TempFile("", "cq")

	// Handle possible errors
	Check(err)

	// Initialize struct
	uploadCurl := UploadCurl{
		ListCurl: ListCurl{
			CurlFp: CurlFp{
				Curl: Curl{Username: username, Password: password},
			Fp: fp},
		Node: node},
	Package: pkg, Uploaded: false}

	// Get XML of cq package content for the given node
	uploadCurl.Upload()

}

func main() {
	// Parse arguments to program
	arguments := ArgParse()

	// If cmmand is list, then execute list command
	switch os.Args[1] {
	case "xml":
		_, fp := xmlWrapper(arguments["NODE"].(string),
			arguments["USERNAME"].(string),
			arguments["PASSWORD"].(string))
		// Remove tempfile
		os.Remove(fp.Name())
	case "list":
		fp := listWrapper(arguments["NODE"].(string),
			arguments["USERNAME"].(string),
			arguments["PASSWORD"].(string))
		// Remove tempfile
		os.Remove(fp.Name())
	case "download":
		downloadWrapper(arguments["NODE"].(string),
			arguments["USERNAME"].(string),
			arguments["PASSWORD"].(string),
			arguments["PACKAGE"].(string))
	case "upload":
		uploadWrapper(arguments["NODE"].(string),
			arguments["USERNAME"].(string),
			arguments["PASSWORD"].(string),
			arguments["PACKAGE"].(string))
	}
}
