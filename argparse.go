package main

import (
	"fmt"
	docopt "github.com/docopt/docopt-go"
	"os"
)

// A little explanation of docopt from what I can gather...
//
// Regular arguments are referenced by the value name passed to the flag in
// the docopt string. You can test if a flag was passed by getting the value
// of the actual flag. The values must be type casted as strings.
//
// For example:
//
// arguments["PASSWORD"].(string) will get you the string, because the
// value for --password is described as PASSWORD
//
// To check if a flag is present, check against the value of the actual
// flag, which is a boolean.
//
// For example:
//
// arguments["--password"].(bool) will give you the boolean value for the
// presence of the flag on the command line when executed.
//
// Options are handled slightly differently. Check for the value of the
// option by actually checking the flag.
//
// For example:
//
// arguments["--port"].(string) will give you the value of the port.
//
// To print out the arguments to figure this out, just print out the
// arguments variable with %v.
func ArgParse() map[string]interface{} {
	usage := `

cqcontent

Work with cq content packages on the command line

Usage:
  cqcontent file     [options] CONFIG_FILE
  cqcontent xml      [options]
  					 (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
  cqcontent list     [options]
  					 (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
  cqcontent download [options]
  					 (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
                     (--package|-k) PACKAGE
  cqcontent upload   [options]
  					 (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
                     (--package|-k) PACKAGE
  cqcontent install  [options]
  					 (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
                     (--package|-k) PACKAGE
                     [--autosave NUMNODES]
                     [--recursive]
                     [--acIgnore]
                     [--acOverwrite]
                     [--acClear]
  cqcontent          [--help|-h|--version]

Options:
  --port PORT        Specify the port. [default: 8080]
`
	// Takes usage statement, argument list, help bool, version, and
	// optionsFirst bool. Returns a map.
	arguments, err := docopt.Parse(usage, os.Args[1:], true, "0.1", false)
	if err != nil {
		fmt.Println("Something went wrong!")
	}

	return arguments
}
