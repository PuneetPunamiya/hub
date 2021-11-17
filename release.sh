#!/usr/bin/env bash

# Copyright © 2021 The Tekton Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

RELEASE_VERSION=""

declare -r SCRIPT_PATH=$(readlink -f "$0")
declare -r SCRIPT_DIR=$(cd $(dirname "$SCRIPT_PATH") && pwd)
declare -r API_DIR="$SCRIPT_DIR/api"
declare -r UI_DIR="$SCRIPT_DIR/ui"
declare -r RELEASES_DIR="$SCRIPT_DIR/releases"
declare -r UPSTREAM_REMOTE="origin"

declare -a BINARIES=(
  git
  hub
)

info() {
  echo "INFO: $@"
}

err() {
  echo "ERROR: $@"
}

getReleaseVersion() {
  [[ -z ${RELEASE_VERSION} ]] && {
    read -r -e -p "Enter a target release (i.e: v0.1.2): " RELEASE_VERSION
    [[ -z ${RELEASE_VERSION} ]] && {
      echo "no target release"
      exit 1
    }
  }
  [[ ${RELEASE_VERSION} =~ v[0-9]+\.[0-9]*\.[0-9]+ ]] || {
    echo "invalid version provided, need to match v\d+\.\d+\.\d+"
    exit 1
  }
}

buildDbMigrationImage() {
  info Building DB Migration Image
  echo -----------------------------------
  cd "$API_DIR"
  docker build -f db.Dockerfile -t docker.io/puneet2147/db-migration:${RELEASE_VERSION} . && docker push docker.io/puneet2147/db-migration:${RELEASE_VERSION}
  info DB Migration Image Build Successfully
  echo -----------------------------------
}

buildApiImage() {
  info Building API Image
  echo -----------------------------------
  cd "$API_DIR"
  docker build -t docker.io/puneet2147/api:${RELEASE_VERSION} . && docker push docker.io/puneet2147/api:${RELEASE_VERSION}
  info API Image Build Successfully
  echo -----------------------------------
}

buildUiImage() {
  info Building UI Image
  echo -----------------------------------
  cd "$UI_DIR"
  docker build -t docker.io/puneet2147/ui:${RELEASE_VERSION} . && docker push docker.io/puneet2147/ui:${RELEASE_VERSION}
  info UI Image Build Successfully
  echo -----------------------------------
}

db(){
	info Creating DB Release Yaml

  make db || {
    err 'db release build failed'
    return 1
  }
}

db-migration(){
	info Creating Db-Migration Release Yaml

  make db-migration || {
    err 'db-migration release build failed'
    return 1
  }
}

api-k8s(){
	info Creating API Release Yaml

  make api-k8s || {
    err 'api release build failed'
    return 1
  }
}

api-openshift(){
	info Creating API Release Yaml

  make api-openshift || {
    err 'api release build failed'
    return 1
  }
}

ui-k8s(){
	info Creating UI Release Yaml

  make ui-k8s || {
    err 'ui release build failed'
    return 1
  }
}

ui-openshift(){
	info Creating UI Release Yaml

  make ui-openshift || {
    err 'ui release build failed'
    return 1
  }
}

replaceImageName() {
  info Changing Image Name

  cd "$RELEASES_DIR"

  #  Replace the db-migration image name
  sed -i "s@image: quay.io/tekton-hub/db-migration@image: quay.io/tekton-hub/db-migration:$RELEASE_VERSION@g" db-migration.yaml

  # Replace the api image
  sed -i "s@image: quay.io/tekton-hub/api@image: quay.io/tekton-hub/api:$RELEASE_VERSION@g" api-k8s.yaml

  sed -i "s@image: quay.io/tekton-hub/api@image: quay.io/tekton-hub/api:$RELEASE_VERSION@g" api-openshift.yaml

  #Replace the ui image
  sed -i "s@image: quay.io/tekton-hub/ui@image: quay.io/tekton-hub/ui:$RELEASE_VERSION@g" ui-k8s.yaml

  sed -i "s@image: quay.io/tekton-hub/ui@image: quay.io/tekton-hub/ui:$RELEASE_VERSION@g" ui-openshift.yaml
}

createGitTag() {
  echo; echo 'Creating tag for new release:  '
#  read -r -e -p "Enter tag message: " TAG_MESSAGE
#  git tag -a "${RELEASE_VERSION}" -m "${TAG_MESSAGE}"
#  git push ${UPSTREAM_REMOTE} --tags

  hub release create -a  /home/Puneet/hub/releases/api-k8s.yaml -a /home/Puneet/hub/releases/api-openshift.yaml -m "Hub v1.6.0" v1.6.0
}

main() {
  # Ask the release version to build images
  getReleaseVersion

  # Generate the release yamls for db, db-migration, api and ui
  echo "********************************************"
  info     Generate the Release Yamls for Hub
  echo "********************************************"
  db
  db-migration
  api-k8s
  api-openshift
  ui-k8s
  ui-openshift

  # Build images for db-migration, api and ui
  echo "********************************************"
  info        Build the Images for Hub
  echo "********************************************"
  buildDbMigrationImage
  buildApiImage
  buildUiImage

  # Change the image name with the release version specified
  echo "********************************************"
  info      Replace the Images with New Version
  echo "********************************************"
  replaceImageName

  echo "********************************************"
  info            Create Git Tag
  echo "********************************************"
  createGitTag

  echo "********************************************"
  echo "***" Release Created for Hub successfully "***"
  echo "********************************************"
}

main $@