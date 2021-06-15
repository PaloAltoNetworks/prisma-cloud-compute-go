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
	EffectAlert	= "alert"
	EffectAllow	= "allow"
	EffectBan	= "ban"
	EffectBlock	= "block"
	EffectDeny	= "deny"
	EffectDisable	= "disable"
	EffectIgnore	= "ignore"
	EffectPrevent	= "prevent"
)
