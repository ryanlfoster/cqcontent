package argparse

import (
	"fmt"
	docopt "github.com/docopt/docopt-go"
	"os"
)

func ArgParse() map[string]interface{} {
	usage := `

cqcontent

Download, upload, and install content packages from one host to a cluster of
hosts in a single network environment.

Usage:
  cqcontent full --host HOST --nodes HOSTNAME --package PACKAGE --env ENV
  cqcontent xml --node NODE --username USERNAME --password PASSWORD
  cqcontent list --node NODE --username USERNAME --password PASSWORD
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
