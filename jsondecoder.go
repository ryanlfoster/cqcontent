package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
)

type Job struct {
	Mode string `json:"mode"`
	TargetNode string `json:"target_node"`
	TargetUsername string `json:"target_username"`
	TargetPassword string `json:"target_password"`
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
			bytes, fp := xmlWrapper(job.TargetNode,
				job.TargetUsername,
				job.TargetPassword)
			fmt.Printf("%s\n", bytes)
			os.Remove(fp.Name())
		}
	}
}
