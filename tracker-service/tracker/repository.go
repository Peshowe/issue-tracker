package tracker

import (
	"github.com/Peshowe/issue-tracker/tracker-service/tracker/issue"
	"github.com/Peshowe/issue-tracker/tracker-service/tracker/project"
)

type TrackerRepository interface {
	project.ProjectRepository
	issue.IssueRepository

	// GetUsersAll() ([]*User, error)
	// GetUserById(id string) (*User, error)
	// CreateUser(user *User) error
	// UpdateUser(user *User) error
	// DeleteUser(id string) error
}
