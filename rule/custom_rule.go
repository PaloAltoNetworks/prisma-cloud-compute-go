package rule

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/paloaltonetworks/prisma-cloud-compute-go/pcc"
)

const CustomRulesEndpoint = "api/v1/custom-rules"

type CustomRule struct {
	Description string `json:"description,omitempty"`
	Id          int    `json:"_id,omitempty"`
	Message     string `json:"message,omitempty"`
	Name        string `json:"name,omitempty"`
	Script      string `json:"script,omitempty"`
	Type        string `json:"type,omitempty"`
}

// Get all custom rules.
func ListCustomRules(c pcc.Client) ([]CustomRule, error) {
	var ans []CustomRule
	if err := c.Request(http.MethodGet, CustomRulesEndpoint, nil, nil, &ans); err != nil {
		return nil, fmt.Errorf("error listing custom rules: %s", err)
	}
	return ans, nil
}

// Get a specific custom rule.
func GetCustomRule(c pcc.Client, id int) (*CustomRule, error) {
	rules, err := ListCustomRules(c)
	if err != nil {
		return nil, fmt.Errorf("error getting custom rules '%d': %s", id, err)
	}
	for _, val := range rules {
		if val.Id == id {
			return &val, nil
		}
	}
	return nil, fmt.Errorf("custom rule '%d' not found", id)
}

// Create a new custom rule.
func CreateCustomRule(c pcc.Client, rule CustomRule) (int, error) {
	id, err := GenerateCustomRuleId(c)
	if err != nil {
		return -1, fmt.Errorf("error getting custom rules '%d': %s", id, err)
	}
	rule.Id = id
	return id, UpdateCustomRule(c, rule)
}

// Helper method to generate id for new custom rule
func GenerateCustomRuleId(c pcc.Client) (int, error) {
	rules, err := ListCustomRules(c)
	if err != nil {
		return -1, fmt.Errorf("error getting custom rules: %s", err)
	}
	rand.Seed(time.Now().UnixNano())
	for {
		id := rand.Intn(100000000)
		unique := true
		for _, val := range rules {
			if val.Id == id {
				unique = false
			}
		}
		if unique {
			return id, nil
		}
	}
	return -1, fmt.Errorf("error generating custom rule id")
}

// Update an existing collection.
func UpdateCustomRule(c pcc.Client, rule CustomRule) error {
	return c.Request(http.MethodPut, fmt.Sprintf("%s/%d", CustomRulesEndpoint, rule.Id), nil, rule, nil)
}

// Delete an existing collection.
func DeleteCustomRule(c pcc.Client, id int) error {
	return c.Request(http.MethodDelete, fmt.Sprintf("%s/%d", CustomRulesEndpoint, id), nil, nil, nil)
}
