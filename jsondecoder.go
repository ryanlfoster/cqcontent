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
	Node           *string `json:"node"`
	Username       *string `json:"username"`
	Password       *string `json:"password"`
	Package        *string `json:"package"`
	Port           *int64  `json:"port"`
	Autosave	   *int64  `json:"autosave"`
	Recursive      *bool   `json:"recursive"`
	AcIgnore       *bool   `json:"acIgnore"`
	AcOverwrite    *bool   `json:"acOverwrite"`
	AcClear        *bool   `json:"acClear"`
}

// Use this struct for the JobLoop
type Job struct {
	Mode           string `json:"mode"`
	Node           string `json:"node"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Package        string `json:"package"`
	Port           int64  `json:"port"`
	Autosave	   int64  `json:"autosave"`
	Recursive      bool   `json:"recursive"`
	AcIgnore       bool   `json:"acIgnore"`
	AcOverwrite    bool   `json:"acOverwrite"`
	AcClear        bool   `json:"acClear"`
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

// Check to see if the port is set since it is optional
func CheckPort(job *JobValidate) bool {
	if job.Port == nil {
		return false
	} else {
		return true
	}
}

// Check to see if the port is set since it is optional
func CheckAutosave(job *JobValidate) bool {
	if job.Autosave == nil {
		return false
	} else {
		return true
	}
}

// Check to see if the recursive is set since it is optional
func CheckRecursive(job *JobValidate) bool {
	if job.Recursive == nil {
		return false
	} else {
		return true
	}
}

// Check to see if the acIgnore is set since it is optional
func CheckAcIgnore(job *JobValidate) bool {
	if job.AcIgnore == nil {
		return false
	} else {
		return true
	}
}

// Check to see if the acOverwrite is set since it is optional
func CheckAcOverwrite(job *JobValidate) bool {
	if job.AcOverwrite == nil {
		return false
	} else {
		return true
	}
}

// Check to see if the acOverwrite is set since it is optional
func CheckAcClear(job *JobValidate) bool {
	if job.AcClear == nil {
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
			if job.Node == nil ||
				job.Username == nil ||
				job.Password == nil {

				color.Red(`
The following settings are requred for the xml job:

node
username
password
				`)
				os.Exit(1)
			}

		case "list":
			if job.Node == nil ||
				job.Username == nil ||
				job.Password == nil {

				color.Red(`
The following settings are requred for the list job:

node
username
password
				`)
				os.Exit(1)
			}
		case "download":
			if job.Node == nil ||
				job.Username == nil ||
				job.Password == nil ||
				job.Package == nil {

				color.Red(`
The following settings are requred for the download job:

node
username
password
package
					`)
				os.Exit(1)

			}
		case "upload":
			if job.Node == nil ||
				job.Username == nil ||
				job.Password == nil ||
				job.Package == nil {

				color.Red(`
The following settings are requred for the upload job:

node
username
password
package
					`)
				os.Exit(1)

			}
		case "install":
			if job.Node == nil ||
				job.Username == nil ||
				job.Password == nil ||
				job.Package == nil {

				color.Red(`
The following settings are requred for the install job:

node
username
password
package
					`)
				os.Exit(1)

			}
		}
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
			if CheckPort(&validateJobs[count]) == true {
				port = job.Port
			} else {
				port = 8080
			}
			str, _ := XmlWrapper(
				job.Node,
				job.Username,
				job.Password,
				port)
			fmt.Printf("%s\n", str)
		case "list":
			var port int64
			if CheckPort(&validateJobs[count]) == true {
				port = job.Port
			} else {
				port = 8080
			}

			ListWrapper(
				job.Node,
				job.Username,
				job.Password,
				port)

		case "download":
			// Handle Port
			var port int64
			if CheckPort(&validateJobs[count]) == true {
				port = job.Port
			} else {
				port = 8080
			}
			DownloadWrapper(
				job.Node,
				job.Username,
				job.Password,
				port,
				job.Package)
		case "upload":
			// Handle Port
			var port int64
			if CheckPort(&validateJobs[count]) == true {
				port = job.Port
			} else {
				port = 8080
			}
			UploadWrapper(
				job.Node,
				job.Username,
				job.Password,
				port,
				job.Package)
		case "install":
			// Handle Port
			var port int64
			if CheckPort(&validateJobs[count]) == true {
				port = job.Port
			} else {
				port = 8080
			}

			// Handle autosave option
			var autosave int64
			if CheckAutosave(&validateJobs[count]) == true {
				autosave = job.Autosave
			} else {
				autosave = 0
			}

			// Handle recursive option
			var recursive bool
			if CheckRecursive(&validateJobs[count]) == true {
				recursive = job.Recursive
			} else {
				recursive = false
			}

			// Handle acIgnore option
			var acIgnore bool
			if CheckAcIgnore(&validateJobs[count]) == true {
				acIgnore = job.AcIgnore
			} else {
				acIgnore = false
			}

			var acOverwrite bool
			if CheckAcOverwrite(&validateJobs[count]) == true {
				acOverwrite = job.AcOverwrite
			} else {
				acOverwrite = false
			}

			var acClear bool
			if CheckAcClear(&validateJobs[count]) == true {
				acClear = job.AcClear
			} else {
				acClear = false
			}

			InstallWrapper(
				job.Node,
				job.Username,
				job.Password,
				port,
				job.Package,
				autosave,
				recursive,
				acIgnore,
				acOverwrite,
				acClear)
		}
		count += 1
	}
}
