package main

import (
	"fmt"
	"os"
)

// Do work based on the json configuration file
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
	listCurl := &ListCurl{
		Curl: &Curl{Username: username, Password: password, Port: port},
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
	port int64) {

	// Initialize struct
	listCurl := &ListCurl{
		Curl: &Curl{Username: username, Password: password, Port: port},
		Node: node}

	// Get XML of cq package content for the given node
	listCurl.List()

}

// Download the package
func DownloadWrapper(
	node string,
	username string,
	password string,
	port int64,
	pkg string) {

	// Initialize struct
	downloadCurl := &DownloadCurl{
		ListCurl: &ListCurl{
			Curl: &Curl{Username: username, Password: password, Port: port},
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
	verifyTimeout int64,
	pkg string) {

	// Initialize struct
	uploadCurl := &UploadCurl{
		DownloadCurl: &DownloadCurl{
			ListCurl: &ListCurl{
				Curl: &Curl{Username: username, Password: password, Port: port},
				Node: node},
			Package: pkg},
		Uploaded: false, VerifyTimeout: verifyTimeout}

	// Get XML of cq package content for the given node
	uploadCurl.Upload()
	var count int64 = 0
	var countPtr *int64 = &count
	uploadCurl.VerifyUpload(countPtr)

}

func InstallWrapper(
	node string,
	username string,
	password string,
	port int64,
	verifyTimeout int64,
	pkg string,
	autosave int64,
	recursive bool,
	acIgnore bool,
	acOverwrite bool,
	acClear bool) {

	// Initialize struct
	installCurl := &InstallCurl{
		UploadCurl: &UploadCurl{
			DownloadCurl: &DownloadCurl{
				ListCurl: &ListCurl{
					Curl: &Curl{Username: username, Password: password, Port: port},
					Node: node},
				Package: pkg},
			Uploaded: false, VerifyTimeout: verifyTimeout},
		Installed: false, Autosave: autosave, Recursive: recursive, AcIgnore: acIgnore, AcOverwrite: acOverwrite, AcClear: acClear}

	// Get XML of cq package content for the given node
	installCurl.Install()
	var count int64 = 0
	var countPtr *int64 = &count
	installCurl.VerifyInstall(countPtr)

}

func DeleteWrapper(
	node string,
	username string,
	password string,
	port int64,
	verifyTimeout int64,
	pkg string) {

	// Initalize struct
	deleteCurl := &DeleteCurl{
		UploadCurl: &UploadCurl{
			DownloadCurl: &DownloadCurl{
				ListCurl: &ListCurl{
					Curl: &Curl{Username: username, Password: password, Port: port},
					Node: node},
				Package: pkg},
			Uploaded: false, VerifyTimeout: verifyTimeout}}

	deleteCurl.Delete()
	var count int64 = 0
	var countPtr *int64 = &count
	deleteCurl.VerifyDelete(countPtr)

}
