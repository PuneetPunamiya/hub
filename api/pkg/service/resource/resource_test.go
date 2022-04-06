// Copyright © 2022 The Tekton Authors.
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

package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreationRawURLBitbucket(t *testing.T) {
	url := "https://bitbucket.org/org/catalog/src/main/task/name/0.1/name.yaml"
	replacer := getStringReplacer(url, "bitbucket")
	rawUrl := replacer.Replace(url)
	expected := "https://bitbucket.org/org/catalog/raw/main/task/name/0.1/name.yaml"
	assert.Equal(t, expected, rawUrl)
}

func TestCreationRawURLGitlab(t *testing.T) {
	url := "https://gitlab.com/org/catalog/-/blob/main/task/name/0.1/name.yaml"
	replacer := getStringReplacer(url, "gitlab")
	rawUrl := replacer.Replace(url)
	expected := "https://gitlab.com/org/catalog/-/raw/main/task/name/0.1/name.yaml"
	assert.Equal(t, expected, rawUrl)
}

func TestCreationRawURLGitlabEnterprise(t *testing.T) {
	url := "https://gitlab.myhost.com/org/catalog/-/blob/main/task/name/0.1/name.yaml"
	replacer := getStringReplacer(url, "gitlab")
	rawUrl := replacer.Replace(url)
	expected := "https://gitlab.myhost.com/org/catalog/-/raw/main/task/name/0.1/name.yaml"
	assert.Equal(t, expected, rawUrl)
}
