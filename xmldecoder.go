package main

import (
	"encoding/xml"
)

//
// We need to map all of the values from the xml to structs, which can then
// be parsed and decoded so we can work with the values
//

type Package struct {
	Group          string `xml:"group"`
	Name           string `xml:"name"`
	Version        string `xml:"version"`
	DownloadName   string `xml:"downloadName"`
	Size           int32  `xml:"size"`
	Created        string `xml:"created"`
	CreatedBy      string `xml:"createdBy"`
	LastModified   string `xml:"listModified"`
	LastModifiedBy string `xml:"lastModifiedBy"`
	LastUnpacked   string `xml:"lastUnpacked"`
	LastUnpackedBy string `xml:"lastUnpackedBy"`
}

type Packages struct {
	XMLName  xml.Name  `xml:"packages"`
	Packages []Package `xml:"package"`
}

type Data struct {
	XMLName  xml.Name `xml:"data"`
	Packages Packages `xml:"packages"`
}

type Response struct {
	XMLName xml.Name `xml:"response"`
	Data    Data     `xml:"data"`
}

type Request struct {
	XMLName xml.Name `xml:"request"`
	Param   string   `xml:"param"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type Crx struct {
	XMLName   xml.Name `xml:"crx"`
	Version   string   `xml:"version,attr"`
	User      string   `xml:"user,attr"`
	Workspace string   `xml:"workspace,attr"`
	Request   Request  `xml:"request"`
	Response  Response `xml:"response"`
}

// Implementation of Decoder()
func (lc ListCurl) Decoder() *Crx {
	// Get xml
	output := lc.Xml()

	// Create a Crx object (the root of the xml)
	var obj Crx

	// pass the byte string output from the XML call and the Crx object
	xml.Unmarshal(output, &obj)

	// return reference to obj
	return &obj
}
