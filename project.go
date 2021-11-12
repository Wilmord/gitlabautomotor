package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/xanzy/go-gitlab"
)

type project struct {
	projectID     int
	projectName   string
	gitlabProject *gitlab.Project
	gitlabClient  *gitlab.Client
}

//Constructor
func newProject(id int, url string) *project {

	gitlabToken := os.Getenv("GITLAB_TOKEN")
	if gitlabToken == "" {
		panic("Missing authentication information in environment.")
	}

	gitClient, err := gitlab.NewClient(gitlabToken, gitlab.WithBaseURL(url))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	//TODO create project and give id as an argument
	proj, _, projectErr := gitClient.Projects.GetProject(id, &gitlab.GetProjectOptions{})
	if projectErr != nil {
		log.Fatalf("Failed to get project: %v", projectErr)
	}

	return &project{projectID: id, projectName: proj.Name, gitlabProject: proj, gitlabClient: gitClient}
}

func (prj project) findLabel(labelStr string) (*gitlab.Label, error) {
	labels, _, labelsErr := prj.gitlabClient.Labels.ListLabels(prj.projectID, &gitlab.ListLabelsOptions{})
	if labelsErr != nil {
		log.Fatalf("Failed to get project labels: %v", labelsErr)
	}

	for _, label := range labels {
		if label.Name == labelStr {
			return label, nil
		}
	}
	return nil, errors.New("label does not exist")
}

func (prj project) getAllMergeRequests() ([]*gitlab.MergeRequest, error) {

	mergeRequests, _, mergeRequestErr := prj.gitlabClient.MergeRequests.ListProjectMergeRequests(prj.projectID,
		&gitlab.ListProjectMergeRequestsOptions{
			State:   gitlab.String("opened"),
			OrderBy: gitlab.String("created_at"),
			Sort:    gitlab.String("asc"),
			ListOptions: gitlab.ListOptions{
				PerPage: 100,
			},
		})

	if mergeRequestErr != nil {
		return nil, errors.New("failed to get merge requests")
	}
	return mergeRequests, nil
}

func (prj project) addLabelToMergeRequests(label *gitlab.Label) {

	mergeRequests, err := prj.getAllMergeRequests()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, mr := range mergeRequests {

			fmt.Println(mr.Title)
			fmt.Println(mr.IID)
			mr.Labels = append(mr.Labels, label.Name)
			prj.gitlabClient.MergeRequests.UpdateMergeRequest(prj.projectID, mr.IID, &gitlab.UpdateMergeRequestOptions{AddLabels: mr.Labels})
			fmt.Println(mr.Labels)
			fmt.Println("---------------------------------------------")

		}
	}
}

func (prj project) removeLabelToMergeRequests(label *gitlab.Label) {
	mergeRequests, err := prj.getAllMergeRequests()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, mr := range mergeRequests {

			fmt.Println(mr.Title)
			fmt.Println(mr.IID)
			fmt.Println(mr.Labels)
			var labels []string = nil
			labels = append(labels, label.Name)
			prj.gitlabClient.MergeRequests.UpdateMergeRequest(prj.projectID, mr.IID, &gitlab.UpdateMergeRequestOptions{RemoveLabels: labels})
			fmt.Println("---------------------------------------------")

		}
	}

}
