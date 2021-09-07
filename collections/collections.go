package collections

import (
	"fmt"
	"net/http"

	pcc "github.com/paloaltonetworks/prisma-cloud-compute-go"
)

const endpoint = "api/v1/collections"

type Collection struct {
	AccountIDs  []string `json:"accountIDs,omitempty"`
	AppIDs      []string `json:"appIDs,omitempty"`
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

func List(client pcc.Client) ([]Collection, error) {
	var ans []Collection
	if err := client.Communicate(http.MethodGet, endpoint, nil, nil, &ans); err != nil {
		return ans, err
	}
	return ans, nil
}

func Get(client pcc.Client, name string) (*Collection, error) {
	collections, err := List(client)
	if err != nil {
		return nil, err
	}

	for _, v := range collections {
		if v.Name == name {
			return &v, err
		}
	}
	return nil, fmt.Errorf("Collection %s not found", name)
}

func Create(client pcc.Client, collection Collection) error {
	return createUpdate(client, collection, false)
}

func Update(client pcc.Client, collection Collection) error {
	return createUpdate(client, collection, true)
}

func Delete(client pcc.Client, name string) error {
	err := client.Communicate(http.MethodDelete, fmt.Sprintf("%s/%s", endpoint, name), nil, nil, nil)
	return err
}

func createUpdate(client pcc.Client, collection Collection, exists bool) error {
	var method, path string
	if exists {
		method = http.MethodPut
		path = fmt.Sprintf("%s/%s", endpoint, collection.Name)
	} else {
		method = http.MethodPost
		path = endpoint
	}
	err := client.Communicate(method, path, nil, collection, nil)
	return err
}
