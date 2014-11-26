package main

import (
	"cqcontent/argparse"
	"cqcontent/cq"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/fatih/color"
)

// Generic error handling function
func check(err error) {
	if err != nil {
		color.Red("ERROR: %v\n", err)
		os.Exit(1)
	}
}

// Print the XML data for the given host
func xml(node string, username string, password string) ([]byte, *os.File) {
	// Initialize tempfile and get pointer
	fp, err := ioutil.TempFile("", "cq")

	// Handle possible errors
	check(err)

	// Initialize struct
	listCurl := cq.ListCurl{
		Curl: cq.Curl{Fp: fp, Username: username, Password: password},
		Node:  node}

	// Get XML of cq package content for the given node
	output := listCurl.Xml()

	// Print the output
	fmt.Printf("%s", output)

	// Close the tempfile
	fp.Close()

	return output, fp
}

// Print XML data parsed for now
func list(node string, username string, password string) *os.File {
	// Initialize tempfile and get pointer
	fp, err := ioutil.TempFile("", "cq")

	// Handle possible errors
	check(err)

	// Initialize struct
	listCurl := cq.ListCurl{
		Curl: cq.Curl{Fp: fp, Username: username, Password: password},
		Node:  node}

	// Get XML of cq package content for the given node
	listCurl.List()

	return fp
}

// Download the package
func download(node string,
	          username string,
	          password string,
	          pkg string) *os.File {

	// Initialize file and get pointer
	fp, err := os.OpenFile(pkg, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0777)
	check(err)

	// Handle possible errors
	check(err)

	// Initialize struct
	downloadCurl := cq.DownloadCurl{
		ListCurl: cq.ListCurl{
		Curl: cq.Curl{Fp: fp, Username: username, Password: password}, Node:  node},
		Package: pkg}

	// Get XML of cq package content for the given node
	downloadCurl.Download()

	return fp

}

func main() {
	// Parse arguments to program
	arguments := argparse.ArgParse()

	// If cmmand is list, then execute list command
	switch os.Args[1] {
	case "xml":
		_, fp := xml(arguments["NODE"].(string),
			arguments["USERNAME"].(string),
			arguments["PASSWORD"].(string))
		// Remove tempfile
		os.Remove(fp.Name())
	case "list":
		fp := list(arguments["NODE"].(string),
			arguments["USERNAME"].(string),
			arguments["PASSWORD"].(string))
		// Remove tempfile
		os.Remove(fp.Name())
	case "download":
		download(arguments["NODE"].(string),
			arguments["USERNAME"].(string),
			arguments["PASSWORD"].(string),
			arguments["PACKAGE"].(string))
	}
}
