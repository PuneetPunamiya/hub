package hub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCatalogEndpoint(t *testing.T) {
	url := catalogEndpoint()
	assert.Equal(t, "/v1/catalogs", url)

}
