package issue

import (
	"errors"
	"time"

	errs "github.com/pkg/errors"
	"gopkg.in/dealancer/validate.v2"
)

// predefine some very general errors (could be more specific like having different errors for Project, Issue, User)
var (
	ErrIssueNotFound = errors.New("Issue Not Found")
	ErrIssueInvalid  = errors.New("Issue Invalid")
)

type issueService struct {
	//reference to our repository (i.e. database)
	issueRepo IssueRepository
}

func NewIssueService(issueRepo IssueRepository) IssueService {
	return &issueService{
		issueRepo,
	}
}

//Issue logic

func (r *issueService) GetIssueById(id string) (*Issue, error) {
	return r.issueRepo.GetIssueById(id)
}

func (r *issueService) GetIssuesByProject(projectId string) ([]*Issue, error) {
	return r.issueRepo.GetIssuesByProject(projectId)
}

func (r *issueService) GetIssuesByUser(userId string) ([]*Issue, error) {
	return r.issueRepo.GetIssuesByUser(userId)
}

func (r *issueService) CreateIssue(issue *Issue) error {
	if err := validate.Validate(issue); err != nil {
		return errs.Wrap(ErrIssueInvalid, "service.Issue.CreateIssue")
	}

	//add the timestamps
	issue.CreatedOn = time.Now().UTC().Unix()
	issue.LastModifiedOn = time.Now().UTC().Unix()

	return r.issueRepo.CreateIssue(issue)
}

func (r *issueService) DeleteIssue(id string) error {
	return r.issueRepo.DeleteIssue(id)
}

func (r *issueService) UpdateStatus(issueId string, newStatus string) error {
	if err := r.updateLastModifiedOn(issueId); err != nil {
		return err
	}
	return r.issueRepo.UpdateStatus(issueId, newStatus)
}

func (r *issueService) UpdateUser(issueId string, userId string) error {
	if err := r.updateLastModifiedOn(issueId); err != nil {
		return err
	}
	return r.issueRepo.UpdateUser(issueId, userId)
}

func (r *issueService) UpdateDescription(issueId string, newDescription string) error {
	if err := r.updateLastModifiedOn(issueId); err != nil {
		return err
	}
	return r.issueRepo.UpdateDescription(issueId, newDescription)
}

func (r *issueService) UpdateBugTrace(issueId string, newBugTrace string) error {
	if err := r.updateLastModifiedOn(issueId); err != nil {
		return err
	}
	return r.issueRepo.UpdateBugTrace(issueId, newBugTrace)
}

func (r *issueService) updateLastModifiedOn(issueId string) error {
	currentTime := time.Now().UTC().Unix()
	return r.issueRepo.UpdateLastModifiedOn(issueId, currentTime)
}
