package gateway

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Authenticator interface {
	//AuthenticationMiddleware is a middleware that makes sure users are authenticated before they are able to visit the endpoints
	AuthenticationMiddleware(http.Handler) http.Handler
	//GetUser should return the currently authenticated user from the session
	GetUser(*http.Request) string
	//SetLoginRedirect overrides the default loginRedirect with a custom login page
	SetLoginRedirect(func(http.ResponseWriter, *http.Request))
	//AddAuthExceptionPath should add a path to a set of excempt from authentication paths
	AddAuthExceptionPath(string) error
	//RegisterEndpoints should register the authentication endpoints
	RegisterEndpoints(chi.Router)
}
