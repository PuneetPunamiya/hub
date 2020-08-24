package admin

import (
	"context"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"goa.design/goa/v3/security"

	"github.com/tektoncd/hub/api/gen/admin"
	"github.com/tektoncd/hub/api/pkg/app"
	"github.com/tektoncd/hub/api/pkg/db/model"
	"github.com/tektoncd/hub/api/pkg/token"
)

type service struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
	jwtKey string
}

var (
	invalidTokenError  = admin.MakeInvalidToken(fmt.Errorf("invalid user token"))
	invalidScopesError = admin.MakeInvalidScopes(fmt.Errorf("user not authorized"))
	internalError      = admin.MakeInternalError(fmt.Errorf("failed to create agent"))
)

// New returns the admin service implementation.
func New(api app.Config) admin.Service {
	return &service{api.Logger(), api.DB(), api.JWTSigningKey()}
}

// JWTAuth implements the authorization logic for service "admin" for the "jwt"
// security scheme.
func (s *service) JWTAuth(ctx context.Context, jwt string, scheme *security.JWTScheme) (context.Context, error) {

	claims, err := token.Verify(jwt, s.jwtKey)
	if err != nil {
		return ctx, invalidTokenError
	}

	err = token.ValidateScopes(claims, scheme)
	if err != nil {
		return ctx, invalidScopesError
	}

	return ctx, nil
}

// Create a agent user with required scopes
func (s *service) CreateAgent(ctx context.Context, p *admin.CreateAgentPayload) (res *admin.CreateAgentResult, err error) {

	tx := s.db.Begin()

	// Check if a normal user exists with the name
	user := &model.User{}
	q := tx.Model(&model.User{}).
		Where("LOWER(github_name) = ?", strings.ToLower(p.Name))

	if err = q.Find(&user).Error; err == nil {
		return nil, admin.MakeInvalidPayload(fmt.Errorf("normal user exists with name: %s", p.Name))
	}

	// Check if an agent exist with the name
	q = tx.Model(&model.User{}).Preload("Scopes").
		Where("LOWER(name) = ? AND LOWER(type) = ?", strings.ToLower(p.Name), model.Agent)

	if err := q.First(&user).Error; err != nil {

		// If agent does not exist then create one
		if gorm.IsRecordNotFoundError(err) {
			return s.addNewAgent(tx, p.Name, p.Scopes)
		}
		s.logger.Error(err)
		return nil, internalError
	}

	// If an agent with name already exists, then update the scopes of agent
	return s.updateAgent(tx, user, p.Scopes)
}

func (s *service) addNewAgent(tx *gorm.DB, name string, scopes []string) (*admin.CreateAgentResult, error) {

	user := &model.User{
		Name: name,
		Type: model.Agent,
	}
	if err := tx.Create(user).Error; err != nil {
		s.logger.Error(err)
		return nil, internalError
	}

	if err := s.addScopesForUser(tx, user, scopes); err != nil {
		return nil, err
	}

	token, err := s.createJWT(user, scopes)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		s.logger.Error(err)
		return nil, internalError
	}

	return &admin.CreateAgentResult{Token: token}, nil
}

func (s *service) updateAgent(tx *gorm.DB, user *model.User, scopes []string) (*admin.CreateAgentResult, error) {

	// Delete existing scopes of agent
	if err := tx.Where("user_id = ?", user.ID).Delete(&model.UserScope{}).Error; err != nil {
		s.logger.Error(err)
		return nil, internalError
	}

	// Add new scopes for agent
	if err := s.addScopesForUser(tx, user, scopes); err != nil {
		return nil, err
	}

	token, err := s.createJWT(user, scopes)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		s.logger.Error(err)
		return nil, internalError
	}

	return &admin.CreateAgentResult{Token: token}, nil
}

func (s *service) addScopesForUser(tx *gorm.DB, user *model.User, scopes []string) error {

	for _, sc := range scopes {

		scope := &model.Scope{}
		if err := tx.Where("name = ?", strings.ToLower(sc)).
			First(&scope).Error; err != nil {

			// If scope in payload does not exist then return
			if gorm.IsRecordNotFoundError(err) {
				tx.Rollback()
				return admin.MakeInvalidPayload(fmt.Errorf("invalid scope: %s", sc))
			}
			s.logger.Error(err)
			return internalError
		}

		us := model.UserScope{UserID: user.ID, ScopeID: scope.ID}
		if err := tx.Create(us).Error; err != nil {
			s.logger.Error(err)
			return internalError
		}
	}

	return nil
}

func (s *service) createJWT(user *model.User, scopes []string) (string, error) {

	claim := jwt.MapClaims{
		"id":     user.ID,
		"name":   user.Name,
		"scopes": scopes,
	}

	token, err := token.Create(claim, s.jwtKey)
	if err != nil {
		s.logger.Error(err)
		return "", internalError
	}

	return token, nil
}
