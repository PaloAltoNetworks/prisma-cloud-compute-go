package registry

import (
	"net/http"

	pcc "github.com/paloaltonetworks/prisma-cloud-compute-go"
)

const Endpoint = "api/v1/settings/registry"

type Registry struct {
	Specifications []Specification `json:"specifications,omitempty"`
}

type Specification struct {
	Cap                      int      `json:"cap,omitempty"`
	Collections              []string `json:"collections,omitempty"`
	CredentialId             string   `json:"credentialID,omitempty"`
	ExcludedRepositories     []string `json:"excludedRepositories,omitempty"`
	ExcludedTags             []string `json:"excludedTags,omitempty"`
	HarborDeploymentSecurity bool     `json:"harborDeploymentSecurity,omitempty"`
	JfrogRepoTypes           []string `json:"jfrogRepoTypes,omitempty"`
	Namespace                string   `json:"namespace,omitempty"`
	Os                       string   `json:"os,omitempty"`
	Tag                      string   `json:"tag,omitempty"`
	Registry                 string   `json:"registry,omitempty"`
	Repository               string   `json:"repository,omitempty"`
	Scanners                 int      `json:"scanners,omitempty"`
	Version                  string   `json:"version,omitempty"`
	VersionPattern           string   `json:"versionPattern,omitempty"`
}

// Return the registry scan settings.
func Get(c pcc.Client) (Registry, error) {
	var ans Registry
	if err := c.Communicate(http.MethodGet, Endpoint, nil, nil, &ans); err != nil {
		return ans, err
	}
	return ans, nil
}

// func Create(client pcc.Client, registry Registry) error {
// 	return CreateUpdate(client, registry, false)
// }

// Update the registry scan settings.
func Update(client pcc.Client, registry Registry) error {
	return CreateUpdate(client, registry, true)
}

func CreateUpdate(c pcc.Client, registry Registry, exists bool) error {
	var method string
	var data interface{}
	// if exists {
	method = http.MethodPut
	data = registry
	// } else {
	// 	method = http.MethodPost
	// 	data = registry.Specifications[0]
	// }
	err := c.Communicate(method, Endpoint, nil, data, nil)
	return err
}
