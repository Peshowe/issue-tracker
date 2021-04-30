package project

import "context"

type ProjectRepository interface {
	GetProjectsAll(ctx context.Context) ([]*Project, error)
	GetProjectById(ctx context.Context, id string) (*Project, error)
	GetProjectsByUser(ctx context.Context, userId string) ([]*Project, error)
	CreateProject(ctx context.Context, project *Project) error
	DeleteProject(ctx context.Context, id string) error

	//AddIssue adds an issue to the given project
	// AddIssue(projectId string, issueId string) error
	//RemoveIssue removes an issue from the given project
	// RemoveIssue(projectId string, issueId string) error

	//AddUser adds a user to the given project
	AddUser(ctx context.Context, projectId string, userId string) error
	//RemoveUser removes a user from the given project
	RemoveUser(ctx context.Context, projectId string, userId string) error
}
