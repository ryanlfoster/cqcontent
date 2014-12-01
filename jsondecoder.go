package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

// Use this struct for CheckValueLoop so you can check for nil against pointers
// to strings for unset values
type JobValidate struct {
	Mode           string  `json:"mode"`
	TargetNode     *string `json:"target_node"`
	TargetUsername *string `json:"target_username"`
	TargetPassword *string `json:"target_password"`
	Package        *string `json:"package"`
	Port           *int64  `json:"port"`
}

// Use this struct for the JobLoop
type Job struct {
	Mode           string `json:"mode"`
	TargetNode     string `json:"target_node"`
	TargetUsername string `json:"target_username"`
	TargetPassword string `json:"target_password"`
	Package        string `json:"package"`
	Port           int64  `json:"port"`
}

// Check if json is syntactically valid
func isJSON(d []byte) bool {
	var jv []JobValidate
	err := json.Unmarshal(d, &jv)
	if err != nil {
		return false
	} else {
		return true
	}
}

// Loop through each job and make sure the right settings are there before
// doing anything
func CheckValueLoop(path string) {
	jbytes, err := ioutil.ReadFile(path)
	Check(err)

	if isJSON(jbytes) == false {
		color.Red(`
The json configuration file provided is not syntactically valid.
		`)
	}
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
		case "download":
			if job.TargetNode == nil ||
				job.TargetUsername == nil ||
				job.TargetPassword == nil ||
				job.Package == nil {

				color.Red(`
The following settings are requred for the xml job:

target_node
target_username
target_password
package
					`)
				os.Exit(1)

			}
		}
	}
}

// Check to see if the port is set since it is optional
func CheckPortSet(job *JobValidate) bool {
	if job.Port == nil {
		return false
	} else {
		return true
	}
}

// Loop through each job and do work
func JobLoop(path string) {
	// Get json as byte array and check for errors
	jbytes, err := ioutil.ReadFile(path)
	Check(err)

	// Unmarshal jobs for work and configuration checking
	var jobs []Job
	json.Unmarshal(jbytes, &jobs)

	var validateJobs []JobValidate
	json.Unmarshal(jbytes, &validateJobs)

	// Loop through each job
	count := 0
	for _, job := range jobs {
		switch job.Mode {
		case "xml":
			// If port was set, then use that value, otherwise, default to 8080
			var port int64
			if CheckPortSet(&validateJobs[count]) == true {
				port = job.Port
			} else {
				port = 8080
			}
			str, _ := XmlWrapper(
				job.TargetNode,
				job.TargetUsername,
				job.TargetPassword,
				port)
			fmt.Printf("%s\n", str)
		case "list":
			var port int64
			if CheckPortSet(&validateJobs[count]) == true {
				port = job.Port
			} else {
				port = 8080
			}

			ListWrapper(
				job.TargetNode,
				job.TargetUsername,
				job.TargetPassword,
				port)

		case "download":
			// Handle Port
			var port int64
			if CheckPortSet(&validateJobs[count]) == true {
				port = job.Port
			} else {
				port = 8080
			}
			DownloadWrapper(
				job.TargetNode,
				job.TargetUsername,
				job.TargetPassword,
				port,
				job.Package)
		}
		count += 1
	}
}
