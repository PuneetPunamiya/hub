package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/tektoncd/hub/api/gen/log"
	"github.com/tektoncd/hub/api/pkg/app"
	"github.com/tektoncd/hub/api/pkg/db/model"
	"gorm.io/gorm"
)

type service struct {
	app.Service
	api app.Config
}

type request struct {
	db            *gorm.DB
	log           *log.Logger
	defaultScopes []string
	jwtConfig     *app.JWTConfig
}

type Provider struct {
	Name string `json:"name"`
}

type ProviderList struct {
	Data []Provider `json:"data"`
}

type AuthService struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Services struct {
	Service AuthService `json:"services"`
}

var (
	UI_URL string
)

type Service interface {
	AuthCallBack(res http.ResponseWriter, req *http.Request)
	HubAuthenticate(res http.ResponseWriter, req *http.Request)
}

// New returns the resource service implementation.
func New(api app.Config) Service {
	return &service{
		Service: api.Service("auth"),
		api:     api,
	}
}

func Status(res http.ResponseWriter, req *http.Request) {

	authSvc := Services{
		AuthService{
			Name:   "auth",
			Status: "ok",
		},
	}

	var log log.Logger
	res.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(res).Encode(authSvc); err != nil {
		log.Error(err)
	}
}

func Authenticate(res http.ResponseWriter, req *http.Request) {
	UI_URL = req.FormValue("redirect_uri")
	gothic.BeginAuthHandler(res, req)
}

func (s *service) AuthCallBack(res http.ResponseWriter, req *http.Request) {
	ghUser, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}

	r := request{
		db:            s.DB(context.Background()),
		log:           s.Logger(context.Background()),
		defaultScopes: s.api.Data().Default.Scopes,
		jwtConfig:     s.api.JWTConfig(),
	}

	params := req.URL.Query()

	if err = r.insertData(ghUser, params.Get("code")); err != nil {
		s.Logger(context.Background()).Error(err)
		res.Header().Set("Location", fmt.Sprintf("%s?status=%d", UI_URL, http.StatusBadRequest))
		res.WriteHeader(http.StatusTemporaryRedirect)
	}

	res.Header().Set("Location", fmt.Sprintf("%s?status=%d&code=%s", UI_URL, http.StatusOK, params.Get("code")))
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func (s *service) HubAuthenticate(res http.ResponseWriter, req *http.Request) {

	code := req.FormValue("code")

	r := request{
		db:            s.DB(context.Background()),
		log:           s.Logger(context.Background()),
		defaultScopes: s.api.Data().Default.Scopes,
		jwtConfig:     s.api.JWTConfig(),
	}

	var User model.User
	// Check if user exist
	q := r.db.Model(&model.User{}).
		Where("code = ?", code)

	err := q.First(&User).Error
	if err != nil {
		r.log.Error(err)
	}

	if err := r.db.Model(&model.User{}).Where("github_login = ?", User.GithubLogin).Update("code", "").Error; err != nil {
		r.log.Error(err)
	}

	// gets user scopes to add in jwt
	scopes, err := r.userScopes(&User)
	if err != nil {
		// return nil, err
		r.log.Error(err)
	}

	user, err := r.createTokens(&User, scopes)
	if err != nil {
		r.log.Error(err)
	}

	res.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(res).Encode(user); err != nil {
		r.log.Error(err)
	}
}

func List(res http.ResponseWriter, req *http.Request) {
	providers := []ProviderList{
		{
			Data: []Provider{
				{
					Name: "github",
				},
				{
					Name: "gitlab",
				},
				{
					Name: "bitbucket",
				},
			},
		},
	}

	var log log.Logger
	res.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(res).Encode(providers); err != nil {
		log.Error(err)
	}
}
