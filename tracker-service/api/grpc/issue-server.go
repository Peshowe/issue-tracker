package grpc

import (
	"context"

	"github.com/Peshowe/issue-tracker/tracker-service/grpc-contract/tracker-service/v1/issue"
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

//decodeIssue converts a *issue.IssueResponse type to a *is.Issue
func decodeIssue(i *issue.IssueResponse) *is.Issue {
	return &is.Issue{
		Id:        i.Id,
		Name:      i.Name,
		Desc:      i.Desc,
		IssueType: i.IssueType,
		Status:    i.Status,
		BugTrace:  i.BugTrace,
		User:      i.User,
		Project:   i.Project,
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
	issueStruct := decodeIssue(request.GetIssue())
	err := s.issueService.CreateIssue(issueStruct)

	if err != nil {
		return nil, err
	}

	return &issue.GenericResponse{}, nil
}

func (s *issueServer) PutIssue(ctx context.Context, request *issue.PutRequest) (*issue.GenericResponse, error) {
	issueStruct := decodeIssue(request.GetIssue())
	err := s.issueService.PutIssue(issueStruct)

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
