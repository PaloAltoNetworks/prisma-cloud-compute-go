package pcc

import (
	"testing"
)

func TestAPIClient(t *testing.T) {
	client, _ := APIClient("https://twistlock.wfg.lab.twistlock.com", "admin", "geauxtigers", false)
	if client.JWT == "" {
		t.Errorf("JWT is empty. Authenticate did not work.")
	}
}
