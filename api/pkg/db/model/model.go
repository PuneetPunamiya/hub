package model

import (
	"github.com/jinzhu/gorm"
)

type (
	Category struct {
		gorm.Model
		Name string `gorm:"size:100;not null;unique"`
		Tags []Tag
	}

	Tag struct {
		gorm.Model
		Name       string `gorm:"size:100;not null;unique"`
		Category   Category
		CategoryID int
		Resources  []*Resource `gorm:"many2many:resource_tags;"`
	}

	Catalog struct {
		gorm.Model
		Name       string `gorm:"size:100;not null;unique"`
		Type       string `gorm:"size:100;not null;unique"`
		URL        string `gorm:"size:100;not null;unique"`
		Owner      string
		ContextDir string
		Resources  []Resource
		Revision   string
	}

	Resource struct {
		gorm.Model
		Name      string `gorm:"size:100;not null;unique"`
		Type      string `gorm:"size:100;not null;unique"`
		Rating    float64
		Catalog   Catalog
		CatalogID uint
		Versions  []ResourceVersion
		Tags      []*Tag `gorm:"many2many:resource_tags;"`
	}

	ResourceVersion struct {
		gorm.Model
		Version             string `gorm:"size:100;not null;unique"`
		Description         string
		URL                 string `gorm:"size:100;not null;unique"`
		DisplayName         string
		MinPipelinesVersion string `gorm:"size:100;not null;unique"`
		Resource            Resource
		ResourceID          uint
	}

	ResourceTag struct {
		ResourceID uint
		TagID      uint
	}
)
