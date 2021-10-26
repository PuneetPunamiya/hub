package provider

import (
	"fmt"
	"os"

	"github.com/markbates/goth/providers/gitlab"
)

func GitlabProvider(AUTH_URL string) provider {
	gitlabAuth := provider{
		Url:          "https://gitlab.com",
		AuthUrl:      gitlab.AuthURL,
		TokenUrl:     gitlab.TokenURL,
		ProfileUrl:   gitlab.ProfileURL,
		ClientId:     os.Getenv("GL_CLIENT_ID"),
		ClientSecret: os.Getenv("GL_CLIENT_SECRET"),
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
