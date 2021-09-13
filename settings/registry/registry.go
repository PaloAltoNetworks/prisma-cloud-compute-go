package registry

import (
	"net/http"

	pcc "github.com/paloaltonetworks/prisma-cloud-compute-go"
)

const endpoint = "api/v1/settings/registry"

type Registry struct {
	Specifications []Specification `json:"specifications,omitempty"`
}

type Specification struct {
	Cap                      int        `json:"cap,omitempty"`
	Collections              []string   `json:"collections,omitempty"`
	Credential               Credential `json:"credential,omitempty"`
	CredentialId             string     `json:"credentialID,omitempty"`
	ExcludedRepositories     []string   `json:"excludedRepositories,omitempty"`
	ExcludedTags             string     `json:"excludedTags,omitempty"`
	HarborDeploymentSecurity bool       `json:"harborDeploymentSecurity,omitempty"`
	JfrogRepoTypes           []string   `json:"jfrogRepoTypes,omitempty"`
	Namespace                string     `json:"namespace,omitempty"`
	Os                       string     `json:"os,omitempty"`
	Tag                      string     `json:"tag,omitempty"`
	Registry                 string     `json:"registry,omitempty"`
	Repository               string     `json:"repository,omitempty"`
	Scanners                 int        `json:"scanners,omitempty"`
	Version                  string     `json:"version,omitempty"`
	VersionPattern           string     `json:"versionPattern,omitempty"`
}

type Credential struct {
	AccountGuid  string       `json:"accountGUID,omitempty"`
	AccountId    string       `json:"accountID,omitempty"`
	ApiToken     StringResult `json:"apiToken,omitempty"`
	CaCert       string       `json:"caCert,omitempty"`
	Created      string       `json:"created,omitempty"`
	Description  string       `json:"description,omitempty"`
	External     bool         `json:"external"`
	Id           string       `json:"_id,omitempty"`
	LastModified string       `json:"lastModified,omitempty"`
	Owner        string       `json:"owner,omitempty"`
	RoleArn      string       `json:"roleArn,omitempty"`
	Secret       StringResult `json:"secret,omitempty"`
	SkipVerify   bool         `json:"skipVerify"`
	Type         string       `json:"type,omitempty"`
	Url          string       `json:"url,omitempty"`
	UseAwsRole   bool         `json:"useAWSRole"`
}

type StringResult struct {
	Encrypted string `json:"encrypted,omitempty"`
	Plain     string `json:"plain,omitempty"`
}

type Token struct {
	AwsAccessKeyId     string       `json:"awsAccessKeyId,omitempty"`
	AwsSecretAccessKey string       `json:"awsSecretAccessKey,omitempty"`
	Duration           int          `json:"duration,omitempty"`
	ExpirationTime     string       `json:"expirationTime,omitempty"`
	Token              StringResult `json:"token,omitempty"`
}

// Return the registry scan settings.
func Get(c pcc.Client) (Registry, error) {
	var ans Registry
	if err := c.Communicate(http.MethodGet, endpoint, nil, nil, &ans); err != nil {
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
	err := c.Communicate(method, endpoint, nil, data, nil)
	return err
}
