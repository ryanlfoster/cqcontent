package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Do work based on the json configuraiton file
func JsonWrapper(path string) {
	CheckValueLoop(path)
	JobLoop(path)
}

// Print the XML data for the given host
func XmlWrapper(
	node string,
	username string,
	password string,
	port int64) ([]byte, *os.File) {

	// Initialize struct
	listCurl := ListCurl{
		Curl: Curl{Username: username, Password: password, Port: port},
		Node: node}

	// Get XML of cq package content for the given node
	output, fp := listCurl.Xml()

	// Print the output
	fmt.Printf("%s", output)

	return output, fp
}

// Print list of installed package for a given node
func ListWrapper(
	node string,
	username string,
	password string,
	port int64) *os.File {

	// Initialize tempfile and get pointer
	fp, err := ioutil.TempFile("", "cq")

	// Handle possible errors
	Check(err)

	// Initialize struct
	listCurl := ListCurl{
		Curl: Curl{Username: username, Password: password, Port: port},
		Node: node}

	// Get XML of cq package content for the given node
	listCurl.List()

	return fp
}

// Download the package
func DownloadWrapper(
	node string,
	username string,
	password string,
	port int64,
	pkg string) {

	// Initialize struct
	downloadCurl := DownloadCurl{
		ListCurl: ListCurl{
			Curl: Curl{Username: username, Password: password, Port: port},
			Node: node},
		Package: pkg}

	// Get XML of cq package content for the given node
	downloadCurl.Download()

}

func UploadWrapper(
	node string,
	username string,
	password string,
	port int64,
	pkg string) *os.File {

	// Initialize tempfile and get pointer
	fp, err := ioutil.TempFile("", "cq")

	// Handle possible errors
	Check(err)

	// Initialize struct
	uploadCurl := UploadCurl{
		DownloadCurl: DownloadCurl{
			ListCurl: ListCurl{
				Curl: Curl{Username: username, Password: password, Port: port},
				Node: node},
			Package: pkg},
		Uploaded: false}

	// Get XML of cq package content for the given node
	uploadCurl.Upload()

	return fp

}

func InstallWrapper(
	node string,
	username string,
	password string,
	port int64,
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
					Curl: Curl{Username: username, Password: password, Port: port},
					Node: node},
				Package: pkg},
			Uploaded: false},
		Installed: false, Autosave: autosave, Recursive: recursive, AcIgnore: acIgnore, AcOverwrite: acOverwrite, AcClear: acClear}

	// Get XML of cq package content for the given node
	installCurl.Install()

	return fp

}
