package hub

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ikawaha/goahttpcheck"
	"github.com/stretchr/testify/assert"

	"github.com/tektoncd/hub/api/gen/http/resource/server"
	resource "github.com/tektoncd/hub/api/gen/resource"
	resourceservice "github.com/tektoncd/hub/api/pkg/service"
)

var (
	resourceSvc resource.Service
)

func TestGetResources(t *testing.T) {
	LoadFixture(db, "../../fixtures")
	checker := goahttpcheck.New()
	checker.Mount(server.NewAllResourcesHandler, server.MountAllResourcesHandler, resource.NewAllResourcesEndpoint(resourceservice.NewResource(testConfig)))
	checker.Test(t, http.MethodGet, "/resources").
		Check().
		HasStatus(200).
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
			assert.Equal(t, len(jsonMap), 3)
		})

}
