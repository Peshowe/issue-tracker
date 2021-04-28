package frontend

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
)

// workDir is the cwd
var workDir string

// RegisterEndpoints registers the frontend endpoints (including the static dir) and puts an authentication middleware where necessary
func RegisterEndpoints(r chi.Router, authMiddleware func(http.Handler) http.Handler) {

	workDir, _ = os.Getwd()

	// Create a route along that will serve contents from
	// the ./frontend/static directory.
	filesDir := http.Dir(filepath.Join(workDir, "frontend", "build", "static"))
	fileServer(r, "/static", filesDir)

	r.Route("/", func(r chi.Router) {
		// Serve index.html on all other requests to /*, which are not for the static dir and run the authentication middleware here
		r.Use(authMiddleware)
		r.Get("/*", indexServer)
	})

}

func LoginRedirect(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path = "/login"
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

// indexServer only serves the index.html file
func indexServer(w http.ResponseWriter, r *http.Request) {

	p := filepath.Join(workDir, "frontend", "build", "index.html")
	http.ServeFile(w, r, p)
}

// fileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}

	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
