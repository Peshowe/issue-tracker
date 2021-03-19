package issue

type IssueService interface {
	GetIssueById(id string) (*Issue, error)
	GetIssuesByProject(projectId string) ([]*Issue, error)
	GetIssuesByUser(userId string) ([]*Issue, error)
	CreateIssue(issue *Issue) error
	DeleteIssue(id string) error

	UpdateStatus(issueId string, newStatus string) error
	UpdateUser(issueId string, userId string) error
	UpdateDescription(issueId string, newDescription string) error
	UpdateBugTrace(issueId string, newBugTrace string) error
}
