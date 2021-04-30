package issue

import "context"

type IssueRepository interface {
	GetIssueById(ctx context.Context, id string) (*Issue, error)
	GetIssuesByProject(ctx context.Context, projectId string) ([]*Issue, error)
	GetIssuesByUser(ctx context.Context, userId string) ([]*Issue, error)
	CreateIssue(ctx context.Context, issue *Issue) error
	PutIssue(ctx context.Context, issue *Issue) error
	DeleteIssue(ctx context.Context, id string) error
}
