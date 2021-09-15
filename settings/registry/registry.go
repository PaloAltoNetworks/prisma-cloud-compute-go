package registry

import (
	"net/http"

	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
)

const Endpoint = "api/v1/settings/registry"

type Registry struct {
	Specifications []Specification `json:"specifications,omitempty"`
}

type Specification struct {
	Cap                      int      `json:"cap,omitempty"`
	Collections              []string `json:"collections,omitempty"`
	Credential               string   `json:"credentialID,omitempty"`
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

// Get the current registry scan settings.
func Get(c pcc.Client) (Registry, error) {
	var ans Registry
	if err := c.Request(http.MethodGet, Endpoint, nil, nil, &ans); err != nil {
		return ans, err
	}
	return ans, nil
}

// Update the current registry scan settings.
func Update(c pcc.Client, registry Registry) error {
	return c.Request(http.MethodPut, Endpoint, nil, registry, nil)
}
