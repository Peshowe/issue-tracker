package project

import (
	"errors"
	"time"

	errs "github.com/pkg/errors"
	"gopkg.in/dealancer/validate.v2"
)

// predefine some general errors
var (
	ErrProjectNotFound = errors.New("Project Not Found")
	ErrProjectInvalid  = errors.New("Project Invalid")
)

//projectService implements the business logic of its domain
type projectService struct {
	//reference to our repository (i.e. database)
	projectRepo ProjectRepository
}

//NewProjectService creates a new instance of the projectService that contains a reference to a repository
func NewProjectService(projectRepo ProjectRepository) ProjectService {
	return &projectService{
		projectRepo,
	}
}

//Project logic

func (r *projectService) GetProjectsAll() ([]*Project, error) {
	return r.projectRepo.GetProjectsAll()
}

func (r *projectService) GetProjectById(id string) (*Project, error) {
	return r.projectRepo.GetProjectById(id)
}

func (r *projectService) GetProjectsByUser(userId string) ([]*Project, error) {
	return r.projectRepo.GetProjectsByUser(userId)
}

func (r *projectService) CreateProject(project *Project) error {
	if err := validate.Validate(project); err != nil {
		return errs.Wrap(ErrProjectInvalid, "service.Project.CreateProject")
	}

	//add the timestamp
	project.CreatedOn = time.Now().UTC().Unix()

	return r.projectRepo.CreateProject(project)
}

func (r *projectService) DeleteProject(id string) error {
	return r.projectRepo.DeleteProject(id)
}

// func (r *projectService) UpdateProject(project *Project) error {
// 	if err := validate.Validate(project); err != nil {
// 		return errs.Wrap(ErrProjectInvalid, "service.Project.UpdateProject")
// 	}
// 	return r.projectRepo.UpdateProject(project)
// }

func (r *projectService) AddUser(projectId string, userId string) error {
	return r.projectRepo.AddUser(projectId, userId)
}

func (r *projectService) RemoveUser(projectId string, userId string) error {
	return r.projectRepo.RemoveUser(projectId, userId)
}
