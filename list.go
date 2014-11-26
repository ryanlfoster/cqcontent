package main

import "fmt"

// Implementation of the List method. See Decoder method for more information
func (lc ListCurl) List() {
	var decoder *Crx
	decoder = lc.Decoder()

	// Loop through every package. If the value of lastUnpackedBy is an empty
	// then the package was only uploaded, and not installed. With that said
	// only print installed packages.
	for _, p := range decoder.Response.Data.Packages.Packages {
		if p.LastUnpackedBy != "" {
			fmt.Printf("%s\n", p.DownloadName)
		}
	}
}
