package main

import (
	"os"
	"strconv"
)

func main() {
	// Parse arguments to program
	arguments := ArgParse()

	// Call the wrapper function based on the sub-command passed as first arg
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
		fp := uploadWrapper(arguments["NODE"].(string),
			arguments["USERNAME"].(string),
			arguments["PASSWORD"].(string),
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
		fp := installWrapper(arguments["NODE"].(string),
			arguments["USERNAME"].(string),
			arguments["PASSWORD"].(string),
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
