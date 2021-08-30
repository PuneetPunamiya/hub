package auth

import (
	"fmt"
	"os"

	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/gitlab"
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

func BitBucketProvider(AUTH_URL string) provider {
	bitbucketAuth := provider{
		ClientId:     os.Getenv("BITBUCKET_ID"),
		ClientSecret: os.Getenv("BITBUCKET_SECRET"),
		CallbackUrl:  fmt.Sprintf(AUTH_URL, "bitbucket"),
	}

	return bitbucketAuth
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

func GitlabProvider(AUTH_URL string) provider {
	gitlabAuth := provider{
		Url:          "https://gitlab.com",
		AuthUrl:      gitlab.AuthURL,
		TokenUrl:     gitlab.TokenURL,
		ProfileUrl:   gitlab.ProfileURL,
		ClientId:     os.Getenv("GITLAB_ID"),
		ClientSecret: os.Getenv("GITLAB_SECRET"),
		CallbackUrl:  fmt.Sprintf(AUTH_URL, "gitlab"),
	}

	if os.Getenv("GLE_URL") != "" {
		gitlabAuth.Url = os.Getenv("GLE_URL")
		gitlabAuth.AuthUrl = fmt.Sprintf("%s/oauth/authorize", gitlabAuth.Url)
		gitlabAuth.TokenUrl = fmt.Sprintf("%s/oauth/token", gitlabAuth.Url)
		gitlabAuth.ProfileUrl = fmt.Sprintf("%s/api/v3/user", gitlabAuth.Url)
	}

	return gitlabAuth
}
