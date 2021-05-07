package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Peshowe/issue-tracker/gateway-service/authentication"
	"github.com/Peshowe/issue-tracker/gateway-service/frontend"
	"github.com/Peshowe/issue-tracker/gateway-service/tracker-proxy/issue"
	"github.com/Peshowe/issue-tracker/gateway-service/tracker-proxy/project"
	"github.com/Peshowe/issue-tracker/gateway-service/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("tracker-service:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	authentication.LoginRedirect = frontend.LoginRedirect
	authentication.AddAuthExceptionPath("/login")

	// register the authentication endpoints
	authentication.RegisterEndpoints(r)

	// register the tracker service (i.e. business logic) API endpoints
	r.Route("/v1", func(r chi.Router) {
		r.Use(utils.JsonContentTypeMiddleware)
		r.Use(utils.GrpcJWTMiddleware(authentication.GetUser))
		project.RegisterEndpoints(r, conn)
		issue.RegisterEndpoints(r, conn)
	})

	r.Route("/", func(r chi.Router) {
		frontend.RegisterEndpoints(r, authentication.AuthenticationMiddleware)
	})

	errs := make(chan error, 2)
	go func() {
		port := httpPort()
		fmt.Println("Listening on port", port)
		errs <- http.ListenAndServe(port, r)

	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)

}

func httpPort() string {
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}
