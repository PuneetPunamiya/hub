package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/tektoncd/hub/api/gen/log"
	"github.com/tektoncd/hub/api/pkg/db/model"
	"gorm.io/gorm"
)

func addUsersDetailsInAccountTable(log *log.Logger) *gormigrate.Migration {

	return &gormigrate.Migration{
		ID: "202103111249_add_git_user",
		Migrate: func(db *gorm.DB) error {

			// Get all users from user table
			if err := db.Exec("CREATE TABLE user_prevs AS	SELECT * FROM users;").Error; err != nil {
				log.Error(err)
				return err
			}

			var users []model.UserPrev
			if err := db.Find(&users).Error; err != nil {
				log.Error(err)
				return err
			}

			// Update the user table based on the model
			db.Migrator().DropColumn(model.User{}, "github_login")

			db.Migrator().DropColumn(model.User{}, "github_name")

			db.Migrator().DropColumn(model.User{}, "avatar_url")

			// Create the account table
			if err := db.AutoMigrate(
				&model.Account{},
			); err != nil {
				log.Error(err)
				return err
			}

			var accounts []model.Account
			for _, user := range users {
				account := model.Account{
					UserID:   user.ID,
					Username: user.GithubLogin,
					Name:     user.GithubName,
				}
				accounts = append(accounts, account)
			}

			// Add user details in account table
			if err := db.Create(&accounts).Error; err != nil {
				log.Error(err)
				return err
			}

			return nil
		},
	}
}
