package issue

import (
	"context"
	"errors"
	"time"

	"github.com/Peshowe/issue-tracker/tracker-service/tracker/project"
	"github.com/Peshowe/issue-tracker/tracker-service/utils"
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
	//eventPublisher is a reference to the logic for publishing domain events
	eventPublisher EventPublisher
	//we need a reference to the project service to deal with authorization (this is not ideal)
	projectService project.ProjectService
}

func NewIssueService(issueRepo IssueRepository, eventPublisher EventPublisher, projectService project.ProjectService) IssueService {
	return &issueService{
		issueRepo,
		eventPublisher,
		projectService,
	}
}

//Issue logic

//userAllowed returns whether a user can perform an action on given issue
func (r *issueService) userAllowed(ctx context.Context, issueId string) bool {
	project, _ := r.issueRepo.GetIssueById(ctx, issueId)

	return r.projectService.UserInProject(ctx, project.Id)
}

func (r *issueService) GetIssueById(ctx context.Context, id string) (*Issue, error) {
	if r.userAllowed(ctx, id) {
		return r.issueRepo.GetIssueById(ctx, id)
	} else {
		return nil, errs.Wrap(utils.ErrNotAllowed, "service.Issue.GetIssueById")
	}
}

func (r *issueService) GetIssuesByProject(ctx context.Context, projectId string) ([]*Issue, error) {
	if r.projectService.UserInProject(ctx, projectId) {
		return r.issueRepo.GetIssuesByProject(ctx, projectId)
	} else {
		return nil, errs.Wrap(utils.ErrNotAllowed, "service.Issue.GetIssuesByProject")
	}
}

func (r *issueService) GetIssuesByUser(ctx context.Context, userId string) ([]*Issue, error) {
	if utils.GetUserFromContext(ctx) == userId {
		return r.issueRepo.GetIssuesByUser(ctx, userId)
	} else {
		return nil, errs.Wrap(utils.ErrNotAllowed, "service.Issue.GetIssuesByUser")
	}
}

func (r *issueService) CreateIssue(ctx context.Context, issue *Issue) error {
	if r.projectService.UserInProject(ctx, issue.Project) {
		if err := validate.Validate(issue); err != nil {
			return errs.Wrap(ErrIssueInvalid, "service.Issue.CreateIssue")
		}

		//add the timestamps
		issue.CreatedOn = time.Now().UTC().Unix()
		issue.LastModifiedOn = time.Now().UTC().Unix()

		if err := r.issueRepo.CreateIssue(ctx, issue); err != nil {
			return errs.Wrap(err, "service.Issue.CreateIssue")
		} else {
			//publish the event
			go r.eventPublisher.PublishEvent(context.Background(), IssueEvent{
				Type:  IssueCreatedEventType,
				Issue: issue,
			})

			return nil
		}

	} else {
		return errs.Wrap(utils.ErrNotAllowed, "service.Issue.CreateIssue")
	}
}

func (r *issueService) PutIssue(ctx context.Context, issue *Issue) error {
	if r.projectService.UserInProject(ctx, issue.Project) {
		if err := validate.Validate(issue); err != nil {
			return errs.Wrap(ErrIssueInvalid, "service.Issue.PutIssue")
		}

		//update timestamp
		currentTime := time.Now().UTC().Unix()
		issue.LastModifiedOn = currentTime

		if err := r.issueRepo.PutIssue(ctx, issue); err != nil {
			return errs.Wrap(err, "service.Issue.PutIssue")
		} else {
			//publish the event
			go r.eventPublisher.PublishEvent(context.Background(), IssueEvent{
				Type:  IssueUpdatedEventType,
				Issue: issue,
			})

			return nil
		}
	} else {
		return errs.Wrap(utils.ErrNotAllowed, "service.Issue.PutIssue")
	}
}

func (r *issueService) DeleteIssue(ctx context.Context, id string) error {
	if r.userAllowed(ctx, id) {
		issue, _ := r.issueRepo.GetIssueById(ctx, id)
		if err := r.issueRepo.DeleteIssue(ctx, id); err != nil {
			return errs.Wrap(err, "service.Issue.DeleteIssue")
		} else {
			//publish the event
			go r.eventPublisher.PublishEvent(context.Background(), IssueEvent{
				Type:  IssueDeletedEventType,
				Issue: issue,
			})

			return nil
		}

	} else {
		return errs.Wrap(utils.ErrNotAllowed, "service.Issue.DeleteIssue")
	}
}
