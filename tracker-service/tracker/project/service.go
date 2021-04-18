package project

type ProjectService interface {
	GetProjectsAll() ([]*Project, error)
	GetProjectById(id string) (*Project, error)
	GetProjectsByUser(userId string) ([]*Project, error)
	CreateProject(project *Project) error
	DeleteProject(id string) error

	//AddUser adds a user to the given project
	AddUser(projectId string, userId string) error
	//RemoveUser removes a user from the given project
	RemoveUser(projectId string, userId string) error
}
