package issue_test

import (
	"context"
	"testing"
	"time"

	"github.com/Peshowe/issue-tracker/tracker-service/tracker/issue"
	mock_issue "github.com/Peshowe/issue-tracker/tracker-service/tracker/issue/mocks"
	mock_project "github.com/Peshowe/issue-tracker/tracker-service/tracker/project/mocks"
	"github.com/golang/mock/gomock"
	errs "github.com/pkg/errors"
)

func TestCreateIssueOk(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//create an IssueService with mock dependencies
	mockRepo := mock_issue.NewMockIssueRepository(mockCtrl)
	mockEventPublisher := mock_issue.NewMockEventPublisher(mockCtrl)
	mockProjectService := mock_project.NewMockProjectService(mockCtrl)
	testService := issue.NewIssueService(mockRepo, mockEventPublisher, mockProjectService)

	testIssue := issue.Issue{
		Name:      "Test Name",
		Desc:      "Test Desc",
		IssueType: "adhoc",
		Status:    "to do",
		User:      "user",
		Project:   "123",
	}
	ctx := context.Background()

	//expect these methods to be called once with given parameters
	mockProjectService.EXPECT().UserInProject(ctx, testIssue.Project).Return(true).Times(1)
	mockRepo.EXPECT().CreateIssue(ctx, &testIssue).Return(nil).Times(1)
	mockEventPublisher.EXPECT().PublishEvent(ctx, issue.IssueEvent{
		Type:  issue.IssueCreatedEventType,
		Issue: &testIssue,
	}).Return(nil).Times(1)

	//call the actual method
	err := testService.CreateIssue(ctx, &testIssue)

	//sleep for a second so that the component under test has time to call the mock event publisher
	time.Sleep(1 * time.Second)

	if err != nil {
		t.Errorf("Got unexpected error: %s", err)
	}
}

func TestCreateIssueErr(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//create an IssueService with mock dependencies
	mockRepo := mock_issue.NewMockIssueRepository(mockCtrl)
	mockEventPublisher := mock_issue.NewMockEventPublisher(mockCtrl)
	mockProjectService := mock_project.NewMockProjectService(mockCtrl)
	testService := issue.NewIssueService(mockRepo, mockEventPublisher, mockProjectService)

	testIssue := issue.Issue{
		Name:      "I",
		Desc:      "Invalid issue",
		IssueType: "adhoc",
		Status:    "to do",
		User:      "user",
		Project:   "123",
	}
	ctx := context.Background()

	//expect these methods to be called once with given parameters
	mockProjectService.EXPECT().UserInProject(ctx, testIssue.Project).Return(true).Times(1)

	//call the actual method
	err := testService.CreateIssue(ctx, &testIssue)
	invalidError := errs.Wrap(issue.ErrIssueInvalid, "service.Issue.CreateIssue")

	if err.Error() != invalidError.Error() {
		t.Errorf("Got wrong error: %s, Expected: %s", err, invalidError)
	}

}
