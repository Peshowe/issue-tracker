package project

import "context"

type ProjectService interface {
	UserInProject(ctx context.Context, projectId string) bool
	GetProjectsAll(ctx context.Context) ([]*Project, error)
	GetProjectById(ctx context.Context, id string) (*Project, error)
	GetProjectsByUser(ctx context.Context, userId string) ([]*Project, error)
	CreateProject(ctx context.Context, project *Project) error
	DeleteProject(ctx context.Context, id string) error

	//AddUser adds a user to the given project
	AddUser(ctx context.Context, projectId string, userId string) error
	//RemoveUser removes a user from the given project
	RemoveUser(ctx context.Context, projectId string, userId string) error
}
