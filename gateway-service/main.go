package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Peshowe/issue-tracker/gateway-service/authentication/mock"
	"github.com/Peshowe/issue-tracker/gateway-service/authentication/oauth2"
	"github.com/Peshowe/issue-tracker/gateway-service/frontend"
	"github.com/Peshowe/issue-tracker/gateway-service/gateway"
	"github.com/Peshowe/issue-tracker/gateway-service/mailer-proxy/mailer"
	"github.com/Peshowe/issue-tracker/gateway-service/tracker-proxy/issue"
	"github.com/Peshowe/issue-tracker/gateway-service/tracker-proxy/project"
	"github.com/Peshowe/issue-tracker/gateway-service/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"google.golang.org/grpc"
)

func main() {
	trackerGrpcConn, err := grpc.Dial("tracker-service:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	mailerGrpcConn, err := grpc.Dial("mail-service:4040", grpc.WithInsecure())
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

	var authenticator gateway.Authenticator

	if mockUser, ok := os.LookupEnv("MOCK_USER"); ok {
		log.Println("Using mock authentication!")
		authenticator = mock.NewAuthenticator(mockUser)
	} else {
		authenticator = oauth2.NewAuthenticator()
	}

	authenticator.SetLoginRedirect(frontend.LoginRedirect)
	authenticator.AddAuthExceptionPath("/login")

	// register the authenticator endpoints
	r.Route("/auth", func(r chi.Router) {

		authenticator.RegisterEndpoints(r)
		//expose an API endpoint for GetUser
		r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
			userResp := map[string]string{"user": authenticator.GetUser(r)}
			json.NewEncoder(w).Encode(userResp)
		})
	})

	// register the tracker service (i.e. business logic) API endpoints
	r.Route("/v1", func(r chi.Router) {
		r.Use(utils.JsonContentTypeMiddleware)
		r.Use(utils.GrpcJWTMiddleware(authenticator.GetUser))
		project.RegisterEndpoints(r, trackerGrpcConn)
		issue.RegisterEndpoints(r, trackerGrpcConn)
		mailer.RegisterEndpoints(r, mailerGrpcConn)
	})

	r.Route("/", func(r chi.Router) {
		frontend.RegisterEndpoints(r, authenticator.AuthenticationMiddleware)
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
