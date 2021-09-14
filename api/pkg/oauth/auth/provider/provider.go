package provider

import (
	"fmt"
	"os"

	"github.com/markbates/goth/providers/github"
)

type provider struct {
	Url          string
	ProfileUrl   string
	EmailUrl     string
	AuthUrl      string
	TokenUrl     string
	ClientId     string
	ClientSecret string
	CallbackUrl  string
}

func GithubProvider(AUTH_URL string) provider {
	githubAuth := provider{
		Url:          "https://github.com",
		ProfileUrl:   github.ProfileURL,
		EmailUrl:     github.EmailURL,
		ClientId:     os.Getenv("GITHUB_ID"),
		ClientSecret: os.Getenv("GITHUB_SECRET"),
		CallbackUrl:  fmt.Sprintf(AUTH_URL, "github"),
	}

	if os.Getenv("GHE_URL") != "" {
		githubAuth.Url = os.Getenv("GHE_URL")
		githubAuth.ProfileUrl = fmt.Sprintf("%s/api/v3/user", githubAuth.Url)
		githubAuth.EmailUrl = fmt.Sprintf("%s/api/v3/user/emails", githubAuth.Url)
	}

	githubAuth.AuthUrl = fmt.Sprintf("%s/login/oauth/authorize", githubAuth.Url)
	githubAuth.TokenUrl = fmt.Sprintf("%s/login/oauth/access_token", githubAuth.Url)

	return githubAuth
}
