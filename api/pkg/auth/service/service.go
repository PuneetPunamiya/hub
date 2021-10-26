// Copyright © 2021 The Tekton Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/markbates/goth"
	"github.com/tektoncd/hub/api/pkg/auth/app"
	"github.com/tektoncd/hub/api/pkg/db/model"
	"github.com/tektoncd/hub/api/pkg/token"
	"gorm.io/gorm"
)

func (r *request) userScopes(user *model.Account) ([]string, error) {

	var userScopes []string = r.defaultScopes

	// q := r.db.Preload("Scopes").Where(&model.User{GithubLogin: user.GithubLogin})
	q := r.db.Preload("Scopes").Where(&model.Account{Username: user.Username})

	dbUser := model.GitUser{}
	if err := q.Model(&dbUser).Joins("JOINS accounts ON accounts.git_user_id = git_users.id").Where("accounts.username = ?", user.Username).Error; err != nil {
		r.log.Error(err)
		return nil, err
	}

	for _, s := range dbUser.Scopes {
		userScopes = append(userScopes, s.Name)
	}

	return userScopes, nil
}

func (r *request) createTokens(user *model.GitUser, scopes []string) (*app.AuthenticateResult, error) {

	req := token.Request{
		GitUser:   user,
		Scopes:    scopes,
		JWTConfig: r.jwtConfig,
	}

	accessToken, accessExpiresAt, err := req.AccessJWT()
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	refreshToken, refreshExpiresAt, err := req.RefreshJWT()
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	user.RefreshTokenChecksum = createChecksum(refreshToken)

	if err = r.db.Save(user).Error; err != nil {
		r.log.Error(err)
		return nil, err
	}

	data := &app.AuthTokens{
		Access: &app.Token{
			Token:           accessToken,
			RefreshInterval: r.jwtConfig.AccessExpiresIn.String(),
			ExpiresAt:       accessExpiresAt,
		},
		Refresh: &app.Token{
			Token:           refreshToken,
			RefreshInterval: r.jwtConfig.RefreshExpiresIn.String(),
			ExpiresAt:       refreshExpiresAt,
		},
	}

	return &app.AuthenticateResult{Data: data}, nil
}

func createChecksum(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func (r *request) insertData(ghUser goth.User, code, provider string) error {

	q := r.db.Model(&model.GitUser{}).
		Where("email = ?", ghUser.Email)

	// Check if user exist

	var acc model.Account
	var user model.GitUser
	err := q.First(&user).Error
	if err != nil {

		s := r.db.Model(&model.Account{}).Where("username = ?", ghUser.NickName).Where("provider = ?", provider)
		err = s.First(&acc).Error

		// If user doesn't exist, create a new record
		if err == gorm.ErrRecordNotFound {

			acc.Name = ghUser.Name
			acc.Username = ghUser.NickName
			acc.Type = model.NormalUserType
			acc.AvatarURL = ghUser.AvatarURL
			acc.Provider = provider

			user.Code = code
			user.Email = ghUser.Email

			result := map[string]interface{}{}
			r.db.Model(&model.GitUser{}).Last(&result)
			user.ID = result["id"].(uint) + 1

			err = r.db.Create(&user).Error
			if err != nil {
				r.log.Error(err)
				return err
			}

			acc.GitUserID = user.ID
			if err = r.db.Create(&acc).Error; err != nil {
				r.log.Error(err)
				return err
			}
		} else {
			if err := r.db.Model(&model.GitUser{}).Where("id = ?", acc.GitUserID).Updates(model.GitUser{Code: code, Email: ghUser.Email}).Error; err != nil {
				r.log.Error(err)
				return err
			}

			if err := r.db.Model(&model.Account{}).Where("id = ?", acc.GitUserID).Updates(model.Account{AvatarURL: ghUser.AvatarURL}).Error; err != nil {
				r.log.Error(err)
				return err
			}
		}
	} else {

		result := map[string]interface{}{}
		// r.db.Model(&model.GitUser{}).First(&result).Where("email = ?", ghUser.Email)

		// This checks is when user logins with different git provider
		q := r.db.Model(&model.Account{}).Where("email = ?", ghUser.Email)
		err = q.First(&result).Error
		if err == nil {
			user.ID = result["id"].(uint)
		}

		s := r.db.Model(&model.Account{}).Where("git_user_id = ?", user.ID).Where("provider = ?", provider)
		err = s.First(&acc).Error

		if err == gorm.ErrRecordNotFound {
			acc.Name = ghUser.Name
			acc.Username = ghUser.NickName
			acc.Type = model.NormalUserType
			acc.AvatarURL = ghUser.AvatarURL
			acc.Provider = provider
			acc.GitUserID = user.ID

			if err = r.db.Create(&acc).Error; err != nil {
				r.log.Error(err)
				return err
			}
		}

		if err := r.db.Model(&model.GitUser{}).Where("email = ?", ghUser.Email).Update("code", code).Error; err != nil {
			r.log.Error(err)
			return err
		}
	}

	// User already exist, check if GitHub Name is empty
	// If Name is empty, then user is inserted through config.yaml
	// Update user with remaining details

	// if user.GithubName == "" {
	// 	user.GithubName = ghUser.Name
	// 	user.Type = model.NormalUserType
	// }

	// if acc.Username == "" {
	// 	acc.Username = ghUser.Name
	// 	acc.Type = model.NormalUserType
	// }

	// // For existing user, check if URL is not added
	// if user.AvatarURL == "" {
	// 	user.AvatarURL = ghUser.AvatarURL
	// 	if err = r.db.Save(&user).Error; err != nil {
	// 		r.log.Error(err)
	// 		return err
	// 	}
	// }

	// if acc.AvatarURL == "" {
	// 	acc.AvatarURL = ghUser.AvatarURL
	// 	if err = r.db.Save(&user).Error; err != nil {
	// 		r.log.Error(err)
	// 		return err
	// 	}
	// }

	return nil
}
