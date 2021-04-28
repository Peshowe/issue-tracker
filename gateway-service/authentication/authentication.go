package authentication

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Peshowe/issue-tracker/gateway-service/utils"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

//store is a CookieStore where we'll be saving the user's authentication data
var store = sessions.NewCookieStore([]byte("secret")) //os.Getenv("SESSION_KEY")
var userSessionName string = "user-session"

var oauthConfig *oauth2.Config

//authExceptionPaths is the set of paths that do not require authentication (e.g. the login page)
var authExceptionPaths map[string]interface{} = make(map[string]interface{})

//LoginRedirect is the page to which we will redirect a user if they are not authenticated. By default it will automatically start the authentication (probably should be overwritten).
var LoginRedirect func(http.ResponseWriter, *http.Request) = authBegin

// AuthenticationMiddleware is a middleware that makes sure users are authenticated before they are able to visit the endpoints
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authenticationOk := true
		session, _ := store.Get(r, userSessionName)

		// if path is in authExceptionPaths, skip authentication
		if _, ok := authExceptionPaths[r.URL.Path]; !ok { //&& !strings.HasPrefix(r.URL.Path, "/static")

			// Make the necessary checks to see if user is authenticated
			if _, ok := session.Values["token"]; !ok {
				// user not authenticated at all
				authenticationOk = false
			} else {
				cached_token, _ := session.Values["token"].(oauth2.Token)
				token, err := oauthConfig.TokenSource(r.Context(), &cached_token).Token()
				if err != nil || !token.Valid() {
					// access token is no longer valid (TokenSource tries to refresh it automatically)
					authenticationOk = false
				}

			}
		}

		if !authenticationOk {
			// save the original path that's trying to be accessed, so we can redirect to it right away after authentication
			session.Values["original_path"] = r.URL.Path
			if err := session.Save(r, w); err != nil {
				utils.HandleError(errors.Wrap(err, "authentication.AuthenticationMiddleware.session.Save"), w)
				return
			}
			//redirect to a page that asks to login
			LoginRedirect(w, r)
		} else {
			// everything is ok, carry on
			next.ServeHTTP(w, r)
		}
	})
}

// GetUser returns the currently authenticated user from the session
func GetUser(r *http.Request) string {
	session, _ := store.Get(r, userSessionName)
	if user, ok := session.Values["user"]; ok {
		return user.(string)
	}
	return ""

}

// getUser exposes an API endpoint for GetUser
func getUser(w http.ResponseWriter, r *http.Request) {
	userResp := map[string]string{"user": GetUser(r)}
	json.NewEncoder(w).Encode(userResp)
}

// AddAuthExceptionPath adds a path to the authExceptionPaths set
func AddAuthExceptionPath(path string) error {
	authExceptionPaths[path] = nil
	return nil

}

// authInit does some initial setups for the oAuth2 config
func authInit() {
	oauthConfig = &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	// register the oauth2.Token type so we can store it in the cookie session
	gob.Register(oauth2.Token{})
}

// authBegin starts the authentication
func authBegin(w http.ResponseWriter, r *http.Request) {
	// generate a random state
	state := uuid.New().String()
	session, _ := store.Get(r, userSessionName)
	// Store the state in the session values.
	session.Values["state"] = state
	if err := session.Save(r, w); err != nil {
		utils.HandleError(errors.Wrap(err, "authentication.authBegin"), w)
		return
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := oauthConfig.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

// authCallback is where the service provider sends the user details
func authCallback(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, userSessionName)

	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"), session.Values["state"].(string), r.Context())
	if err != nil {
		utils.HandleError(errors.Wrap(err, "authentication.authCallback.getUserInfo"), w)
		// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	//Store the user and their token in the session store
	session.Values["user"] = content["email"]
	session.Values["token"] = content["token"]
	if err := session.Save(r, w); err != nil {
		utils.HandleError(errors.Wrap(err, "authentication.authCallback.session.Save"), w)
		return
	}

	if original_path, ok := session.Values["original_path"]; ok {
		//redirect to the original URL
		http.Redirect(w, r, original_path.(string), http.StatusTemporaryRedirect)
	} else {
		//if we don't have the original URL in the session, just go to index
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

// getUserInfo processes the data sent from the service provider
func getUserInfo(state string, code string, oauthState string, ctx context.Context) (map[string]interface{}, error) {

	// check if the state strings match
	if state != oauthState {
		return nil, fmt.Errorf("invalid oauth state")
	}

	// get access token
	token, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	// get the user info
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	// put the response in a map
	m := make(map[string]interface{})
	err = json.Unmarshal(contents, &m)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	// put the token in there as well
	m["token"] = token

	return m, nil
}

// logout logs out the current user by clearing their session
func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, userSessionName)
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		utils.HandleError(errors.Wrap(err, "authentication.logout.session.Save"), w)
		return
	}
}

// RegisterEndpoints registers the authentication endpoints in the router
func RegisterEndpoints(r chi.Router) {
	authInit()

	r.Route("/auth", func(r chi.Router) {

		r.Get("/google/callback", authCallback)
		r.Get("/google", authBegin)
		r.Get("/user", getUser)
		r.Get("/logout", logout)
	})

}
