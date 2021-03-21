package issue

import (
	"encoding/json"
	"github.com/Peshowe/issue-tracker/gateway/utils"
	"github.com/Peshowe/issue-tracker/grpc-contract/tracker-service/v1/issue"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
)

type IssueHandler interface {
	GetIssueById(http.ResponseWriter, *http.Request)
	GetIssuesByProject(http.ResponseWriter, *http.Request)
	GetIssuesByUser(http.ResponseWriter, *http.Request)
	CreateIssue(http.ResponseWriter, *http.Request)
	DeleteIssue(http.ResponseWriter, *http.Request)

	UpdateStatus(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
	UpdateDescription(http.ResponseWriter, *http.Request)
	UpdateBugTrace(http.ResponseWriter, *http.Request)
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
		r.Get("/byproject/{projectid}", issueHandler.GetIssuesByProject)
		r.Get("/byuser/{userid}", issueHandler.GetIssuesByUser)

		r.Post("/", issueHandler.CreateIssue)
		r.Delete("/", issueHandler.DeleteIssue)

		r.Put("/status", issueHandler.UpdateStatus)
		r.Put("/user", issueHandler.UpdateUser)
		r.Put("/description", issueHandler.UpdateDescription)
		r.Put("/bugtrace", issueHandler.UpdateBugTrace)

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

func (h *handler) UpdateStatus(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	req := &issue.UpdateStatusRequest{}
	if err := json.Unmarshal(requestBody, req); err != nil {
		utils.HandleError(errors.Wrap(err, "issueHandler.UpdateStatus"), w)
		return
	}
	resp, err := h.issueClient.UpdateStatus(r.Context(), req)
	if err != nil {
		utils.HandleError(errors.Wrap(err, "issueHandler.UpdateStatus"), w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)

}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	req := &issue.UpdateUserRequest{}
	if err := json.Unmarshal(requestBody, req); err != nil {
		utils.HandleError(errors.Wrap(err, "issueHandler.UpdateUser"), w)
		return
	}
	resp, err := h.issueClient.UpdateUser(r.Context(), req)
	if err != nil {
		utils.HandleError(errors.Wrap(err, "issueHandler.UpdateUser"), w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)

}

func (h *handler) UpdateDescription(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	req := &issue.UpdateDescriptionRequest{}
	if err := json.Unmarshal(requestBody, req); err != nil {
		utils.HandleError(errors.Wrap(err, "issueHandler.UpdateDescription"), w)
		return
	}
	resp, err := h.issueClient.UpdateDescription(r.Context(), req)
	if err != nil {
		utils.HandleError(errors.Wrap(err, "issueHandler.UpdateDescription"), w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)

}

func (h *handler) UpdateBugTrace(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	req := &issue.UpdateBugTraceRequest{}
	if err := json.Unmarshal(requestBody, req); err != nil {
		utils.HandleError(errors.Wrap(err, "issueHandler.UpdateBugTrace"), w)
		return
	}
	resp, err := h.issueClient.UpdateBugTrace(r.Context(), req)
	if err != nil {
		utils.HandleError(errors.Wrap(err, "issueHandler.UpdateBugTrace"), w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)

}
