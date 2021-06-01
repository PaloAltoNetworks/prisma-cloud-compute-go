package policyRuntimeContainer

import (
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collection"
)

type Policy struct {
	PolicyId                   string               `json:"_id,omitempty"`
    LearningDisabled           bool                 `json:"learningDisabled,omitempty"`
	Rules                      []Rule                `json:"rules,omitempty"`    
}

type Rule struct {
	AdvancedProtection         bool                        `json:"advancedProtection,omitempty"`
	CloudMetadataEnforcement   bool                        `json:"cloudMetadataEnforcement,omitempty"`
	Collections                []collection.Collection     `json:"collections,omitempty"`
	CustomRules                []CustomRule                `json:"customRules,omitempty"`
	Disabled                   bool                        `json:"disabled,omitempty"`
	Dns                        Dns                         `json:"dns,omitempty"`
	Filesystem                 Filesystem                  `json:"filesystem,omitempty"`
	KubernetesEnforcement      bool                        `json:"kubernetesEnforcement,omitempty"`
	Modified                   string                      `json:"modified,omitempty"`
	Name                       string                      `json:"name,omitempty"`
	Network                    Network                     `json:"network,omitempty"`
	Notes                      string                      `json:"notes,omitempty"`
	Owner                      string                      `json:"owner,omitempty"`
    PreviousName               string                      `json:"previousName,omitempty"`
	Processes                  Processes                   `json:"processes,omitempty"`
	WildFireAnalysis           []string                    `json:"wildFireAnalysis,omitempty"`    
}

type CustomRule struct {
	Id            int                  `json:"_id,omitempty"`
	Action         []string             `json:"action,omitempty"`
	Effect         []string             `json:"effect,omitempty"`    
}

type Dns struct {
    Blacklist           []string               `json:"blacklist,omitempty"`
	Effect              []string               `json:"effect,omitempty"`
	Whitelist           []string               `json:"whitelist,omitempty"`
}

type Filesystem struct {
	BackdoorFiles              bool                 `json:"backdoorFiles,omitempty"`
    Blacklist                  []string             `json:"blacklist,omitempty"`
	CheckNewFiles              bool                 `json:"checkNewFiles,omitempty"`
	Effect                     []string             `json:"effect,omitempty"`
	SkipEncryptedBinaries      bool                 `json:"skipEncryptedBinaries,omitempty"`
	SuspiciousELFHeaders       bool                 `json:"suspiciousELFHeaders,omitempty"`
	Whitelist                  string               `json:"whitelist,omitempty"`
}

type Network struct {
	BlacklistIPs               []string        `json:"blacklistIPs,omitempty"`
	BlacklistListeningPorts    []ListPort      `json:"blacklistListeningPorts,omitempty"`
    BlacklistOutboundPorts     []ListPort      `json:"blacklistOutboundPorts,omitempty"`
	DetectPortScan             bool            `json:"detectPortScan,omitempty"`
	Effect                     []string        `json:"effect,omitempty"`
	SkipModifiedProc           bool            `json:"skipModifiedProc,omitempty"`
	SkipRawSockets             bool            `json:"skipRawSockets,omitempty"`
	WhitelistIPs               string          `json:"whitelistIPs,omitempty"`
    WhitelistListeningPorts    []ListPort      `json:"whitelistListeningPorts,omitempty"`
    WhitelistOutboundPorts     []ListPort      `json:"whitelistOutboundPorts,omitempty"`
}


type Processes struct {
    Blacklist                  []string        `json:"blacklist,omitempty"`
	BlockAllBinaries           bool            `json:"blockAllBinaries,omitempty"`
	CheckCryptoMiners          bool            `json:"checkCryptoMiners,omitempty"`
	CheckLateralMovement       bool            `json:"checkLateralMovement,omitempty"`
	CheckNewBinaries           bool            `json:"checkNewBinaries,omitempty"`
	CheckParentChild           bool            `json:"checkParentChild,omitempty"`
	CheckSuidBinaries          bool            `json:"checkSuidBinaries,omitempty"`
    Effect                     []string        `json:"effect,omitempty"`
	SkipModified               bool            `json:"skipModified,omitempty"`
	SkipReverseShell           bool            `json:"skipReverseShell,omitempty"`
    Whitelist                  []string         `json:"whitelist,omitempty"`
}

type ListPort struct {
	Deny       bool                `json:"deny,omitempty"`
	End        int                 `json:"end,omitempty"`
	Start      int                 `json:"start,omitempty"`
}