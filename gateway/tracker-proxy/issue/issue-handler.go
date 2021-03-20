package issue

import (
	"io/ioutil"
	"log"
	// "context"
	"github.com/Peshowe/issue-tracker/grpc-contract/tracker-service/v1/issue"
	// "google.golang.org/grpc"
	"encoding/json"
	// "fmt"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net/http"
	// "github.com/pkg/errors"
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

func NewIssueHandler(grpcConn grpc.ClientConnInterface) IssueHandler {
	return &handler{issueClient: issue.NewIssueServiceClient(grpcConn)}
}
