package cq

func (lc ListCurl) Download() {
	var decoder *Crx
	var foundPackage *Package
	var url string

	decoder = lc.Decoder()
	for _, p := range decoder.Response.Data.Packages.Packages {
		if lc.Package == p.DownloadName {
			foundPackage = p
			break
		}
	}

}
