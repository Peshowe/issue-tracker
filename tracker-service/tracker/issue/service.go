package issue

type IssueService interface {
	GetIssueById(id string) (*Issue, error)
	GetIssuesByProject(projectId string) ([]*Issue, error)
	GetIssuesByUser(userId string) ([]*Issue, error)
	CreateIssue(issue *Issue) error
	PutIssue(issue *Issue) error
	DeleteIssue(id string) error
}
