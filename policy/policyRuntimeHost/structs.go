package policyRuntimeHost

import (
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"
)

type Policy struct {
	PolicyId                   string               `json:"_id,omitempty"`
    	Owner           string                 `json:"owner,omitempty"`
	Rules                      []policy.Rule        `json:"rules,omitempty"`    
}
