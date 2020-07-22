package main

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
	"github.com/stretchr/testify/assert"
	"github.com/tektoncd/hub/api/pkg/db/model"
	"github.com/tektoncd/hub/api/pkg/testutils"
	"gopkg.in/gormigrate.v1"
)

var migrationsw = []*gormigrate.Migration{
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
	{
		ID: "202007141723",
		Migrate: func(tx *gorm.DB) error {

			tx.Model(&model.Catalog{}).AddUniqueIndex("idx_id", "id")

			catalog_query := `ALTER TABLE catalogs
				alter column name set NOT NULL,
				alter column type set NOT NULL,
				alter column url set NOT NULL;`

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
}

func tableCount(t *testing.T, db *gorm.DB, tableName string) (count int) {
	assert.NoError(t, db.Table(tableName).Count(&count).Error)
	return
}

func TestMigration(t *testing.T) {
	tc := testutils.Setup(t)

	err := tc.Error()

	m := gormigrate.New(tc.DB(), gormigrate.DefaultOptions, migrationsw)

	err = m.Migrate()

	assert.NoError(t, err)
	assert.True(t, tc.DB().HasTable(&model.Category{}))
}

func TestMigrateTo(t *testing.T) {
	tc := testutils.Setup(t)

	err := tc.Error()

	m := gormigrate.New(tc.DB(), gormigrate.DefaultOptions, migrationsw)
	err = m.MigrateTo("202007141723")

	assert.NoError(t, err)
	assert.True(t, tc.DB().HasTable(&model.Catalog{}))
	assert.True(t, tc.DB().HasTable(&model.Category{}))
	assert.True(t, tc.DB().HasTable(&model.Resource{}))
	assert.True(t, tc.DB().HasTable(&model.ResourceTag{}))
	assert.True(t, tc.DB().HasTable(&model.ResourceVersion{}))
	assert.True(t, tc.DB().HasTable(&model.Tag{}))

	assert.Equal(t, 0, tableCount(t, tc.DB(), "catalogs"))
}

func TestInitSchemaNoMigrations(t *testing.T) {
	tc := testutils.Setup(t)

	err := tc.Error()

	m := gormigrate.New(tc.DB(), gormigrate.DefaultOptions, []*gormigrate.Migration{})
	m.InitSchema(func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(&model.Category{},
			&model.Tag{},
			&model.Catalog{},
			&model.Resource{},
			&model.ResourceVersion{}).Error; err != nil {
			return err
		}
		return nil
	})

	assert.NoError(t, m.Migrate())
	assert.NoError(t, err)
	assert.True(t, tc.DB().HasTable(&model.Catalog{}))
	assert.True(t, tc.DB().HasTable(&model.Category{}))
	assert.True(t, tc.DB().HasTable(&model.Resource{}))
	assert.True(t, tc.DB().HasTable(&model.ResourceTag{}))
	assert.True(t, tc.DB().HasTable(&model.ResourceVersion{}))
	assert.True(t, tc.DB().HasTable(&model.Tag{}))
}

func TestInitSchemaWithMigrations(t *testing.T) {
	tc := testutils.Setup(t)

	err := tc.Error()

	m := gormigrate.New(tc.DB(), gormigrate.DefaultOptions, migrationsw)
	m.InitSchema(func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(&model.Category{},
			&model.Tag{},
			&model.Catalog{},
			&model.Resource{},
			&model.ResourceVersion{}).Error; err != nil {
			return err
		}
		return nil
	})

	assert.NoError(t, m.Migrate())
	assert.NoError(t, err)
	assert.True(t, tc.DB().HasTable(&model.Catalog{}))
	assert.True(t, tc.DB().HasTable(&model.Resource{}))
	assert.True(t, tc.DB().HasTable(&model.ResourceTag{}))
	assert.True(t, tc.DB().HasTable(&model.ResourceVersion{}))
}

func TestMissingID(t *testing.T) {
	tc := testutils.Setup(t)
	migrationsMissingID := []*gormigrate.Migration{
		{
			Migrate: func(tx *gorm.DB) error {
				return nil
			},
		},
	}

	m := gormigrate.New(tc.DB(), gormigrate.DefaultOptions, migrationsMissingID)
	assert.Equal(t, gormigrate.ErrMissingID, m.Migrate())
}
