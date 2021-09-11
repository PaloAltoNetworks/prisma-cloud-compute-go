package policies

import (
	"net/http"

	pcc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collections"
)

const (
	ComplianceCiImagesEndpoint    = "api/v1/policies/compliance/ci/images"
	ComplianceContainerEndpoint   = "api/v1/policies/compliance/container"
	ComplianceHostEndpoint        = "api/v1/policies/compliance/host"
	RuntimeContainerEndpoint      = "api/v1/policies/runtime/container"
	RuntimeHostEndpoint           = "api/v1/policies/runtime/host"
	VulnerabilityCiImagesEndpoint = "api/v1/policies/vulnerability/ci/images"
	VulnerabilityHostEndpoint     = "api/v1/policies/vulnerability/host"
	VulnerabilityImagesEndpoint   = "api/v1/policies/vulnerability/images"
)

type Policy struct {
	LearningDisabled bool   `json:"learningDisabled,omitempty"`
	PolicyId         string `json:"_id,omitempty"`
	PolicyType       string `json:"policyType,omitempty"`
	Rules            []Rule `json:"rules,omitempty"`
}

type Rule struct {
	Action                   []string                 `json:"action,omitempty"`
	AdvancedProtection       bool                     `json:"advancedProtection"`
	AlertThreshold           Threshold                `json:"alertThreshold,omitempty"`
	AllCompliance            bool                     `json:"allCompliance"`
	AntiMalware              AntiMalware              `json:"antiMalware,omitempty"`
	AuditAllowed             bool                     `json:"auditAllowed"`
	BlockMsg                 string                   `json:"blockMsg,omitempty"`
	BlockThreshold           Threshold                `json:"blockThreshold,omitempty"`
	CloudMetadataEnforcement bool                     `json:"cloudMetadataEnforcement"`
	Collections              []collections.Collection `json:"collections,omitempty"`
	Condition                Condition                `json:"condition,omitempty"`
	CustomRules              []CustomRule             `json:"customRules,omitempty"`
	CveRules                 []CveRule                `json:"cveRules,omitempty"`
	Disabled                 bool                     `json:"disabled"`
	Dns                      Dns                      `json:"dns,omitempty"`
	Effect                   string                   `json:"effect,omitempty"`
	FileIntegrityRules       []FileIntegrityRule      `json:"fileIntegrityRules,omitempty"`
	Filesystem               Filesystem               `json:"filesystem,omitempty"`
	Forensic                 Forensic                 `json:"forensic,omitempty"`
	GraceDays                int                      `json:"graceDays,omitempty"`
	Group                    []string                 `json:"group,omitempty"`
	KubernetesEnforcement    bool                     `json:"kubernetesEnforcement"`
	License                  License                  `json:"license,omitempty"`
	LogInspectionRules       []LogInspectionRule      `json:"logInspectionRules,omitempty"`
	Modified                 string                   `json:"modified,omitempty"`
	Name                     string                   `json:"name,omitempty"`
	Network                  Network                  `json:"network,omitempty"`
	Notes                    string                   `json:"notes,omitempty"`
	OnlyFixed                bool                     `json:"onlyFixed"`
	Owner                    string                   `json:"owner,omitempty"`
	PreviousName             string                   `json:"previousName,omitempty"`
	Principal                []string                 `json:"principal,omitempty"`
	Processes                Processes                `json:"processes,omitempty"`
	Tags                     []Tag                    `json:"tags,omitempty"`
	Verbose                  bool                     `json:"verbose"`
	WildFireAnalysis         string                   `json:"wildFireAnalysis,omitempty"`
}

type FileIntegrityRule struct {
	Dir           bool     `json:"dir"`
	Exclusions    []string `json:"exclusions,omitempty"`
	Metadata      bool     `json:"metadata"`
	Path          string   `json:"path,omitempty"`
	ProcWhitelist []string `json:"procWhitelist,omitempty"`
	Read          bool     `json:"read"`
	Recursive     bool     `json:"recursive"`
	Write         bool     `json:"write"`
}

type Forensic struct {
	ActivitiesDisabled       bool `json:"activitiesDisabled"`
	DockerEnabled            bool `json:"dockerEnabled"`
	ReadonlyDockerEnabled    bool `json:"readonlyDockerEnabled"`
	ServiceActivitiesEnabled bool `json:"serviceActivitiesEnabled"`
	SshdEnabled              bool `json:"sshdEnabled"`
	SudoEnabled              bool `json:"sudoEnabled"`
}
type LogInspectionRule struct {
	Path  string   `json:"path,omitempty"`
	Regex []string `json:"regex,omitempty"`
}

type DeniedProcesses struct {
	Effect string   `json:"effect,omitempty"`
	Paths  []string `json:"paths,omitempty"`
}

type AntiMalware struct {
	AllowedProcesses              []string        `json:"allowedProcesses,omitempty"`
	CryptoMiner                   string          `json:"cryptoMiner,omitempty"`
	CustomFeed                    string          `json:"customFeed,omitempty"`
	DeniedProcesses               DeniedProcesses `json:"deniedProcesses,omitempty"`
	DetectCompilerGeneratedBinary bool            `json:"detectCompilerGeneratedBinary"`
	EncryptedBinaries             string          `json:"encryptedBinaries,omitempty"`
	ExecutionFlowHijack           string          `json:"executionFlowHijack,omitempty"`
	IntelligenceFeed              string          `json:"intelligenceFeed,omitempty"`
	Paths                         []string        `json:"paths,omitempty"`
	ReverseShell                  string          `json:"reverseShell,omitempty"`
	ServiceUnknownOriginBinary    string          `json:"serviceUnknownOriginBinary,omitempty"`
	SkipSSHTracking               bool            `json:"skipSSHTracking"`
	SuspiciousELFHeaders          string          `json:"suspiciousELFHeaders,omitempty"`
	TempFSProc                    string          `json:"tempFSProc,omitempty"`
	UserUnknownOriginBinary       string          `json:"userUnknownOriginBinary,omitempty"`
	WebShell                      string          `json:"webShell,omitempty"`
	WildFireAnalysis              string          `json:"wildFireAnalysis,omitempty"`
}

type Threshold struct {
	Disabled bool `json:"disabled"`
	Enabled  bool `json:"enabled"`
	Value    int  `json:"value,omitempty"`
}

type Condition struct {
	Device          string          `json:"device,omitempty"`
	Readonly        bool            `json:"readonly"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities,omitempty"`
}

type Vulnerability struct {
	Block       bool `json:"block"`
	Id          int  `json:"id,omitempty"`
	MinSeverity int  `json:"minSeverity,omitempty"`
}

type CveRule struct {
	Description string     `json:"description,omitempty"`
	Effect      string     `json:"effect,omitempty"`
	Expiration  Expiration `json:"expiration,omitempty"`
	Id          string     `json:"id,omitempty"`
}

type Expiration struct {
	Date    string `json:"date,omitempty"`
	Enabled bool   `json:"enabled"`
}

type License struct {
	AlertThreshold Threshold `json:"alertThreshold,omitempty"`
	BlockThreshold Threshold `json:"blockThreshold,omitempty"`
	Critical       []string  `json:"critical,omitempty"`
	High           []string  `json:"high,omitempty"`
	Low            []string  `json:"low,omitempty"`
	Medium         []string  `json:"medium,omitempty"`
}

type Tag struct {
	Description string     `json:"description,omitempty"`
	Effect      string     `json:"effect,omitempty"`
	Expiration  Expiration `json:"expiration,omitempty"`
	Name        string     `json:"name,omitempty"`
}

type CustomRule struct {
	Action []string `json:"action,omitempty"`
	Effect string   `json:"effect,omitempty"`
	Id     int      `json:"_id,omitempty"`
}

type Dns struct {
	Allow            []string `json:"allow,omitempty"`
	Blacklist        []string `json:"blacklist,omitempty"`
	Deny             []string `json:"deny,omitempty"`
	DenyListEffect   string   `json:"denyListEffect,omitempty"`
	Effect           string   `json:"effect,omitempty"`
	IntelligenceFeed string   `json:"intelligenceFeed,omitempty"`
	Whitelist        []string `json:"whitelist,omitempty"`
}

type Filesystem struct {
	BackdoorFiles         bool     `json:"backdoorFiles"`
	Blacklist             []string `json:"blacklist,omitempty"`
	CheckNewFiles         bool     `json:"checkNewFiles"`
	Effect                string   `json:"effect,omitempty"`
	SkipEncryptedBinaries bool     `json:"skipEncryptedBinaries"`
	SuspiciousELFHeaders  bool     `json:"suspiciousELFHeaders"`
	Whitelist             []string `json:"whitelist,omitempty"`
}

type Network struct {
	AllowedOutboundIPs      []string   `json:"allowedOutboundIPs,omitempty"`
	BlacklistIPs            []string   `json:"blacklistIPs,omitempty"`
	BlacklistListeningPorts []ListPort `json:"blacklistListeningPorts,omitempty"`
	BlacklistOutboundPorts  []ListPort `json:"blacklistOutboundPorts,omitempty"`
	CustomFeed              string     `json:"customFeed,omitempty"`
	DeniedListeningPorts    []ListPort `json:"deniedListeningPorts,omitempty"`
	DeniedOutboundIPs       []string   `json:"deniedOutboundIPs,omitempty"`
	DeniedOutboundPorts     []ListPort `json:"deniedOutboundPorts,omitempty"`
	DenyListEffect          string     `json:"denyListEffect,omitempty"`
	DetectPortScan          bool       `json:"detectPortScan"`
	Effect                  string     `json:"effect,omitempty"`
	IntelligenceFeed        string     `json:"intelligenceFeed,omitempty"`
	SkipModifiedProc        bool       `json:"skipModifiedProc"`
	SkipRawSockets          bool       `json:"skipRawSockets"`
	WhitelistIPs            []string   `json:"whitelistIPs,omitempty"`
	WhitelistListeningPorts []ListPort `json:"whitelistListeningPorts,omitempty"`
	WhitelistOutboundPorts  []ListPort `json:"whitelistOutboundPorts,omitempty"`
}

type Processes struct {
	Blacklist            []string `json:"blacklist,omitempty"`
	BlockAllBinaries     bool     `json:"blockAllBinaries"`
	CheckCryptoMiners    bool     `json:"checkCryptoMiners"`
	CheckLateralMovement bool     `json:"checkLateralMovement"`
	CheckNewBinaries     bool     `json:"checkNewBinaries"`
	CheckParentChild     bool     `json:"checkParentChild"`
	CheckSuidBinaries    bool     `json:"checkSuidBinaries"`
	Effect               string   `json:"effect,omitempty"`
	SkipModified         bool     `json:"skipModified"`
	SkipReverseShell     bool     `json:"skipReverseShell"`
	Whitelist            []string `json:"whitelist,omitempty"`
}

type ListPort struct {
	Deny  bool `json:"deny"`
	End   int  `json:"end,omitempty"`
	Start int  `json:"start,omitempty"`
}

// Return the current policy.
func Get(c pcc.Client, endpoint string) (Policy, error) {
	var ans Policy
	if err := c.Communicate(http.MethodGet, endpoint, nil, nil, &ans); err != nil {
		return ans, err
	}
	return ans, nil
}

// Update the current policy.
func Update(c pcc.Client, endpoint string, policy Policy) error {
	err := c.Communicate(http.MethodPut, endpoint, nil, policy, nil)
	return err
}
