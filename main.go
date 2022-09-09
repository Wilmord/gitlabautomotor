package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func stringTrimmer(str string) string {
	if runtime.GOOS == "windows" {
		return strings.Replace(str, "\r\n", "", -1)
	} else {
		return strings.Replace(str, "\n", "", -1)
	}
}

func main() {

	//fmt.Println(os.Args[1])
	//relNumber := os.Args[1]

	fmt.Println("---My gitlab automotor---")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Choose project you want:")
		fmt.Println("1. Imaging")
		fmt.Println("2. Velox")
		fmt.Println("Any other button to exit")

		inputId, _ := reader.ReadString('\n')
		inputId = stringTrimmer(inputId)
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

		gitlabProj := newProject(projectID, "https://bro-gitlab.w2k.feico.com/")

		fmt.Println("Enter the number of operation you want:")
		fmt.Println("1. Label Opertion")
		fmt.Println("2. Dependency information")

		inputOperation, _ := reader.ReadString('\n')
		inputOperation = stringTrimmer(inputOperation)
		if inputOperation == "1" {
			label, err := gitlabProj.findLabel("label-name")
			if err != nil {
				fmt.Println(err.Error())
			}
			//gitlabProj.removeLabelToMergeRequests(label)
			gitlabProj.addLabelToMergeRequests(label)
		} else if inputOperation == "2" {

			fmt.Println("Please enter release tag, for example: REL-3.6.0")

			inputVersion, _ := reader.ReadString('\n')
			inputVersion = stringTrimmer(inputVersion)
			gitlabProj.getDependencyInformation(inputVersion)
		} else {
			fmt.Println("Invalid Operation!")
			break
		}
	}

}
