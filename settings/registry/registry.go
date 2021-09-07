package registry

import (
	"net/http"

	pcc "github.com/paloaltonetworks/prisma-cloud-compute-go"
)

const endpoint = "api/v1/settings/registry"

type Registry struct {
	HarborScannerUrlSuffix string          `json:"harborScannerUrlSuffix,omitempty"`
	Specifications         []Specification `json:"specifications,omitempty"`
	WebhookUrlSuffix       string          `json:"webhookUrlSuffix,omitempty"`
}

type Specification struct {
	Cap                      int        `json:"cap,omitempty"`
	Collections              []string   `json:"collections,omitempty"`
	Credential               Credential `json:"credential,omitempty"`
	CredentialID             string     `json:"credentialID,omitempty"`
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
	AccountGUID  string       `json:"accountGUID,omitempty"`
	AccountID    string       `json:"accountID,omitempty"`
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
	UseAWSRole   bool         `json:"useAWSRole"`
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

// Update the registry scan settings.
func Update(c pcc.Client, registry Registry) error {
	err := c.Communicate(http.MethodPut, endpoint, nil, registry, nil)
	return err
}
