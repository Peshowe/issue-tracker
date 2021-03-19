package tracker

// import (
// 	"errors"
// 	errs "github.com/pkg/errors"
// 	"gopkg.in/dealancer/validate.v2"
// 	"time"
// )

// // predefine some very general errors (could be more specific like having different errors for Project, Issue, User)
// var (
// 	ErrObjectNotFound = errors.New("Object Not Found")
// 	ErrObjectInvalid  = errors.New("Object Invalid")
// )

// type trackerService struct {
// 	//reference to our repository (i.e. database)
// 	trackerRepo TrackerRepository
// }

// func NewTrackerService(trackerRepo TrackerRepository) TrackerService {
// 	return &trackerService{
// 		trackerRepo,
// 	}
// }

// //Project logic

// func (r *trackerService) GetProjectsAll() ([]*Project, error) {
// 	return r.trackerRepo.GetProjectsAll()
// }

// func (r *trackerService) GetProjectById(id string) (*Project, error) {
// 	return r.trackerRepo.GetProjectById(id)
// }

// func (r *trackerService) GetProjectsByUser(userId string) ([]*Project, error) {
// 	return r.trackerRepo.GetProjectsByUser(userId)
// }

// func (r *trackerService) CreateProject(project *Project) error {
// 	if err := validate.Validate(project); err != nil {
// 		return errs.Wrap(ErrObjectInvalid, "service.Tracker.CreateProject")
// 	}

// 	//add the timestamp
// 	project.CreatedOn = time.Now().UTC().Unix()

// 	return r.trackerRepo.CreateProject(project)
// }

// func (r *trackerService) UpdateProject(project *Project) error {
// 	if err := validate.Validate(project); err != nil {
// 		return errs.Wrap(ErrObjectInvalid, "service.Tracker.UpdateProject")
// 	}
// 	return r.trackerRepo.UpdateProject(project)
// }

// func (r *trackerService) DeleteProject(id string) error {
// 	return r.trackerRepo.DeleteProject(id)
// }

// //Issue logic

// func (r *trackerService) GetIssueById(id string) (*Issue, error) {
// 	return r.trackerRepo.GetIssueById(id)
// }

// func (r *trackerService) GetIssuesByProject(projectId string) ([]*Issue, error) {
// 	return r.trackerRepo.GetIssuesByProject(projectId)
// }

// func (r *trackerService) GetIssuesByUser(userId string) ([]*Issue, error) {
// 	return r.trackerRepo.GetIssuesByUser(userId)
// }

// func (r *trackerService) CreateIssue(issue *Issue) error {
// 	if err := validate.Validate(issue); err != nil {
// 		return errs.Wrap(ErrObjectInvalid, "service.Tracker.CreateIssue")
// 	}

// 	//add the timestamps
// 	issue.CreatedOn = time.Now().UTC().Unix()
// 	issue.LastModifiedOn = time.Now().UTC().Unix()

// 	return r.trackerRepo.CreateIssue(issue)
// }

// func (r *trackerService) UpdateIssue(issue *Issue) error {
// 	if err := validate.Validate(issue); err != nil {
// 		return errs.Wrap(ErrObjectInvalid, "service.Tracker.UpdateIssue")
// 	}

// 	//update the timestamp
// 	issue.LastModifiedOn = time.Now().UTC().Unix()

// 	return r.trackerRepo.UpdateIssue(issue)
// }

// func (r *trackerService) DeleteIssue(id string) error {
// 	return r.trackerRepo.DeleteIssue(id)
// }

// //User logic

// func (r *trackerService) GetUsersAll() ([]*User, error) {
// 	return r.trackerRepo.GetUsersAll()
// }

// func (r *trackerService) GetUserById(id string) (*User, error) {
// 	return r.trackerRepo.GetUserById(id)
// }

// func (r *trackerService) CreateUser(user *User) error {
// 	if err := validate.Validate(user); err != nil {
// 		return errs.Wrap(ErrObjectInvalid, "service.Tracker.CreateUser")
// 	}

// 	return r.trackerRepo.CreateUser(user)
// }

// func (r *trackerService) UpdateUser(user *User) error {
// 	if err := validate.Validate(user); err != nil {
// 		return errs.Wrap(ErrObjectInvalid, "service.Tracker.UpdateUser")
// 	}

// 	return r.trackerRepo.UpdateUser(user)
// }

// func (r *trackerService) DeleteUser(id string) error {
// 	return r.trackerRepo.DeleteUser(id)
// }
