package initializer

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"

	"github.com/tektoncd/hub/api/pkg/app"
	"github.com/tektoncd/hub/api/pkg/db/model"
)

// Initializer defines the configuration required for initailizer
// to populate the tables
type Initializer struct {
	log  *zap.SugaredLogger
	db   *gorm.DB
	data *app.Data
}

// New returns the Initializer implementation.
func New(api app.BaseConfig) *Initializer {
	return &Initializer{
		log:  api.Logger().With("component", "initiailizer"),
		db:   api.DB(),
		data: api.Data(),
	}
}

// Run executes the func which populate the tables
func (i *Initializer) Run() error {

	if err := i.addCategories(); err != nil {
		return err
	}
	if err := i.addCatalogs(); err != nil {
		return err
	}
	if err := i.addUsers(); err != nil {
		return err
	}
	return nil
}

func (i *Initializer) addCategories() error {

	db := i.db

	// Checks if tables exists
	if !db.HasTable(&model.Category{}) || !db.HasTable(model.Tag{}) {
		return fmt.Errorf("categories or tags table not found")
	}

	for _, c := range i.data.Categories {
		cat := &model.Category{Name: c.Name}
		if err := i.db.Where(cat).FirstOrCreate(cat).
			Error; err != nil {
			i.log.Error(err)
			return err
		}
		for _, t := range c.Tags {
			tag := &model.Tag{Name: t, CategoryID: cat.ID}
			if err := db.Where(tag).FirstOrCreate(tag).
				Error; err != nil {
				i.log.Error(err)
				return err
			}
		}
	}
	return nil
}

func (i *Initializer) addCatalogs() error {

	db := i.db

	// Checks if tables exists
	if !db.HasTable(&model.Catalog{}) {
		return fmt.Errorf("catalogs table not found")
	}

	for _, c := range i.data.Catalogs {
		cat := &model.Catalog{
			Name:       c.Name,
			Org:        c.Org,
			Type:       c.Type,
			URL:        c.URL,
			Revision:   c.Revision,
			ContextDir: c.ContextDir,
		}
		if err := db.Where(&model.Catalog{Name: c.Name, Org: c.Org}).
			FirstOrCreate(cat).
			Error; err != nil {
			i.log.Error(err)
			return err
		}
	}
	return nil
}

func (i *Initializer) addUsers() error {

	db := i.db
	// Checks if tables exists
	if !db.HasTable(&model.User{}) || !db.HasTable(model.Scope{}) {
		return fmt.Errorf("user or scope table not found")
	}

	for _, u := range i.data.Users {

		// If user does not exists then skip and move to next one
		user := &model.User{}
		if err := db.Where("LOWER(github_login) = ?", strings.ToLower(u.Name)).
			First(&user).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				i.log.Errorf("user %s not found: %s", u.Name, err)
				continue
			}
			i.log.Error(err)
			return err
		}
		for _, s := range u.Scopes {

			// If scope does not exists then return
			scope := &model.Scope{}
			if err := db.Where("name = ?", strings.ToLower(s)).
				First(&scope).Error; err != nil {
				if gorm.IsRecordNotFoundError(err) {
					i.log.Errorf("scope (%s) does not exist: %s", s, err)
					return fmt.Errorf("invalid-scope")
				}
				i.log.Error(err)
				return err
			}

			us := model.UserScope{UserID: user.ID, ScopeID: scope.ID}
			if err := db.Model(&model.UserScope{}).Where(&us).
				FirstOrCreate(&us).Error; err != nil {
				i.log.Error(err)
				return err
			}
		}
	}
	return nil
}
