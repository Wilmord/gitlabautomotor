package main

import (
	"fmt"
)

func main() {
	fmt.Println("---My gitlab automotor---")
<<<<<<< HEAD
	gitlabProj := newProject(401, "https://bro-gitlab.w2k.feico.com/")
	label, err := gitlabProj.findLabel("ccb-pending")
=======
	gitlabProj := newProject(401, "https://gitlab.com/")
	label, err := gitlabProj.findLabel("label-name")
>>>>>>> b5bc889 (first commit)
	if err != nil {
		fmt.Println(err.Error())
	}

	gitlabProj.removeLabelToMergeRequests(label)
}
