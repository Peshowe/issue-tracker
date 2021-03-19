package project

import (
	"io/ioutil"
	"log"
	// "context"
	"github.com/Peshowe/issue-tracker/grpc-contract/tracker-service/v1/project"
	// "google.golang.org/grpc"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net/http"
	// "github.com/go-chi/chi"
	// "github.com/pkg/errors"
)

type ProjectHandler interface {
	GetProjectsAll(http.ResponseWriter, *http.Request)
	// GetProjectById(http.ResponseWriter, *http.Request)
	// GetProjectsByUser(http.ResponseWriter, *http.Request)

	CreateProject(http.ResponseWriter, *http.Request)
	// DeleteProject(http.ResponseWriter, *http.Request)

	// //AddIssue adds an issue to the given project
	// AddIssue(http.ResponseWriter, *http.Request)
	// //RemoveIssue removes an issue from the given project
	// RemoveIssue(http.ResponseWriter, *http.Request)

	// //AddUser adds a user to the given project
	// AddUser(http.ResponseWriter, *http.Request)
	// //RemoveUser removes a user from the given project
	// RemoveUser(http.ResponseWriter, *http.Request)
}

type handler struct {
	//the gRPC client
	projectClient project.ProjectServiceClient
}

func NewProjectHandler(grpcConn grpc.ClientConnInterface) ProjectHandler {
	return &handler{projectClient: project.NewProjectServiceClient(grpcConn)}
}

//handleError handles any errors that might pop up
func handleError(err error, w http.ResponseWriter) {
	log.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (h *handler) GetProjectsAll(w http.ResponseWriter, r *http.Request) {
	resp, err := h.projectClient.GetProjectsAll(r.Context(), &project.ProjectsAllRequest{})
	if err != nil {
		handleError(errors.Wrap(err, "projectHandler.GetProjectsAll"), w)
		return
	}

	json.NewEncoder(w).Encode(resp)

}

func (h *handler) CreateProject(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(err, w)
		return
	}

	req := &project.CreateRequest{}
	if err := json.Unmarshal(requestBody, req); err != nil {
		handleError(errors.Wrap(err, "projectHandler.CreateProject"), w)
		return
	}
	resp, err := h.projectClient.CreateProject(r.Context(), req)
	if err != nil {
		handleError(errors.Wrap(err, "projectHandler.CreateProject"), w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)

}
