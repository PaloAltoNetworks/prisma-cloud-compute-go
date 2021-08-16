package policy

const (
	singular = "policy"
	plural   = "policies"
)

// Valid values for action
const (
	ActionAudit    = "audit"
	ActionIncident = "incident"
)

// Valid values for effect
const (
	EffectAlert   = "alert"
	EffectAllow   = "allow"
	EffectBan     = "ban"
	EffectBlock   = "block"
	EffectDeny    = "deny"
	EffectDisable = "disable"
	EffectIgnore  = "ignore"
	EffectPrevent = "prevent"
)

const (
	PolicyTypeDocker                    = "docker"
	PolicyTypeContainerVulnerability    = "containerVulnerability"
	PolicyTypeContainerCompliance       = "containerCompliance"
	PolicyTypeCiImagesVulnerability     = "ciImagesVulnerability"
	PolicyTypeCiImagesCompliance        = "ciImagesCompliance"
	PolicyTypeHostVulnerability         = "hostVulnerability"
	PolicyTypeHostCompliance            = "hostCompliance"
	PolicyTypeVmVulnerability           = "vmVulnerability"
	PolicyTypeVmCompliance              = "vmCompliance"
	PolicyTypeServerlessCompliance      = "serverlessCompliance"
	PolicyTypeCiServerlessCompliance    = "ciServerlessCompliance"
	PolicyTypeServerlessVulnerability   = "serverlessVulnerability"
	PolicyTypeCiServerlessVulnerability = "ciServerlessVulnerability"
	PolicyTypeContainerRuntime          = "containerRuntime"
	PolicyTypeAppEmbeddedRuntime        = "appEmbeddedRuntime"
	PolicyTypeContainerAppFirewall      = "containerAppFirewall"
	PolicyTypeHostAppFirewall           = "hostAppFirewall"
	PolicyTypeAppEmbeddedAppFirewall    = "appEmbeddedAppFirewall"
	PolicyTypeServerlessAppFirewall     = "serverlessAppFirewall"
	PolicyTypeNetworkFirewall           = "networkFirewall"
	PolicyTypeSecrets                   = "secrets"
	PolicyTypeHostRuntime               = "hostRuntime"
	PolicyTypeServerlessRuntime         = "serverlessRuntime"
	PolicyTypeKubernetesAudit           = "kubernetesAudit"
	PolicyTypeTrust                     = "trust"
	PolicyTypeCloud                     = "cloud"
	PolicyTypeAdmission                 = "admission"
	PolicyTypeCodeRepoVulnerability     = "codeRepoVulnerability"
	PolicyTypeCiCodeRepoVulnerability   = "ciCodeRepoVulnerability"
	PolicyTypeCodeRepoCompliance        = "codeRepoCompliance"
	PolicyTypeCiCodeRepoCompliance      = "ciCodeRepoCompliance"
)
