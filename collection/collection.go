package collection

import (
	"fmt"
	"net/http"

	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
)

const CollectionsEndpoint = "api/v1/collections"

type Collection struct {
	AccountIds  []string `json:"accountIDs,omitempty"`
	AppIds      []string `json:"appIDs,omitempty"`
	Clusters    []string `json:"clusters,omitempty"`
	CodeRepos   []string `json:"codeRepos,omitempty"`
	Color       string   `json:"color,omitempty"`
	Containers  []string `json:"containers,omitempty"`
	Description string   `json:"description,omitempty"`
	Functions   []string `json:"functions,omitempty"`
	Hosts       []string `json:"hosts,omitempty"`
	Images      []string `json:"images,omitempty"`
	Labels      []string `json:"labels,omitempty"`
	Name        string   `json:"name,omitempty"`
	Namespaces  []string `json:"namespaces,omitempty"`
}

// Get all collections.
func List(c pcc.Client) ([]Collection, error) {
	var ans []Collection
	if err := c.Request(http.MethodGet, CollectionsEndpoint, nil, nil, &ans); err != nil {
		return nil, fmt.Errorf("error listing collections: %s", err)
	}
	return ans, nil
}

// Get a specific collection.
func Get(c pcc.Client, name string) (*Collection, error) {
	collections, err := List(c)
	if err != nil {
		return nil, fmt.Errorf("error listing collections: %s", err)
	}
	for _, v := range collections {
		if v.Name == name {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("collection %s not found", name)
}

// Create a new collection.
func Create(c pcc.Client, collection Collection) error {
	return c.Request(http.MethodPost, CollectionsEndpoint, nil, collection, nil)
}

// Update an existing collection.
func Update(c pcc.Client, collection Collection) error {
	return c.Request(http.MethodPut, fmt.Sprintf("%s/%s", CollectionsEndpoint, collection.Name), nil, collection, nil)
}

// Delete an existing collection.
func Delete(c pcc.Client, name string) error {
	return c.Request(http.MethodDelete, fmt.Sprintf("%s/%s", CollectionsEndpoint, name), nil, nil, nil)
}
