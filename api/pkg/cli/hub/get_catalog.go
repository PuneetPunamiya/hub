package hub

import (
	"encoding/json"
	"net/http"

	cclient "github.com/tektoncd/hub/api/v1/gen/http/catalog/client"
)

type CatalogOptions struct {
}

type CatalogResult struct {
	data    []byte
	status  int
	err     error
	Catalog CatalogData
}
type CatalogData = cclient.ListResponseBody

func (c *client) GetAllCatalogs() CatalogResult {
	data, status, err := c.Get(catalogEndpoint())
	if status == http.StatusNotFound {
		err = nil
	}

	return CatalogResult{data: data, status: status, err: err}
}

// Typed returns unmarshalled API response as CatalogResponse
func (sr *CatalogResult) Type() (CatalogData, error) {
	if sr.Catalog.Data != nil || sr.err != nil {
		return sr.Catalog, sr.err
	}
	res := &CatalogData{}
	if sr.status == http.StatusNotFound {
		return sr.Catalog, sr.err
	}

	sr.err = json.Unmarshal(sr.data, res)
	sr.Catalog.Data = res.Data

	return sr.Catalog, sr.err
}

// Endpoint computes the endpoint url using input provided
func catalogEndpoint() string {
	return "/v1/catalogs"
}

func (h *client) GetCatalogsList() ([]string, error) {
	// Get all catalogs
	c := h.GetAllCatalogs()

	// Unmarshal the data
	var err error
	typed, err := c.Type()
	if err != nil {
		return nil, err
	}

	var data = struct {
		Catalogs CatalogData
	}{
		Catalogs: typed,
	}

	Catalog := data.Catalogs
	// // Get all catalog names
	var cat []string
	for i := range Catalog.Data {
		cat = append(cat, *Catalog.Data[i].Name)
	}

	return cat, nil
}
