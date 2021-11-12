package main

import (
	"fmt"
)

func main() {
	fmt.Println("---My gitlab automotor---")
	gitlabProj := newProject(401, "https://bro-gitlab.w2k.feico.com/")
	label, err := gitlabProj.findLabel("ccb-pending")
	if err != nil {
		fmt.Println(err.Error())
	}

	gitlabProj.removeLabelToMergeRequests(label)
}
