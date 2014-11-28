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
  cqcontent file     CONFIG_FILE
  cqcontent xml      (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
  cqcontent list     (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
  cqcontent download (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
                     (--package|-k) PACKAGE
  cqcontent upload   (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
                     (--package|-k) PACKAGE
  cqcontent install  (--node|-n) NODE
                     (--username|-u) USERNAME
                     (--password|-p) PASSWORD
                     (--package|-k) PACKAGE
                     [--autosave NUMNODES]
                     [--recursive]
                     [--acIgnore]
                     [--acOverwrite]
                     [--acClear]
  cqcontent          [--help|-h|--version]
`
	// Takes usage statement, argument list, help bool, version, and
	// optionsFirst bool. Returns a map.
	arguments, err := docopt.Parse(usage, os.Args[1:], true, "0.1", false)
	if err != nil {
		fmt.Println("Something went wrong!")
	}

	return arguments
}
