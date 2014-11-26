package main

import (
	"regexp"
	"os"
	"path/filepath"
	"github.com/fatih/color"
)

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

func FileExists (str string) (bool, string) {
	absPath := FullPath(str)
	if _, err := os.Stat(absPath); err == nil {
		return true, absPath
	} else {
		return false, ""
	}
}
