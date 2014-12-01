package main

import (
	"os"
	"strconv"
)

func main() {
	// Parse arguments to program
	arguments := ArgParse()

	// Call the wrapper function based on the sub-command passed as first arg.
	// See wrapper.go for implementation of each wrapper function. See
	// argparse.go for argument parsing implementation.
	switch os.Args[1] {
	case "file":
		JsonWrapper(os.Args[2])
	case "xml":
		XmlWrapper(
			arguments["NODE"].(string),
			arguments["USERNAME"].(string),
			arguments["PASSWORD"].(string),
			StrToInt(arguments["--port"].(string)))

	case "list":
		fp := ListWrapper(
			arguments["NODE"].(string),
			arguments["USERNAME"].(string),
			arguments["PASSWORD"].(string),
			StrToInt(arguments["--port"].(string)))

		// Remove tempfile
		os.Remove(fp.Name())

	case "download":
		DownloadWrapper(
			arguments["NODE"].(string),
			arguments["USERNAME"].(string),
			arguments["PASSWORD"].(string),
			StrToInt(arguments["--port"].(string)),
			arguments["PACKAGE"].(string))

	case "upload":
		fp := UploadWrapper(
			arguments["NODE"].(string),
			arguments["USERNAME"].(string),
			arguments["PASSWORD"].(string),
			StrToInt(arguments["--port"].(string)),
			arguments["PACKAGE"].(string))
		// Remove tempfile
		os.Remove(fp.Name())

	case "install":
		// Set declare var for each install option
		var autosave int64
		var recursive bool
		var acIgnore bool
		var acOverwrite bool
		var acClear bool

		// Loop through the arguments and check for the options. If the option
		// is set on the command line, then set the corresponding var to that
		// set value
		for key, value := range arguments {
			switch key {
			case "--autosave":
				// Check if flag is present
				if value.(bool) == true {
					// Get string value passed
					s := arguments["NUMNODES"].(string)

					// Convert to int
					i, err := strconv.ParseInt(s, 10, 0)

					Check(err)
					autosave = i
				}
			case "--recursive":
				recursive = value.(bool)
			case "--acIgnore":
				acIgnore = value.(bool)
			case "--acOverwrite":
				acOverwrite = value.(bool)
			case "--acClear":
				acClear = value.(bool)
			}
		}

		// Pass all arguments and options to the installWrapper
		fp := InstallWrapper(
			arguments["NODE"].(string),
			arguments["USERNAME"].(string),
			arguments["PASSWORD"].(string),
			StrToInt(arguments["--port"].(string)),
			arguments["PACKAGE"].(string),
			autosave,
			recursive,
			acIgnore,
			acOverwrite,
			acClear)

		// Remove tempfile
		os.Remove(fp.Name())
	}
}
