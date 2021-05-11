package mailer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Peshowe/issue-tracker/gateway-service/grpc-contract/mail-service/v1/mailer"
	"github.com/Peshowe/issue-tracker/gateway-service/utils"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type MailerHandler interface {
	GetUserPreference(http.ResponseWriter, *http.Request)
	SetUserPreference(http.ResponseWriter, *http.Request)
}

type handler struct {
	//the gRPC client
	mailClient mailer.MailServiceClient
}

func newMailerHandler(grpcConn grpc.ClientConnInterface) MailerHandler {
	return &handler{mailClient: mailer.NewMailServiceClient(grpcConn)}
}

//RegisterEndpoints registers the endpoints of our API for the mailer subdomain
func RegisterEndpoints(r chi.Router, grpcConn grpc.ClientConnInterface) {

	mailerHandler := newMailerHandler(grpcConn)

	r.Route("/mailer", func(r chi.Router) {

		r.Get("/preferences/{user}", mailerHandler.GetUserPreference)
		r.Put("/preferences", mailerHandler.SetUserPreference)

	})

}

func (h *handler) GetUserPreference(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "user")
	resp, err := h.mailClient.GetUserPreference(r.Context(), &mailer.UserPreferenceRequest{User: user})
	if err != nil {
		utils.HandleError(errors.Wrap(err, "mailerHandler.GetUserPreference"), w)
		return
	}

	json.NewEncoder(w).Encode(resp)

}

func (h *handler) SetUserPreference(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HandleError(err, w)
		return
	}

	req := &mailer.UserPreference{}
	if err := json.Unmarshal(requestBody, req); err != nil {
		utils.HandleError(errors.Wrap(err, "mailerHandler.SetUserPreference"), w)
		return
	}

	resp, err := h.mailClient.SetUserPreference(r.Context(), req)
	if err != nil {
		utils.HandleError(errors.Wrap(err, "mailerHandler.SetUserPreference"), w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)

}
