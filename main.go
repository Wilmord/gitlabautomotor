package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func main() {

	fmt.Println("---My gitlab automotor---")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Choose project you want:")
		fmt.Println("1. Project 1")
		fmt.Println("2. Project 2")
		fmt.Println("Any other button to exit")

		inputId, _ := reader.ReadString('\n')
		if runtime.GOOS == "windows" {
			inputId = strings.Replace(inputId, "\r\n", "", -1)
		} else {
			inputId = strings.Replace(inputId, "\n", "", -1)
		}

		var projectID int
		if inputId == "1" {
			projectID = 401
		} else if inputId == "2" {
			projectID = 155
		} else {
			fmt.Println("Invalid Project")
			fmt.Println("Exited")
			break
		}

		gitlabProj := newProject(projectID, "https://gitlab.com/")

		fmt.Println("Enter the number of operation you want:")
		fmt.Println("1. Label Opertion")
		fmt.Println("2. Dependency information")

		inputOperation, _ := reader.ReadString('\n')
		if runtime.GOOS == "windows" {
			inputOperation = strings.Replace(inputId, "\r\n", "", -1)
		} else {
			inputOperation = strings.Replace(inputId, "\n", "", -1)
		}

		if inputOperation == "1" {
			label, err := gitlabProj.findLabel("label-name")
			if err != nil {
				fmt.Println(err.Error())
			}
			//gitlabProj.removeLabelToMergeRequests(label)
			gitlabProj.addLabelToMergeRequests(label)
		} else if inputOperation == "2" {
			gitlabProj.getDependencyInformation()
		} else {
			fmt.Println("Invalid Operation!")
			break
		}
	}

}
