package main

import (
	"github.com/jinzhu/gorm"
	"github.com/tektoncd/hub/api/pkg/app"
	"github.com/tektoncd/hub/api/pkg/db/model"
	"go.uber.org/zap"
	"gopkg.in/gormigrate.v1"
)

// Migrate create tables and populates master tables
func Migrate(api *app.APIConfig) error {

	log := api.Logger()

	// NOTE: when writing a migration for a new table, add the same in InitSchema
	migration := gormigrate.New(
		api.DB(),
		gormigrate.DefaultOptions,
		[]*gormigrate.Migration{
			{
				ID: "202007141723",
				Migrate: func(tx *gorm.DB) error {

					tx.Model(&model.Catalog{}).AddUniqueIndex("idx_id", "id")

					catalog_query := `ALTER TABLE catalogs
						alter column name set NOT NULL,
						alter column type set NOT NULL,
						alter column url set NOT NULL,
						alter column owner set NOT NULL;`

					resource_query := `ALTER TABLE resources
						alter column name set NOT NULL,
						 alter column type set NOT NULL;`

					resource_version_query := `ALTER TABLE resource_versions
						 alter column version set NOT NULL,
						 alter column description set NOT NULL,
						 alter column min_pipelines_version set NOT NULL;`

					if err := tx.Exec(catalog_query).Error; err != nil {
						log.Error(err)
						return err
					}
					if err := tx.Exec(resource_query).Error; err != nil {
						log.Error(err)
						return err
					}
					if err := tx.Exec(resource_version_query).Error; err != nil {
						log.Error(err)
						return err
					}
					return nil
				},
			},
			{
				ID: "202006071000",
				Migrate: func(tx *gorm.DB) error {

					if err := tx.AutoMigrate(
						&model.Tag{}, &model.Catalog{},
						&model.Resource{}, &model.ResourceVersion{}).Error; err != nil {
						log.Error(err)
						return err
					}

					if err := fkey(log, tx, model.Resource{}, "catalog_id", "catalogs"); err != nil {
						return err
					}
					if err := fkey(log, tx, model.ResourceVersion{}, "resource_id", "resources"); err != nil {
						return err
					}
					if err := fkey(log, tx, model.ResourceTag{},
						"resource_id", "resources",
						"tag_id", "tags"); err != nil {
						return err
					}
					return nil
				},
			},
			{
				ID: "202006091100",
				Migrate: func(tx *gorm.DB) error {
					if err := tx.AutoMigrate(
						&model.ResourceVersion{}).Error; err != nil {
						log.Error(err)
						return err
					}
					return nil
				},
			},
		})

	migration.InitSchema(func(db *gorm.DB) error {
		if err := db.AutoMigrate(
			&model.Category{},
			&model.Tag{},
			&model.Catalog{},
			&model.Resource{},
			&model.ResourceVersion{},
		).Error; err != nil {
			log.Error(err)
			return err
		}

		db.Model(&model.Catalog{}).AddUniqueIndex("idx_id", "id")

		if err := fkey(log, db, model.Tag{}, "category_id", "categories"); err != nil {
			return err
		}

		if err := fkey(log, db, model.Resource{}, "catalog_id", "catalogs(id)"); err != nil {
			return err
		}

		if err := fkey(log, db, model.ResourceVersion{}, "resource_id", "resources"); err != nil {
			return err
		}

		if err := fkey(log, db, model.ResourceTag{},
			"resource_id", "resources",
			"tag_id", "tags"); err != nil {
			return err
		}

		log.Info("Schema initialised successfully !!")

		if err := initialiseTables(log, db); err != nil {
			log.Info("Data initialisation failed !!")
			return err
		}
		log.Info("Data added successfully !!")

		return nil
	})

	if err := migration.Migrate(); err != nil {
		log.Error(err, " failed to migrate")
		return err
	}

	log.Info("Migration ran successfully !!")
	return nil
}

func fkey(log *zap.SugaredLogger, db *gorm.DB, model interface{}, args ...string) error {
	for i := 0; i < len(args); i += 2 {
		col := args[i]
		table := args[i+1]
		err := db.Model(model).AddForeignKey(col, table, "CASCADE", "CASCADE").Error
		if err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}

// Adds data to category & tag table
func initialiseTables(log *zap.SugaredLogger, db *gorm.DB) error {
	var categories = map[string][]string{
		"Others":         []string{},
		"Build Tools":    []string{"build-tool"},
		"CLI":            []string{"cli"},
		"Cloud":          []string{"gcp", "aws", "azure", "cloud"},
		"Deploy":         []string{"deploy"},
		"Image Build":    []string{"image-build"},
		"Notification":   []string{"notification"},
		"Test Framework": []string{"test"},
	}

	for name, tags := range categories {
		cat := &model.Category{Name: name}
		if err := db.Create(cat).Error; err != nil {
			log.Error(err)
			return err
		}

		for _, tag := range tags {
			if err := db.Model(&cat).Association("Tags").
				Append(&model.Tag{Name: tag}).Error; err != nil {
				log.Error(err)
				return err
			}
		}
	}
	return nil
}
