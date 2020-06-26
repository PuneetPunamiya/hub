package hub

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"

	resource "github.com/tektoncd/hub/api/gen/resource"
	app "github.com/tektoncd/hub/api/pkg/app"
	"github.com/tektoncd/hub/api/pkg/db/model"
)

// resource service example implementation.
// The example methods log the requests and return zero values.
type resourcesrvc struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
}

// NewResource returns the resource service implementation.
func NewResource(api *app.ApiConfig) resource.Service {
	return &resourcesrvc{api.Logger(), api.DB()}
}

// Get all Resources
func (s *resourcesrvc) AllResources(ctx context.Context, p *resource.AllResourcesPayload) (res resource.ResourceCollection, err error) {

	var all []model.Resource

	if err := s.db.Order("rating desc, name").Limit(p.Limit).
		Preload("Catalog").
		Preload("Versions", func(db *gorm.DB) *gorm.DB {
			return db.Order("resource_versions.id ASC")
		}).
		Preload("Tags").
		Find(&all).Error; err != nil {
		return []*resource.Resource{}, errors.New("Failed to fetch Resources")
	}

	for _, r := range all {
		tags := []*resource.Tag{}
		for _, t := range r.Tags {
			tags = append(tags, &resource.Tag{
				ID:   t.ID,
				Name: t.Name,
			})
		}
		latestVersion := r.Versions[len(r.Versions)-1]
		res = append(res, &resource.Resource{
			ID:            r.ID,
			Name:          r.Name,
			DisplayName:   latestVersion.DisplayName,
			Tags:          tags,
			Catalog:       &resource.Catalog{ID: r.Catalog.ID, Type: r.Catalog.Type},
			Type:          r.Type,
			LatestVersion: latestVersion.Version,
			Description:   latestVersion.Description,
			Rating:        uint(r.Rating),
			LastUpdatedAt: r.UpdatedAt.String(),
		})
	}

	return res, nil
}
