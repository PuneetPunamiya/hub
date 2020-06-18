package hub

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/ikawaha/goahttpcheck"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	category "github.com/tektoncd/hub/api/gen/category"
	server "github.com/tektoncd/hub/api/gen/http/category/server"
	categoryservice "github.com/tektoncd/hub/api/pkg/service"
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

func TestGetCategories(t *testing.T) {
	checker := goahttpcheck.New()
	LoadFixture(db, "../fixtures")
	checker.Mount(server.NewAllHandler, server.MountAllHandler, category.NewAllEndpoint(categoryservice.NewCategory(testConfig)))

	checker.Test(t, http.MethodGet, "/categories").
		Check().
		HasStatus(http.StatusOK).
		Cb(func(r *http.Response) {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("unexpected error, %v", err)
			}
			defer r.Body.Close()
			var jsonMap []map[string]interface{}
			err = json.Unmarshal([]byte(b), &jsonMap)
			if err != nil {
				panic(err)
			}
			assert.Equal(t, 3, len(jsonMap))
			assert.Equal(t, "bird", jsonMap[0]["name"])
		})
}
