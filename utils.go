package main

import (
	"regexp"
	"os"
	"path/filepath"
	"github.com/fatih/color"
)

// Generic error handling function
func Check(err error) {
	if err != nil {
		color.Red("ERROR: %v\n", err)
		os.Exit(1)
	}
}

// callback function for OPT_WRITEFUNCTION. See libcurl docs.
func WriteData(ptr []byte, userdata interface{}) bool {

	file := userdata.(*os.File)
	if _, err := file.Write(ptr); err != nil {
		return false
	} else {
		return true
	}
}

func IsRelative (str string) bool {
	regex, err := regexp.Compile("/")
	if err != nil {
		color.Red("%v\n", err)
		os.Exit(1)
	}
	match := regex.FindString(str)

	if match == "" {
		return true
	} else {
		return false
	}
}

func FullPath (str string) string {
	if IsRelative(str) == true {
		absPath, err := filepath.Abs(str)
		Check(err)
		return absPath
	} else {
		absPath := str
		return absPath
	}
}

func RelPath (str string) string {
	if IsRelative(str) == false {
		relPath := filepath.Base(str)
		return relPath
	} else {
		relPath := str
		return relPath
	}
}

func FileExists (str string) (bool, string) {
	absPath := FullPath(str)
	if _, err := os.Stat(absPath); err == nil {
		return true, absPath
	} else {
		return false, ""
	}
}
