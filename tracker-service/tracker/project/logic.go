package project

import (
	"context"
	"errors"
	"time"

	"github.com/Peshowe/issue-tracker/tracker-service/utils"
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
	//eventPublisher is a reference to the logic for publishing domain events
	eventPublisher EventPublisher
}

//NewProjectService creates a new instance of the projectService that contains a reference to a repository
func NewProjectService(projectRepo ProjectRepository, eventPublisher EventPublisher) ProjectService {
	return &projectService{
		projectRepo,
		eventPublisher,
	}
}

// UserInProject checks if given user belongs to given project
func (r *projectService) UserInProject(ctx context.Context, projectId string) bool {
	user := utils.GetUserFromContext(ctx)
	project, _ := r.projectRepo.GetProjectById(ctx, projectId)

	for _, projectUser := range project.Users {
		if user == projectUser {
			return true
		}
	}
	return false
}

//Project logic

func (r *projectService) GetProjectsAll(ctx context.Context) ([]*Project, error) {

	// return r.projectRepo.GetProjectsAll(ctx)
	// return only the projects that the current user is part of
	return r.projectRepo.GetProjectsByUser(ctx, utils.GetUserFromContext(ctx))
}

func (r *projectService) GetProjectById(ctx context.Context, id string) (*Project, error) {
	if r.UserInProject(ctx, id) {
		return r.projectRepo.GetProjectById(ctx, id)
	} else {
		return nil, errs.Wrap(utils.ErrNotAllowed, "service.Project.GetProjectById")
	}
}

func (r *projectService) GetProjectsByUser(ctx context.Context, userId string) ([]*Project, error) {
	if utils.GetUserFromContext(ctx) == userId {
		return r.projectRepo.GetProjectsByUser(ctx, userId)
	} else {
		return nil, errs.Wrap(utils.ErrNotAllowed, "service.Project.GetProjectsByUser")
	}
}

func (r *projectService) CreateProject(ctx context.Context, project *Project) error {
	if err := validate.Validate(project); err != nil {
		return errs.Wrap(ErrProjectInvalid, "service.Project.CreateProject")
	}

	//add the timestamp
	project.CreatedOn = time.Now().UTC().Unix()

	//add the current user to the list of users for the new project
	project.Users = append(project.Users, utils.GetUserFromContext(ctx))

	return r.projectRepo.CreateProject(ctx, project)
}

func (r *projectService) DeleteProject(ctx context.Context, id string) error {
	if r.UserInProject(ctx, id) {
		return r.projectRepo.DeleteProject(ctx, id)
	} else {
		return errs.Wrap(utils.ErrNotAllowed, "service.Project.DeleteProject")
	}

}

// func (r *projectService) UpdateProject(project *Project) error {
// 	if err := validate.Validate(project); err != nil {
// 		return errs.Wrap(ErrProjectInvalid, "service.Project.UpdateProject")
// 	}
// 	return r.projectRepo.UpdateProject(project)
// }

func (r *projectService) AddUser(ctx context.Context, projectId string, userId string) error {
	if r.UserInProject(ctx, projectId) {
		return r.projectRepo.AddUser(ctx, projectId, userId)
	} else {
		return errs.Wrap(utils.ErrNotAllowed, "service.Project.AddUser")
	}
}

func (r *projectService) RemoveUser(ctx context.Context, projectId string, userId string) error {
	if r.UserInProject(ctx, projectId) {
		return r.projectRepo.RemoveUser(ctx, projectId, userId)
	} else {
		return errs.Wrap(utils.ErrNotAllowed, "service.Project.RemoveUser")
	}
}
