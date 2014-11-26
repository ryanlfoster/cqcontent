package main

import (
	"fmt"
	docopt "github.com/docopt/docopt-go"
	"os"
)

func ArgParse() map[string]interface{} {
	usage := `

cqcontent

Work with cq content packages on the command line

Usage:
  cqcontent xml --node NODE --username USERNAME --password PASSWORD
  cqcontent list --node NODE --username USERNAME --password PASSWORD
  cqcontent download --node NODE --username USERNAME --password PASSWORD --package PACKAGE
  cqcontent upload --node NODE --username USERNAME --password PASSWORD --package PACKAGE
  cqcontent [ --help | -h | --version ]
`
	// Takes usage statement, argument list, help bool, version, and
	// optionsFirst bool. Returns a map.
	arguments, err := docopt.Parse(usage, os.Args[1:], true, "0.1", false)
	if err != nil {
		fmt.Println("Something went wrong!")
	}

	return arguments
}
