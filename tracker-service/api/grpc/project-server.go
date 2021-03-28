package grpc

import (
	"context"
	"github.com/Peshowe/issue-tracker/tracker-service/grpc-contract/tracker-service/v1/project"
	pr "github.com/Peshowe/issue-tracker/tracker-service/tracker/project"
)

//projectServer implements the gRPC interface and calls the business logic for the Project subdomain
type projectServer struct {
	projectService pr.ProjectService
	project.UnimplementedProjectServiceServer
}

//encodeProject converts a *pr.Project type to a *project.ProjectResponse
func encodeProject(p *pr.Project) *project.ProjectResponse {
	return &project.ProjectResponse{
		Id:        p.Id,
		Name:      p.Name,
		Issues:    p.Issues,
		Users:     p.Users,
		CreatedOn: p.CreatedOn,
	}
}

//encodeProjects coverts an array of *pr.Project values to an array of *project.ProjectResponse
func encodeProjects(p []*pr.Project) []*project.ProjectResponse {
	projects := make([]*project.ProjectResponse, len(p))
	for i := range projects {
		projects[i] = encodeProject(p[i])
	}
	return projects
}

func (s *projectServer) GetProjectsAll(ctx context.Context, request *project.ProjectsAllRequest) (*project.ProjectsResponse, error) {
	projects, err := s.projectService.GetProjectsAll()
	if err != nil {
		return nil, err
	}

	projectsResponse := &project.ProjectsResponse{
		Projects: encodeProjects(projects),
	}

	return projectsResponse, nil
}

func (s *projectServer) GetProjectById(ctx context.Context, request *project.ProjectByIdRequest) (*project.ProjectResponse, error) {
	projectId := request.GetId()
	proj, err := s.projectService.GetProjectById(projectId)

	if err != nil {
		return nil, err
	}

	return encodeProject(proj), nil
}

func (s *projectServer) GetProjectsByUser(ctx context.Context, request *project.ProjectsByUserRequest) (*project.ProjectsResponse, error) {
	userId := request.GetUserId()
	projects, err := s.projectService.GetProjectsByUser(userId)
	if err != nil {
		return nil, err
	}

	projectsResponse := &project.ProjectsResponse{
		Projects: encodeProjects(projects),
	}

	return projectsResponse, nil
}

func (s *projectServer) CreateProject(ctx context.Context, request *project.CreateRequest) (*project.GenericResponse, error) {
	p := &pr.Project{
		Name: request.GetName(),
	}
	err := s.projectService.CreateProject(p)

	if err != nil {
		return nil, err
	}

	return &project.GenericResponse{}, nil
}

func (s *projectServer) DeleteProject(ctx context.Context, request *project.DeleteRequest) (*project.GenericResponse, error) {
	projectId := request.GetId()
	err := s.projectService.DeleteProject(projectId)

	if err != nil {
		return nil, err
	}

	return &project.GenericResponse{}, nil
}

func (s *projectServer) AddIssue(ctx context.Context, request *project.IssueRequest) (*project.GenericResponse, error) {
	err := s.projectService.AddIssue(request.GetProjectId(), request.GetIssueId())

	if err != nil {
		return nil, err
	}

	return &project.GenericResponse{}, nil
}

func (s *projectServer) RemoveIssue(ctx context.Context, request *project.IssueRequest) (*project.GenericResponse, error) {
	err := s.projectService.RemoveIssue(request.GetProjectId(), request.GetIssueId())

	if err != nil {
		return nil, err
	}

	return &project.GenericResponse{}, nil
}

func (s *projectServer) AddUser(ctx context.Context, request *project.UserRequest) (*project.GenericResponse, error) {
	err := s.projectService.AddUser(request.GetProjectId(), request.GetUserId())

	if err != nil {
		return nil, err
	}

	return &project.GenericResponse{}, nil
}

func (s *projectServer) RemoveUser(ctx context.Context, request *project.UserRequest) (*project.GenericResponse, error) {
	err := s.projectService.RemoveUser(request.GetProjectId(), request.GetUserId())

	if err != nil {
		return nil, err
	}

	return &project.GenericResponse{}, nil
}
