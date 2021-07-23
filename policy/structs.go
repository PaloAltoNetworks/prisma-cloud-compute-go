package policy

import (
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collection"
)

type Rule struct {
	Action				[]string			`json:"action,omitempty"`
	AlertThreshold			Threshold			`json:"alertThreshold,omitempty"`
	AllCompliance			bool				`json:"allCompliance,omitempty"`
	AuditAllowed			bool				`json:"auditAllowed,omitempty"`
	BlockMsg			string				`json:"blockMsg,omitempty"`
	BlockThreshold			Threshold			`json:"blockThreshold,omitempty"`
	Condition			Condition			`json:"condition,omitempty"`
	CveRules			[]CveRule			`json:"cveRules,omitempty"`
	Effect				string				`json:"effect,omitempty"`
	GraceDays			int				`json:"graceDays,omitempty"`
	Group				[]string			`json:"group,omitempty"`
	License			License			`json:"license,omitempty"`
	OnlyFixed			bool				`json:"onlyFixed,omitempty"`
	Principal			[]string			`json:"principal,omitempty"`
	Tags				[]Tag				`json:"tags,omitempty"`
	Verbose			bool				`json:"verbose,omitempty"`
	AdvancedProtection		bool				`json:"advancedProtection,omitempty"`
	CloudMetadataEnforcement	bool				`json:"cloudMetadataEnforcement,omitempty"`
	Collections			[]collection.Collection	`json:"collections,omitempty"`
	CustomRules			[]CustomRule			`json:"customRules,omitempty"`
	Disabled 			bool				`json:"disabled,omitempty"`
	Dns				Dns				`json:"dns,omitempty"`
	Filesystem			Filesystem			`json:"filesystem,omitempty"`
	KubernetesEnforcement		bool				`json:"kubernetesEnforcement,omitempty"`
	Modified			string				`json:"modified,omitempty"`
	Name				string				`json:"name,omitempty"`
	Network			Network			`json:"network,omitempty"`
	Notes				string				`json:"notes,omitempty"`
	Owner				string				`json:"owner,omitempty"`
	PreviousName			string				`json:"previousName,omitempty"`
	Processes			Processes			`json:"processes,omitempty"`
	WildFireAnalysis		string				`json:"wildFireAnalysis,omitempty"`    
}

type Threshold struct {
	Enabled		bool			`json:"enabled,omitempty"`
	Disabled		bool			`json:"disabled,omitempty"`
	Value			int			`json:"value,omitempty"`
}

type Condition struct {
	Device			string			`json:"device,omitempty"`
	Readonly		bool			`json:"readonly,omitempty"`
	Vulnerabilities	[]Vulnerability	`json:"vulnerabilities,omitempty"`
}

type Vulnerability struct {
	Block			bool			`json:"block"`
	Id			int			`json:"id,omitempty"`
	MinSeverity			int			`json:"minSeverity,omitempty"`
}

type CveRule struct {
	Description		string			`json:"description,omitempty"`
	Effect			string			`json:"effect,omitempty"`
	Id			string			`json:"id,omitempty"`
	Expiration		Expiration		`json:"expiration,omitempty"`
}

type Expiration struct {
	Date			string			`json:"date,omitempty"`
	Enabled		bool			`json:"enabled,omitempty"`
}

type License struct {
	AlertThreshold		Threshold		`json:"alertThreshold,omitempty"`
	BlockThreshold		Threshold		`json:"blockThreshold,omitempty"`
	Critical		[]string		`json:"critical,omitempty"`
	High			[]string		`json:"high,omitempty"`
	Low			[]string		`json:"low,omitempty"`
	Medium			[]string		`json:"medium,omitempty"`
}

type Tag struct {
	Description		string			`json:"description,omitempty"`
	Effect			string			`json:"effect,omitempty"`
	Expiration		[]Expiration		`json:"expiration,omitempty"`
	Name			string			`json:"name,omitempty"`
}

type CustomRule struct {
	Id			int			`json:"_id,omitempty"`
	Action         	[]string		`json:"action,omitempty"`
	Effect         	string			`json:"effect,omitempty"`    
}

type Dns struct {
    	Blacklist		[]string		`json:"blacklist,omitempty"`
	Effect			string			`json:"effect,omitempty"`
	Whitelist		[]string		`json:"whitelist,omitempty"`
}

type Filesystem struct {
	BackdoorFiles			bool		`json:"backdoorFiles,omitempty"`
	Blacklist			[]string	`json:"blacklist,omitempty"`
	CheckNewFiles			bool		`json:"checkNewFiles,omitempty"`
	Effect				string		`json:"effect,omitempty"`
	SkipEncryptedBinaries		bool		`json:"skipEncryptedBinaries,omitempty"`
	SuspiciousELFHeaders		bool		`json:"suspiciousELFHeaders,omitempty"`
	Whitelist			[]string	`json:"whitelist,omitempty"`
}

type Network struct {
	BlacklistIPs			[]string	`json:"blacklistIPs,omitempty"`
	BlacklistListeningPorts	[]ListPort	`json:"blacklistListeningPorts,omitempty"`
	BlacklistOutboundPorts		[]ListPort	`json:"blacklistOutboundPorts,omitempty"`
	DetectPortScan			bool		`json:"detectPortScan,omitempty"`
	Effect				string		`json:"effect,omitempty"`
	SkipModifiedProc		bool		`json:"skipModifiedProc,omitempty"`
	SkipRawSockets			bool		`json:"skipRawSockets,omitempty"`
	WhitelistIPs			[]string	`json:"whitelistIPs,omitempty"`
    	WhitelistListeningPorts	[]ListPort	`json:"whitelistListeningPorts,omitempty"`
    	WhitelistOutboundPorts		[]ListPort	`json:"whitelistOutboundPorts,omitempty"`
}


type Processes struct {
    	Blacklist			[]string	`json:"blacklist,omitempty"`
	BlockAllBinaries		bool		`json:"blockAllBinaries,omitempty"`
	CheckCryptoMiners		bool		`json:"checkCryptoMiners,omitempty"`
	CheckLateralMovement		bool		`json:"checkLateralMovement,omitempty"`
	CheckNewBinaries		bool		`json:"checkNewBinaries,omitempty"`
	CheckParentChild		bool		`json:"checkParentChild,omitempty"`
	CheckSuidBinaries		bool		`json:"checkSuidBinaries,omitempty"`
    	Effect				string		`json:"effect,omitempty"`
	SkipModified			bool		`json:"skipModified,omitempty"`
	SkipReverseShell		bool		`json:"skipReverseShell,omitempty"`
    	Whitelist			[]string	`json:"whitelist,omitempty"`
}

type ListPort struct {
	Deny				bool		`json:"deny,omitempty"`
	End				int		`json:"end,omitempty"`
	Start				int		`json:"start,omitempty"`
}
