package grpc

import (
	"context"
	"github.com/Peshowe/issue-tracker/grpc-contract/tracker-service/v1/issue"
	is "github.com/Peshowe/issue-tracker/tracker-service/tracker/issue"
)

//issueServer implements the gRPC interface and calls the business logic for the Project subdomain
type issueServer struct {
	issueService is.IssueService
	issue.UnimplementedIssueServiceServer
}

//encodeIssue converts a *is.Issue type to a *issue.IssueResponse
func encodeIssue(i *is.Issue) *issue.IssueResponse {
	return &issue.IssueResponse{
		Id:             i.Id,
		Name:           i.Name,
		Desc:           i.Desc,
		IssueType:      i.IssueType,
		Status:         i.Status,
		BugTrace:       i.BugTrace,
		User:           i.User,
		Project:        i.Project,
		CreatedOn:      i.CreatedOn,
		LastModifiedOn: i.LastModifiedOn,
	}
}

//encodeIssues coverts an array of *is.Issue values to an array of *issue.IssueResponse
func encodeIssues(i_slice []*is.Issue) []*issue.IssueResponse {
	issues := make([]*issue.IssueResponse, len(i_slice))
	for i := range issues {
		issues[i] = encodeIssue(i_slice[i])
	}
	return issues
}

func (s *issueServer) GetIssueById(ctx context.Context, request *issue.IssueByIdRequest) (*issue.IssueResponse, error) {
	issueId := request.GetId()
	issueResp, err := s.issueService.GetIssueById(issueId)

	if err != nil {
		return nil, err
	}

	return encodeIssue(issueResp), nil
}

func (s *issueServer) GetIssuesByProject(ctx context.Context, request *issue.IssuesByProjectRequest) (*issue.IssuesResponse, error) {
	projectId := request.GetProjectId()
	issues, err := s.issueService.GetIssuesByProject(projectId)
	if err != nil {
		return nil, err
	}

	issuesResponse := &issue.IssuesResponse{
		Issues: encodeIssues(issues),
	}

	return issuesResponse, nil
}

func (s *issueServer) GetIssuesByUser(ctx context.Context, request *issue.IssuesByUserRequest) (*issue.IssuesResponse, error) {
	userId := request.GetUserId()
	issues, err := s.issueService.GetIssuesByUser(userId)
	if err != nil {
		return nil, err
	}

	issuesResponse := &issue.IssuesResponse{
		Issues: encodeIssues(issues),
	}

	return issuesResponse, nil
}

func (s *issueServer) CreateIssue(ctx context.Context, request *issue.CreateRequest) (*issue.GenericResponse, error) {
	i := request.GetIssue()
	issueStruct := &is.Issue{
		Name:      i.Name,
		Desc:      i.Desc,
		IssueType: i.IssueType,
		Status:    i.Status,
		BugTrace:  i.BugTrace,
		User:      i.User,
		Project:   i.Project,
	}
	err := s.issueService.CreateIssue(issueStruct)

	if err != nil {
		return nil, err
	}

	return &issue.GenericResponse{}, nil
}

func (s *issueServer) DeleteIssue(ctx context.Context, request *issue.DeleteRequest) (*issue.GenericResponse, error) {
	projectId := request.GetId()
	err := s.issueService.DeleteIssue(projectId)

	if err != nil {
		return nil, err
	}

	return &issue.GenericResponse{}, nil
}

func (s *issueServer) UpdateStatus(ctx context.Context, request *issue.UpdateStatusRequest) (*issue.GenericResponse, error) {
	err := s.issueService.UpdateStatus(request.GetIssueId(), request.GetNewStatus())

	if err != nil {
		return nil, err
	}

	return &issue.GenericResponse{}, nil
}

func (s *issueServer) UpdateUser(ctx context.Context, request *issue.UpdateUserRequest) (*issue.GenericResponse, error) {
	err := s.issueService.UpdateUser(request.GetIssueId(), request.GetUserId())

	if err != nil {
		return nil, err
	}

	return &issue.GenericResponse{}, nil
}

func (s *issueServer) UpdateDescription(ctx context.Context, request *issue.UpdateDescriptionRequest) (*issue.GenericResponse, error) {
	err := s.issueService.UpdateDescription(request.GetIssueId(), request.GetNewDescription())

	if err != nil {
		return nil, err
	}

	return &issue.GenericResponse{}, nil
}

func (s *issueServer) UpdateBugTrace(ctx context.Context, request *issue.UpdateBugTraceRequest) (*issue.GenericResponse, error) {
	err := s.issueService.UpdateBugTrace(request.GetIssueId(), request.GetNewBugTrace())

	if err != nil {
		return nil, err
	}

	return &issue.GenericResponse{}, nil
}
