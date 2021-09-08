package pcc

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestAPIClient(t *testing.T) {
	credsFile, err := os.Open("examples/creds.json")
	if err != nil {
		fmt.Printf("error opening creds file: %v", err)
	}
	defer credsFile.Close()

	fileContent, err := io.ReadAll(credsFile)
	if err != nil {
		fmt.Printf("error reading creds file: %v", err)
		return
	}
	var creds Credentials
	if err := json.Unmarshal(fileContent, &creds); err != nil {
		fmt.Printf("error unmarshalling creds file: %v", err)
		return
	}

	client, _ := APIClient(creds.ConsoleURL, creds.Username, creds.Password, creds.SkipCertVerification)
	if client.JWT == "" {
		t.Errorf("JWT is empty. Authenticate did not work.")
	}
}
