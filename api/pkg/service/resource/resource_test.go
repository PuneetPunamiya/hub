package resource

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tektoncd/hub/api/gen/resource"
	"github.com/tektoncd/hub/api/pkg/testutils"
)

func TestVersionsByID(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.VersionsByIDPayload{ID: 6}
	res, err := resourceSvc.VersionsByID(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, res.Location, "/v1/resource/6/versions")
}

func TestByCatalogKindNameVersion(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ByCatalogKindNameVersionPayload{Catalog: "catalog-official", Kind: "task", Name: "tkn", Version: "0.1"}
	res, err := resourceSvc.ByCatalogKindNameVersion(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, res.Location, "/v1/resource/catalog-official/task/tkn/0.1")
}

func TestByCatalogKindNameVersionReadme(t *testing.T) {
	os.Setenv("CLONE_BASE_PATH", "testdata/catalog")
	defer os.Unsetenv("CLONE_BASE_PATH")

	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ByCatalogKindNameVersionReadmePayload{Catalog: "catalog-official", Kind: "task", Name: "tkn", Version: "0.1"}
	res, err := resourceSvc.ByCatalogKindNameVersionReadme(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, res.Location, "/v1/resource/catalog-official/task/tkn/0.1/readme")
}

func TestByCatalogKindNameVersionYaml(t *testing.T) {
	os.Setenv("CLONE_BASE_PATH", "testdata/catalog")
	defer os.Unsetenv("CLONE_BASE_PATH")

	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ByCatalogKindNameVersionYamlPayload{Catalog: "catalog-official", Kind: "task", Name: "tkn", Version: "0.1"}
	res, err := resourceSvc.ByCatalogKindNameVersionYaml(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, res.Location, "/v1/resource/catalog-official/task/tkn/0.1/yaml")
}

func TestByVersionID(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ByVersionIDPayload{VersionID: 6}
	res, err := resourceSvc.ByVersionID(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, res.Location, "/v1/resource/version/6")
}

func TestByCatalogKindName(t *testing.T) {
	tc := testutils.Setup(t)
	testutils.LoadFixtures(t, tc.FixturePath())

	resourceSvc := New(tc)
	payload := &resource.ByCatalogKindNamePayload{Catalog: "catalog-community", Kind: "task", Name: "img"}
	res, err := resourceSvc.ByCatalogKindName(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, res.Location, "/v1/resource/catalog-community/task/img")
}
