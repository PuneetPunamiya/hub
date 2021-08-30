package auth

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/bitbucket"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/gitlab"
	"github.com/tektoncd/hub/api/pkg/app"
)

func AuthProvider(r *mux.Router, api app.Config) {

	key := ""            // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30 // 30 days
	isProd := false      // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	var AUTH_BASE_URL = os.Getenv("AUTH_BASE_URL")
	var AUTH_URL = AUTH_BASE_URL + "/auth/%s/callback"

	bitbucketAuth := BitBucketProvider(AUTH_URL)
	githubAuth := GithubProvider(AUTH_URL)
	gitlabAuth := GitlabProvider(AUTH_URL)

	goth.UseProviders(
		github.NewCustomisedURL(
			githubAuth.ClientId,
			githubAuth.ClientSecret,
			githubAuth.CallbackUrl,
			githubAuth.AuthUrl,
			githubAuth.TokenUrl,
			githubAuth.ProfileUrl,
			githubAuth.EmailUrl),

		bitbucket.New(
			bitbucketAuth.ClientId,
			bitbucketAuth.ClientSecret,
			bitbucketAuth.CallbackUrl,
		),

		gitlab.NewCustomisedURL(
			gitlabAuth.ClientId,
			gitlabAuth.ClientSecret,
			gitlabAuth.CallbackUrl,
			gitlabAuth.AuthUrl,
			gitlabAuth.TokenUrl,
			gitlabAuth.ProfileUrl,
		),
	)

	authSvc := New(api)

	r.HandleFunc("/", Status)
	s := r.PathPrefix("/auth").Subrouter()

	s.HandleFunc("/providers", List)

	s.HandleFunc("/login", authSvc.HubAuthenticate).Methods(http.MethodPost)

	s.HandleFunc("/{provider}/callback", authSvc.AuthCallBack)

	s.HandleFunc("/{provider}", Authenticate)
}
