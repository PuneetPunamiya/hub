package model

import (
	"github.com/jinzhu/gorm"
)

type (
	Category struct {
		gorm.Model
		Name string `gorm:"not null"`
		Tags []Tag
	}

	Tag struct {
		gorm.Model
		Name       string `gorm:"not null"`
		Category   Category
		CategoryID int
		Resources  []*Resource `gorm:"many2many:resource_tags;"`
	}

	Catalog struct {
		gorm.Model
		Name       string `gorm:"primary_key"`
		Type       string `gorm:"primary_key"`
		URL        string `gorm:"not null"`
		ContextDir string
		Resources  []Resource
		Revision   string `gorm:"not null"`
	}

	Resource struct {
		gorm.Model
		Name      string `gorm:"not null"`
		Type      string `gorm:"not null"`
		Rating    float64
		Catalog   Catalog
		CatalogID uint
		Versions  []ResourceVersion
		Tags      []*Tag `gorm:"many2many:resource_tags;"`
	}

	ResourceVersion struct {
		gorm.Model
		Version             string `gorm:"not null"`
		Description         string
		URL                 string `gorm:"not null"`
		DisplayName         string
		MinPipelinesVersion string `gorm:"not null"`
		Resource            Resource
		ResourceID          uint
	}

	ResourceTag struct {
		ResourceID uint
		TagID      uint
	}
)
