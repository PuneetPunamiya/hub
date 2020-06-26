package hub

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	resource "github.com/tektoncd/hub/api/gen/resource"
)

var (
	resourceSvc resource.Service
)

func Test_AllResources(t *testing.T) {
	resourceSvc = NewResource(testConfig)
	LoadFixture(db, "../../fixtures")
	payload := &resource.AllResourcesPayload{Limit: 2}
	all, err := resourceSvc.AllResources(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, len(all), 2)
}
