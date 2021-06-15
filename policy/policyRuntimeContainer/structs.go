package policyRuntimeContainer

import (
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"
)

type Policy struct {
	PolicyId                   string               `json:"_id,omitempty"`
    	LearningDisabled           bool                 `json:"learningDisabled,omitempty"`
	Rules                      []policy.Rule        `json:"rules,omitempty"`    
}
