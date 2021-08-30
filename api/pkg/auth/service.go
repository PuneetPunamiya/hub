package auth

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/markbates/goth"
	"github.com/tektoncd/hub/api/pkg/db/model"
	"github.com/tektoncd/hub/api/pkg/token"
	"gorm.io/gorm"
)

func (r *request) userScopes(user *model.User) ([]string, error) {

	var userScopes []string = r.defaultScopes

	// db := svc.DB(context.Background())
	q := r.db.Preload("Scopes").Where(&model.User{GithubLogin: user.GithubLogin})

	dbUser := model.User{}
	if err := q.Find(&dbUser).Error; err != nil {
		r.log.Error(err)
		// return nil, internalError
		return nil, err
	}

	for _, s := range dbUser.Scopes {
		userScopes = append(userScopes, s.Name)
	}

	return userScopes, nil
}

func (r *request) createTokens(user *model.User, scopes []string) (*AuthenticateResult, error) {

	req := token.Request{
		User:      user,
		Scopes:    scopes,
		JWTConfig: r.jwtConfig,
	}

	accessToken, accessExpiresAt, err := req.AccessJWT()
	if err != nil {
		r.log.Error(err)
		return nil, err
		// return nil, internalError
	}

	refreshToken, refreshExpiresAt, err := req.RefreshJWT()
	if err != nil {
		r.log.Error(err)
		// return nil, internalError
	}

	user.RefreshTokenChecksum = createChecksum(refreshToken)

	if err = r.db.Save(user).Error; err != nil {
		r.log.Error(err)
		return nil, err
		// return nil, internalError
	}

	data := &AuthTokens{
		Access: &Token{
			Token:           accessToken,
			RefreshInterval: r.jwtConfig.AccessExpiresIn.String(),
			ExpiresAt:       accessExpiresAt,
		},
		Refresh: &Token{
			Token:           refreshToken,
			RefreshInterval: r.jwtConfig.AccessExpiresIn.String(),
			ExpiresAt:       refreshExpiresAt,
		},
	}

	return &AuthenticateResult{Data: data}, nil
}

func createChecksum(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func (r *request) insertData(ghUser goth.User, code string) error {

	// Check if user exist
	q := r.db.Model(&model.User{}).
		Where("github_login = ?", ghUser.NickName)

	var user model.User
	err := q.First(&user).Error
	if err != nil {
		// If user doesn't exist, create a new record
		if err == gorm.ErrRecordNotFound {

			user.GithubName = ghUser.Name
			user.GithubLogin = ghUser.NickName
			user.Type = model.NormalUserType
			user.AvatarURL = ghUser.AvatarURL
			user.Code = code

			err = r.db.Create(&user).Error
			if err != nil {
				r.log.Error(err)
				// return nil, internalError
				return err
			}
			// return &user, nil
		}
	} else {
		if err := r.db.Model(&model.User{}).Where("github_login = ?", ghUser.NickName).Update("code", code).Error; err != nil {
			r.log.Error(err)
			return err
		}
	}

	// User already exist, check if GitHub Name is empty
	// If Name is empty, then user is inserted through config.yaml
	// Update user with remaining details

	if user.GithubName == "" {
		user.GithubName = ghUser.Name
		user.Type = model.NormalUserType
	}

	// For existing user, check if URL is not added
	if user.AvatarURL == "" {
		user.AvatarURL = ghUser.AvatarURL
		if err = r.db.Save(&user).Error; err != nil {
			r.log.Error(err)
			return err
		}
	}

	return nil
}
