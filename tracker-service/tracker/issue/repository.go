package issue

type IssueRepository interface {
	GetIssueById(id string) (*Issue, error)
	// GetIssuesById(ids []string) ([]*Issue, error)
	GetIssuesByProject(projectId string) ([]*Issue, error)
	GetIssuesByUser(userId string) ([]*Issue, error)
	CreateIssue(issue *Issue) error
	DeleteIssue(id string) error

	UpdateStatus(issueId string, newStatus string) error
	UpdateDescription(issueId string, newDescription string) error
	UpdateBugTrace(issueId string, newBugTrace string) error
	UpdateUser(issueId string, userId string) error
	UpdateLastModifiedOn(string, int64) error
}
