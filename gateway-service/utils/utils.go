package utils

import (
	"log"
	"net/http"
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
