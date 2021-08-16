package policyComplianceCiImages

import (
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"
)

type Policy struct {
	PolicyId   string        `json:"_id,omitempty"`
	PolicyType string        `json:"policyType,omitempty"`
	Rules      []policy.Rule `json:"rules,omitempty"`
}
