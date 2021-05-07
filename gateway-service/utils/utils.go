package utils

import (
	"log"
	"net/http"

	"google.golang.org/grpc/metadata"
)

//HandleError handles any errors that might pop up
func HandleError(err error, w http.ResponseWriter) {
	log.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// JsonContentTypeMiddleware is middleware used to set the content type of all responses to application/json
func JsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// GrpcJWTMiddleware builds a middleware used to add the JWT containing user data in the request's context (at the moment it's not actually using JWT)
func GrpcJWTMiddleware(getUser func(*http.Request) string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(metadata.AppendToOutgoingContext(r.Context(), "token", getUser(r)))
			next.ServeHTTP(w, r)
		})
	}
}
