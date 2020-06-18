package hub

import (
	"context"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	category "github.com/tektoncd/hub/api/gen/category"
)

var (
	categorySvc category.Service
)

// LoadFixture ...
func LoadFixture(db *gorm.DB, fixtureDir string) error {
	fixtures, err := testfixtures.New(
		testfixtures.Database(db.DB()),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(fixtureDir),
	)
	if err != nil {
		return err
	}
	if err := fixtures.Load(); err != nil {
		return err
	}
	return nil
}

func Test_All(t *testing.T) {
	categorySvc = NewCategory(testConfig)
	LoadFixture(db, "../../fixtures")
	all, err := categorySvc.All(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 3, len(all))
	assert.Equal(t, 2, len(all[0].Tags))
	assert.Equal(t, "bird", all[0].Name) // categories are sorted by name
}
