package issue

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Peshowe/issue-tracker/gateway-service/grpc-contract/tracker-service/v1/issue"
	"github.com/Peshowe/issue-tracker/gateway-service/utils"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type IssueHandler interface {
	GetIssueById(http.ResponseWriter, *http.Request)
	GetIssuesByProject(http.ResponseWriter, *http.Request)
	GetIssuesByUser(http.ResponseWriter, *http.Request)
	CreateIssue(http.ResponseWriter, *http.Request)
	PutIssue(http.ResponseWriter, *http.Request)
	DeleteIssue(http.ResponseWriter, *http.Request)
}

type handler struct {
	//the gRPC client
	issueClient issue.IssueServiceClient
}

func newIssueHandler(grpcConn grpc.ClientConnInterface) IssueHandler {
	return &handler{issueClient: issue.NewIssueServiceClient(grpcConn)}
}

//RegisterEndpoints registers the endpoints of our API for the issue subdomain
func RegisterEndpoints(r chi.Router, grpcConn grpc.ClientConnInterface) {

	issueHandler := newIssueHandler(grpcConn)

	r.Route("/issues", func(r chi.Router) {

		r.Get("/byid/{id}", issueHandler.GetIssueById)
		r.Get("/byproject/{projectId}", issueHandler.GetIssuesByProject)
		r.Get("/byuser/{userId}", issueHandler.GetIssuesByUser)

		r.Post("/", issueHandler.CreateIssue)
		r.Put("/", issueHandler.PutIssue)
		r.Delete("/", issueHandler.DeleteIssue)

	})

}

func (h *handler) GetIssueById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	resp, err := h.issueClient.GetIssueById(r.Context(), &issue.IssueByIdRequest{Id: id})
	if err != nil {
		utils.HandleError(errors.Wrap(err, "issueHandler.GetIssueById"), w)
		return
	}

	json.NewEncoder(w).Encode(resp)

}

func (h *handler) GetIssuesByProject(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	resp, err := h.issueClient.GetIssuesByProject(r.Context(), &issue.IssuesByProjectRequest{ProjectId: projectId})
	if err != nil {
		utils.HandleError(errors.Wrap(err, "issueHandler.GetIssuesByProject"), w)
		return
	}

	json.NewEncoder(w).Encode(resp)

}

func (h *handler) GetIssuesByUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	resp, err := h.issueClient.GetIssuesByUser(r.Context(), &issue.IssuesByUserRequest{UserId: userId})
	if err != nil {
		utils.HandleError(errors.Wrap(err, "issueHandler.GetIssuesByUser"), w)
		return
	}

	json.NewEncoder(w).Encode(resp)

}

func (h *handler) CreateIssue(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	req := &issue.CreateRequest{}
	if err := json.Unmarshal(requestBody, req); err != nil {
		utils.HandleError(errors.Wrap(err, "issueHandler.CreateIssue"), w)
		return
	}

	resp, err := h.issueClient.CreateIssue(r.Context(), req)
	if err != nil {
		utils.HandleError(errors.Wrap(err, "issueHandler.CreateIssue"), w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)

}

func (h *handler) PutIssue(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	req := &issue.PutRequest{}
	if err := json.Unmarshal(requestBody, req); err != nil {
		utils.HandleError(errors.Wrap(err, "issueHandler.PutIssue"), w)
		return
	}

	resp, err := h.issueClient.PutIssue(r.Context(), req)
	if err != nil {
		utils.HandleError(errors.Wrap(err, "issueHandler.PutIssue"), w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)

}

func (h *handler) DeleteIssue(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	req := &issue.DeleteRequest{}
	if err := json.Unmarshal(requestBody, req); err != nil {
		utils.HandleError(errors.Wrap(err, "issueHandler.DeleteIssue"), w)
		return
	}
	resp, err := h.issueClient.DeleteIssue(r.Context(), req)
	if err != nil {
		utils.HandleError(errors.Wrap(err, "issueHandler.DeleteIssue"), w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)

}
