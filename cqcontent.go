package main

import (
	"cqcontent/argparse"
	"cqcontent/cq"
	"fmt"
	"io/ioutil"
	"os"
)

// Generic error handling function
func check(err error) {
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
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

	// Take methods of Curler interface
	lc := cq.Curler(listCurl)

	// Get XML of cq package content for the given node
	output := lc.Xml()

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

	// Take methods of Curler interface
	lc := cq.Curler(listCurl)

	// Get XML of cq package content for the given node
	lc.List()

	return fp
}

func main() {
	// Parse arguments to program
	arguments := argparse.ArgParse()

	// If cmmand is list, then execute list command
	if os.Args[1] == "xml" {
		_, fp := xml(arguments["NODE"].(string),
			arguments["USERNAME"].(string),
			arguments["PASSWORD"].(string))
		// Remove tempfile
		os.Remove(fp.Name())
	} else if os.Args[1] == "list" {
		fp := list(arguments["NODE"].(string),
			arguments["USERNAME"].(string),
			arguments["PASSWORD"].(string))
		// Remove tempfile
		os.Remove(fp.Name())
	}
}
