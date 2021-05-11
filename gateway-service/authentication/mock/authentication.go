package mock

//A mock implementation of gateway.Authenticator, always authenticated and returns the mockUser that it's intialised with

import (
	"net/http"

	"github.com/Peshowe/issue-tracker/gateway-service/gateway"
	"github.com/go-chi/chi"
)

type authenticator struct {
	mockUser string
}

func NewAuthenticator(mockUser string) gateway.Authenticator {
	return &authenticator{mockUser}
}

//AuthenticationMiddleware is a middleware that makes sure users are authenticated before they are able to visit the endpoints
func (a *authenticator) AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

//GetUser should return the currently authenticated user from the session
func (a *authenticator) GetUser(r *http.Request) string {
	return a.mockUser
}

//SetLoginRedirect overrides the default loginRedirect with a custom login page
func (a *authenticator) SetLoginRedirect(func(http.ResponseWriter, *http.Request)) {

}

//AddAuthExceptionPath should add a path to a set of excempt from authentication paths
func (a *authenticator) AddAuthExceptionPath(string) error {
	return nil
}

//RegisterEndpoints should register the authentication endpoints
func (a *authenticator) RegisterEndpoints(chi.Router) {

}
