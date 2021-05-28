package policyRuntimeContainer

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
	EffectBlock    = "block"
	EffectPrevent  = "prevent"
	EffectAlert    = "alert"
	EffectAllow    = "allow"
	EffectBan      = "ban"
	EffectDisable  = "disable"
)

/*
// Valid values for Policy.Rule.RuleType.
const (
	RuleTypeConfig     = "Config"
	RuleTypeAuditEvent = "AuditEvent"
	RuleTypeNetwork    = "Network"
)

// Valid values for Policy.PolicyType.
const (
	PolicyTypeConfig     = "config"
	PolicyTypeAuditEvent = "audit_event"
	PolicyTypeNetwork    = "network"
)

// Valid values for Policy.Rule.Severity.
const (
	SeverityLow    = "low"
	SeverityMedium = "medium"
	SeverityHigh   = "high"
)
*/
var Suffix = []string{"policy"}
