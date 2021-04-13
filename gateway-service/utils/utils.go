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
