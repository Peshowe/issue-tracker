package project

import (
	"encoding/json"
	"github.com/Peshowe/issue-tracker/gateway/grpc-contract/tracker-service/v1/project"
	"github.com/Peshowe/issue-tracker/gateway/utils"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
)

type ProjectHandler interface {
	GetProjectsAll(http.ResponseWriter, *http.Request)
	GetProjectById(http.ResponseWriter, *http.Request)
	GetProjectsByUser(http.ResponseWriter, *http.Request)

	CreateProject(http.ResponseWriter, *http.Request)
	DeleteProject(http.ResponseWriter, *http.Request)

	//AddIssue adds an issue to the given project
	AddIssue(http.ResponseWriter, *http.Request)
	//RemoveIssue removes an issue from the given project
	RemoveIssue(http.ResponseWriter, *http.Request)

	//AddUser adds a user to the given project
	AddUser(http.ResponseWriter, *http.Request)
	//RemoveUser removes a user from the given project
	RemoveUser(http.ResponseWriter, *http.Request)
}

type handler struct {
	//the gRPC client
	projectClient project.ProjectServiceClient
}

func newProjectHandler(grpcConn grpc.ClientConnInterface) ProjectHandler {
	return &handler{projectClient: project.NewProjectServiceClient(grpcConn)}
}

//RegisterEndpoints registers the endpoints of our API for the project subdomain
func RegisterEndpoints(r chi.Router, grpcConn grpc.ClientConnInterface) {

	projectHandler := newProjectHandler(grpcConn)
	r.Route("/projects", func(r chi.Router) {

		r.Get("/", projectHandler.GetProjectsAll)
		r.Get("/byid/{id}", projectHandler.GetProjectById)
		r.Get("/byuser/{userid}", projectHandler.GetProjectsByUser)

		r.Post("/", projectHandler.CreateProject)
		r.Delete("/", projectHandler.DeleteProject)

		//issue subroutes
		r.Route("/issues", func(r chi.Router) {
			r.Post("/", projectHandler.AddIssue)
			r.Delete("/", projectHandler.RemoveIssue)
		})

		//user subroutes
		r.Route("/users", func(r chi.Router) {
			r.Post("/", projectHandler.AddUser)
			r.Delete("/", projectHandler.RemoveUser)
		})

	})

}

func (h *handler) GetProjectsAll(w http.ResponseWriter, r *http.Request) {
	resp, err := h.projectClient.GetProjectsAll(r.Context(), &project.ProjectsAllRequest{})
	if err != nil {
		utils.HandleError(errors.Wrap(err, "projectHandler.GetProjectsAll"), w)
		return
	}

	json.NewEncoder(w).Encode(resp)

}

func (h *handler) GetProjectById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	resp, err := h.projectClient.GetProjectById(r.Context(), &project.ProjectByIdRequest{Id: id})
	if err != nil {
		utils.HandleError(errors.Wrap(err, "projectHandler.GetProjectById"), w)
		return
	}

	json.NewEncoder(w).Encode(resp)

}

func (h *handler) GetProjectsByUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	resp, err := h.projectClient.GetProjectsByUser(r.Context(), &project.ProjectsByUserRequest{UserId: userId})
	if err != nil {
		utils.HandleError(errors.Wrap(err, "projectHandler.GetProjectsByUser"), w)
		return
	}

	json.NewEncoder(w).Encode(resp)

}

func (h *handler) CreateProject(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	req := &project.CreateRequest{}
	if err := json.Unmarshal(requestBody, req); err != nil {
		utils.HandleError(errors.Wrap(err, "projectHandler.CreateProject"), w)
		return
	}
	resp, err := h.projectClient.CreateProject(r.Context(), req)
	if err != nil {
		utils.HandleError(errors.Wrap(err, "projectHandler.CreateProject"), w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)

}

func (h *handler) DeleteProject(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	req := &project.DeleteRequest{}
	if err := json.Unmarshal(requestBody, req); err != nil {
		utils.HandleError(errors.Wrap(err, "projectHandler.DeleteProject"), w)
		return
	}
	resp, err := h.projectClient.DeleteProject(r.Context(), req)
	if err != nil {
		utils.HandleError(errors.Wrap(err, "projectHandler.DeleteProject"), w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)

}

func (h *handler) AddIssue(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	req := &project.IssueRequest{}
	if err := json.Unmarshal(requestBody, req); err != nil {
		utils.HandleError(errors.Wrap(err, "projectHandler.AddIssue"), w)
		return
	}
	resp, err := h.projectClient.AddIssue(r.Context(), req)
	if err != nil {
		utils.HandleError(errors.Wrap(err, "projectHandler.AddIssue"), w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)

}

func (h *handler) RemoveIssue(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	req := &project.IssueRequest{}
	if err := json.Unmarshal(requestBody, req); err != nil {
		utils.HandleError(errors.Wrap(err, "projectHandler.RemoveIssue"), w)
		return
	}
	resp, err := h.projectClient.RemoveIssue(r.Context(), req)
	if err != nil {
		utils.HandleError(errors.Wrap(err, "projectHandler.RemoveIssue"), w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)

}

func (h *handler) AddUser(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	req := &project.UserRequest{}
	if err := json.Unmarshal(requestBody, req); err != nil {
		utils.HandleError(errors.Wrap(err, "projectHandler.AddUser"), w)
		return
	}
	resp, err := h.projectClient.AddUser(r.Context(), req)
	if err != nil {
		utils.HandleError(errors.Wrap(err, "projectHandler.AddUser"), w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)

}

func (h *handler) RemoveUser(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	req := &project.UserRequest{}
	if err := json.Unmarshal(requestBody, req); err != nil {
		utils.HandleError(errors.Wrap(err, "projectHandler.RemoveUser"), w)
		return
	}
	resp, err := h.projectClient.RemoveUser(r.Context(), req)
	if err != nil {
		utils.HandleError(errors.Wrap(err, "projectHandler.RemoveUser"), w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)

}
