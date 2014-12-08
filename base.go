/*
cq package to work with cq content packages

http://curl.haxx.se/libcurl/c/libcurl-tutorial.html
https://github.com/andelf/go-curl/
the --libcurl option to the curl binary is your friend.
*/

package main

// Structs for each action

type Curl struct {
	Username string
	Password string
	Port     int64
}

type ListCurl struct {
	*Curl
	Node string
}

type DownloadCurl struct {
	*ListCurl
	Package string
}

type DeleteCurl struct {
	*UploadCurl
}

// Augmented struct to hold pkgpath and node value
type UploadCurl struct {
	*DownloadCurl
	Uploaded      bool
	VerifyTimeout int64
}

// Augmented struct
// Inherits ListCurl struct and adds support for some options:
//
// autosave:    Number of nodes after which to perform intermediate saves. Set
//              to 0 to omit this option.
// recursive:   Set to true to install subpackages as well. Set to false to
//              omit this option
// acIgnore:    Ignores the packaged access control and leaves the target
//              unchanged. Set to false to omit this option
// acOverwrite: Applies the access control provided with the package to the
//              target. This also removes the existing access controls. Set to
//              false to omit.
// acClear:     Clears all access control on the target system.
//
// http://docs.adobe.com/docs/en/crx/current/how_to/package_manager.html#Installing packages (CLI)
type InstallCurl struct {
	*UploadCurl
	Installed   bool
	Autosave    int64
	Recursive   bool
	AcIgnore    bool
	AcOverwrite bool
	AcClear     bool
}
