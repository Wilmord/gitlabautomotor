package main

import (
	"fmt"
)

func main() {
	fmt.Println("---My gitlab automotor---")
	gitlabProj := newProject(projectID, "https://gitlab.com/")
	label, err := gitlabProj.findLabel("label-name")
	if err != nil {
		fmt.Println(err.Error())
	}
	gitlabProj.removeLabelToMergeRequests(label)
	//gitlabProj.addLabelToMergeRequests(label)
}
