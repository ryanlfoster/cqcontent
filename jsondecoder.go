package main

import (
	"github.com/fatih/color"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
)

// Use this struct for CheckValueLoop so you can check for nil against pointers
// to strings for unset values
type JobValidate struct {
	Mode string `json:"mode"`
	TargetNode *string `json:"target_node"`
	TargetUsername *string `json:"target_username"`
	TargetPassword *string `json:"target_password"`
}

// Use this struct for the JobLoop
type Job struct {
	Mode string `json:"mode"`
	TargetNode string `json:"target_node"`
	TargetUsername string `json:"target_username"`
	TargetPassword string `json:"target_password"`
}

func CheckValueLoop (path string) {
	jbytes, err := ioutil.ReadFile(path)
	Check(err)

	var jobs []JobValidate
	json.Unmarshal(jbytes, &jobs)

	for _, job := range jobs {
		switch job.Mode {
		case "xml":
			if job.TargetNode == nil ||
				job.TargetUsername == nil ||
				job.TargetPassword == nil {

				color.Red(`
The following settings are requred for the xml job:

target_node
target_username
target_password

				`)
				os.Exit(1)
			}
		case "list":
			if job.TargetNode == nil ||
				job.TargetUsername == nil ||
				job.TargetPassword == nil {

				color.Red(`
The following settings are requred for the xml job:

target_node
target_username
target_password

				`)
				os.Exit(1)
			}
		}
	}
}

func JobLoop (path string) {
	jbytes, err := ioutil.ReadFile(path)
	Check(err)

	var jobs []Job
	json.Unmarshal(jbytes, &jobs)

	for _, job := range jobs {
		switch job.Mode {
		case "xml":
			bytes, fp := xmlWrapper(job.TargetNode,
				job.TargetUsername,
				job.TargetPassword)
			fmt.Printf("%s\n", bytes)
			os.Remove(fp.Name())
		case "list":
			fp := listWrapper(job.TargetNode,
				job.TargetUsername,
				job.TargetPassword)
			os.Remove(fp.Name())
		}
	}
}
