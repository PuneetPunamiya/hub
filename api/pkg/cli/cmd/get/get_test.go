// Copyright © 2020 The Tekton Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package get

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	res "github.com/tektoncd/hub/api/gen/resource"
	"github.com/tektoncd/hub/api/pkg/cli/test"
	cat "github.com/tektoncd/hub/api/v1/gen/catalog"
	goa "goa.design/goa/v3/pkg"
	"gopkg.in/h2non/gock.v1"
	"gotest.tools/v3/golden"
)

type InfoOptions struct {
	ResId      int
	Name       string
	Kind       string
	Catalog    string
	Version    string
	Resource   *res.ResourceData
	VersionArr *res.Versions
}

var catalogs = &cat.ListResult{
	Data: []*cat.Catalog{
		{
			ID:   1,
			Name: "tekton",
			Type: "community",
			URL:  "https://github.com/tektoncd/catalog",
		},
		{
			ID:   2,
			Name: "fruit",
			Type: "community",
			URL:  "https://github.com/fruit/catalog",
		},
	},
}

var pipelineRes = &res.ResourceData{
	ID:   1,
	Name: "apple",
	Kind: "Pipeline",
	Catalog: &res.Catalog{
		ID:   1,
		Name: "fruit",
		Type: "community",
	},
	Rating: 2.3,
	LatestVersion: &res.ResourceVersionData{
		ID:                  03,
		Version:             "0.1",
		Description:         "v0.1 of Pipeline apple",
		DisplayName:         "Red apple",
		MinPipelinesVersion: "0.17.1",
		RawURL:              "http://raw.github.url/apple/0.1/apple.yaml",
		WebURL:              "http://web.github.com/apple/0.1/apple.yaml",
		UpdatedAt:           "2020-01-01 12:00:00 +0000 UTC",
	},
	Tags: []*res.Tag{
		{
			ID:   3,
			Name: "fruit",
		},
	},
	Versions: []*res.ResourceVersionData{
		{
			ID:      11,
			Version: "0.1",
		},
	},
}

var pipelineResWithLatestVersion = &res.ResourceVersionData{
	ID:                  11,
	Version:             "0.3",
	DisplayName:         "mango",
	Description:         "v0.3 of Pipeline mango",
	MinPipelinesVersion: "0.12",
	RawURL:              "http://raw.github.url/mango/0.3/mango.yaml",
	WebURL:              "http://web.github.com/mango/0.3/mango.yaml",
	UpdatedAt:           "2020-01-01 12:00:00 +0000 UTC",
	Resource: &res.ResourceData{
		ID:   7,
		Name: "foo",
		Kind: "Pipeline",
		Catalog: &res.Catalog{
			ID:   1,
			Name: "fruit",
			Type: "community",
		},
		Rating: 4.3,
		Tags: []*res.Tag{
			{
				ID:   3,
				Name: "fruit",
			},
		},
	},
}

var pipelineResWithOldVersion = &res.ResourceVersionData{
	ID:                  11,
	Version:             "0.2",
	DisplayName:         "mango",
	Description:         "v0.3 of Pipeline mango",
	MinPipelinesVersion: "0.12",
	RawURL:              "http://raw.github.url/mango/0.2/mango.yaml",
	WebURL:              "http://web.github.com/mango/0.2/mango.yaml",
	UpdatedAt:           "2020-01-01 12:00:00 +0000 UTC",
	Resource: &res.ResourceData{
		ID:   7,
		Name: "foo",
		Kind: "Pipeline",
		Catalog: &res.Catalog{
			ID:   1,
			Name: "fruit",
			Type: "community",
		},
		Rating: 4.3,
		Tags: []*res.Tag{
			{
				ID:   3,
				Name: "fruit",
			},
		},
	},
}

var ver = &res.Versions{
	Latest: &res.ResourceVersionData{
		ID:      11,
		Version: "0.3",
		RawURL:  "http://raw.github.url/mango/0.3/mango.yaml",
		WebURL:  "http://web.github.com/mango/0.3/mango.yaml",
	},
	Versions: []*res.ResourceVersionData{
		{
			ID:      11,
			Version: "0.3",
			RawURL:  "http://raw.github.url/mango/0.3/mango.yaml",
			WebURL:  "http://web.github.com/mango/0.3/mango.yaml",
		},
		{
			ID:      10,
			Version: "0.2",
			RawURL:  "http://raw.github.url/mango/0.2/mango.yaml",
			WebURL:  "http://web.github.com/mango/0.2/mango.yaml",
		},
	},
}

var want string = `
Get a Abc of name 'foo':

    tkn hub get abc foo

or

Get a Abc of name 'foo' of version '0.3':

    tkn hub get abc foo --version 0.3
`

func mockApi(io InfoOptions, resourceWithVersion *res.ResourceVersionData) {

	// Get all catalogs
	gock.New(test.API).
		Get("/v1/catalogs").
		Reply(200).
		JSON(&catalogs)

	// Get all resources with catalog as Fruit
	rArr := &res.Resources{Data: res.ResourceDataCollection{io.Resource}}
	resources := res.NewViewedResources(rArr, "withoutVersion")
	gock.New(test.API).
		Get("/query").
		MatchParam("catalogs", io.Catalog).
		Reply(200).
		JSON(&resources.Projected)

	// Get ResourceId in order to get all versions of resource
	rVer := &res.ResourceVersion{Data: resourceWithVersion}
	resWithVersion := res.NewViewedResourceVersion(rVer, "default")

	resInfo := fmt.Sprintf("%s/%s/%s", io.Catalog, io.Kind, io.Name)
	gock.New(test.API).
		Get("/resource/" + resInfo).
		Reply(200).
		JSON(&resWithVersion.Projected)

	// Get all versios of the resource
	v := &res.ResourceVersions{Data: io.VersionArr}
	version := res.NewViewedResourceVersions(v, "min")

	vArr := fmt.Sprintf("%d/versions", io.ResId)
	gock.New(test.API).
		Get("/resource/" + vArr).
		Reply(200).
		JSON(&version.Projected)

	gock.New(test.API).
		Get("/resource/" + resInfo + "/" + io.Version).
		Reply(200).
		JSON(&resWithVersion.Projected)
}

func TestValidate(t *testing.T) {
	opt := options{
		version: "0.1",
	}
	err := opt.validate()
	assert.NoError(t, err)

	opt = options{
		version: "0.3.1",
	}
	err = opt.validate()
	assert.NoError(t, err)
}

func TestValidate_ErrorCase(t *testing.T) {
	opt := options{
		version: "abc",
	}
	err := opt.validate()
	assert.EqualError(t, err, "invalid value \"abc\" set for option version. valid eg. 0.1, 1.2.1")
}

func TestGetResource_WithNewVersion(t *testing.T) {
	cli := test.NewCLI()

	defer gock.Off()

	mockApi(InfoOptions{
		ResId:      12,
		Name:       "mango",
		Kind:       "pipeline",
		Catalog:    "fruit",
		Version:    "0.3",
		Resource:   pipelineRes,
		VersionArr: ver,
	}, pipelineResWithLatestVersion)

	gock.New("http://raw.github.url").
		Get("/mango/0.3/mango.yaml").
		Reply(200).
		File("./testdata/pipeline-mango-v0.3.yaml")

	buf := new(bytes.Buffer)
	cli.SetStream(buf, buf)

	opt := &options{
		cli:     cli,
		kind:    "pipeline",
		args:    []string{"mango"},
		from:    "fruit",
		version: "0.3",
	}

	err := opt.run()
	assert.NoError(t, err)
	golden.Assert(t, buf.String(), fmt.Sprintf("%s.golden", t.Name()))
	assert.Equal(t, gock.IsDone(), false)
}

func TestGetResource_WithOldVersion(t *testing.T) {
	cli := test.NewCLI()

	defer gock.Off()

	mockApi(InfoOptions{
		ResId:      12,
		Name:       "mango",
		Kind:       "pipeline",
		Catalog:    "fruit",
		Version:    "0.2",
		Resource:   pipelineRes,
		VersionArr: ver,
	}, pipelineResWithOldVersion)

	gock.New("http://raw.github.url").
		Get("/mango/0.2/mango.yaml").
		Reply(200).
		File("./testdata/pipeline-mango-v0.2.yaml")

	buf := new(bytes.Buffer)
	cli.SetStream(buf, buf)

	opt := &options{
		cli:     cli,
		kind:    "pipeline",
		args:    []string{"mango"},
		from:    "fruit",
		version: "0.2",
	}

	err := opt.run()
	assert.NoError(t, err)
	golden.Assert(t, buf.String(), fmt.Sprintf("%s.golden", t.Name()))
	assert.Equal(t, gock.IsDone(), false)
}

func TestGet_ResourceNotFound(t *testing.T) {
	cli := test.NewCLI()

	defer gock.Off()

	gock.New(test.API).
		Get("/resource/tekton/pipeline/xyz").
		Reply(404).
		JSON(&goa.ServiceError{
			ID:      "123456",
			Name:    "not-found",
			Message: "resource not found",
		})

	buf := new(bytes.Buffer)
	cli.SetStream(buf, buf)

	opt := &options{
		cli:  cli,
		kind: "pipeline",
		args: []string{"xyz"},
		from: "tekton",
	}

	err := opt.run()
	assert.Error(t, err)
	assert.EqualError(t, err, "No Resource Found")
	assert.Equal(t, gock.IsDone(), true)
}

func Test_examples(t *testing.T) {
	got := examples("abc")
	assert.Equal(t, want, got)
}
