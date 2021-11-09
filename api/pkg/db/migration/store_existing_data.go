package migration

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/tektoncd/hub/api/gen/log"
	"github.com/tektoncd/hub/api/pkg/db/model"
	"gorm.io/gorm"
)

func addGitUsers(log *log.Logger) *gormigrate.Migration {

	return &gormigrate.Migration{
		ID: "202103111249_add_git_user",
		Migrate: func(db *gorm.DB) error {
			// if err := db.AutoMigrate(&model.Config{}); err != nil {
			// 	log.Error(err)
			// 	return err
			// }

			// Get all users from user table
			var user []model.User
			if err := db.Find(&user).Error; err != nil {
				fmt.Println("Bhai chal jaaa")
				log.Error(err)
				return err
			}

			fmt.Println("---------------------------------")
			fmt.Println(user)
			fmt.Println("---------------------------------")

			// if err := db.Migrator().DropTable(
			// 	&model.Account{},
			// 	&model.GitUser{},
			// ); err != nil {
			// 	log.Error(err)
			// 	return err
			// }

			if err := db.AutoMigrate(
				&model.User{},
				&model.Account{},
			); err != nil {
				log.Error(err)
				return err
			}

			// var users_last_value uint
			// if err := db.Table("users_id_seq").Pluck("last_value", &users_last_value).Error; err != nil {
			// 	log.Error(err)
			// 	return err
			// }

			// gitUser := make([]model.GitUser, 0, 8)
			// account := make([]model.Account, 0, 7)

			for i := range user {

				fmt.Println("--------------works till here------------------------------")
				fmt.Println(i, user[i].ID)
				fmt.Println("--------------works till here------------------------------")
				// fmt.Println(gitUser[i])
				// fmt.Println("--------------works till here------------------------------")
				// fmt.Println(len(gitUser))

				// gitUser[i].ID = user[i].ID
				// gitUser[i].RefreshTokenChecksum = user[i].RefreshTokenChecksum

				// account[i].GitUserID = user[i].ID
				// account[i].Username = user[i].GithubLogin
				// account[i].Name = user[i].GithubName
			}

			// if err := db.Create(
			// 	&gitUser,
			// ).Error; err != nil {
			// 	log.Error(err)
			// 	return err
			// }

			// if err := db.Create(
			// 	&account,
			// ).Error; err != nil {
			// 	log.Error(err)
			// 	return err
			// }

			return nil
		},
	}
}
