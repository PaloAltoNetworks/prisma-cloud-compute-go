package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	pcc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collection"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/settings/registry"
)

func main() {
	credsFile, err := os.Open("creds.json")
	if err != nil {
		fmt.Printf("error opening creds file: %v", err)
	}
	defer credsFile.Close()

	fileContent, err := io.ReadAll(credsFile)
	if err != nil {
		fmt.Printf("error reading creds file: %v", err)
		return
	}
	var creds pcc.Credentials
	if err := json.Unmarshal(fileContent, &creds); err != nil {
		fmt.Printf("error unmarshalling creds file: %v", err)
		return
	}

	client, err := pcc.APIClient(creds.ConsoleURL, creds.Username, creds.Password, creds.SkipCertVerification)
	if err != nil {
		fmt.Printf("failed creating API client: %v", err)
		return
	}

	/*
		COLLECTIONS
	*/
	fmt.Printf("create collection\n")
	newColl := collection.Collection{Name: "My Collection"}
	err = collection.Create(*client, newColl)
	if err != nil {
		fmt.Printf("Failed to create collection: %s\n", err)
		return
	}

	fmt.Printf("\nlist collections:\n")
	colls, err := collection.List(*client)
	if err != nil {
		fmt.Printf("Failed to list collections: %s\n", err)
	}

	for _, v := range colls {
		fmt.Printf("* %s %s\n", v.Name, v.Color)
	}

	fmt.Printf("\nget collection:\n")
	coll, err := collection.Get(*client, "My Collection")
	if err != nil {
		fmt.Printf("Failed to get collection: %s\n", err)
	}
	fmt.Printf("* %s %s\n", coll.Name, coll.Color)

	fmt.Printf("\nupdate collection\n")
	existingColl := collection.Collection{Name: "My Collection", Color: "#FFFFFF"}
	err = collection.Update(*client, existingColl)
	if err != nil {
		fmt.Printf("Failed to update collection: %s\n", err)
	}

	fmt.Printf("\nlist collections:\n")
	colls, err = collection.List(*client)
	if err != nil {
		fmt.Printf("Failed to list collections: %s\n", err)
	}

	for _, v := range colls {
		fmt.Printf("* %s %s\n", v.Name, v.Color)
	}

	fmt.Printf("\ndelete collection\n")
	err = collection.Delete(*client, "My Collection")
	if err != nil {
		fmt.Printf("Failed to delete collection: %s\n", err)
	}

	fmt.Printf("\nlist collections:\n")
	colls, err = collection.List(*client)
	if err != nil {
		fmt.Printf("Failed to list collections: %s\n", err)
	}

	for _, v := range colls {
		fmt.Printf("* %s %s\n", v.Name, v.Color)
	}

	/*
		CI IMAGE COMPLIANCE
	*/
	complianceCiImageColl := collection.Collection{Name: "All"}
	complianceCiImageVuln1 := policy.Vulnerability{Id: 41, Block: false}
	complianceCiImageVuln2 := policy.Vulnerability{Id: 422, Block: true}
	complianceCiImageCondition := policy.Condition{Vulnerabilities: []policy.Vulnerability{complianceCiImageVuln1, complianceCiImageVuln2}}
	complianceCiImageRule := policy.Rule{Name: "example ci image compliance rule", Effect: "alert, block", Collections: []collection.Collection{complianceCiImageColl}, Condition: complianceCiImageCondition}
	complianceCiImageRules := []policy.Rule{complianceCiImageRule}
	complianceCiImagePolicy := policy.Policy{PolicyType: "ciImagesCompliance", Rules: complianceCiImageRules}

	fmt.Printf("\nupdate CI image compliance policy\n")
	complianceCiImageErr := policy.Update(*client, policy.ComplianceCiImagesEndpoint, complianceCiImagePolicy)
	if complianceCiImageErr != nil {
		fmt.Printf("\nFailed to update CI image compliance policy: %v\n", complianceCiImageErr)
	}

	fmt.Printf("\nget CI image compliance policy:\n")
	retrievedPolicy, complianceCiImageErr := policy.Get(*client, policy.ComplianceCiImagesEndpoint)
	if complianceCiImageErr != nil {
		fmt.Printf("failed to get CI image compliance policy: %s\n", complianceCiImageErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	fmt.Printf("\nupdate CI image compliance policy\n")
	complianceCiImageRule.Name = "name change"
	complianceCiImagePolicy = policy.Policy{PolicyType: "ciImagesCompliance", Rules: []policy.Rule{complianceCiImageRule}}
	complianceCiImageErr = policy.Update(*client, policy.ComplianceCiImagesEndpoint, complianceCiImagePolicy)
	if complianceCiImageErr != nil {
		fmt.Printf("failed to update CI image compliance policy: %s\n", complianceCiImageErr)
	}

	fmt.Printf("\nget CI image compliance policy:\n")
	retrievedPolicy, complianceCiImageErr = policy.Get(*client, policy.ComplianceCiImagesEndpoint)
	if complianceCiImageErr != nil {
		fmt.Printf("failed to get CI image compliance policy: %s\n", complianceCiImageErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	/*
		CONTAINER COMPLIANCE
	*/
	complianceContainerColl := collection.Collection{Name: "All"}
	complianceContainerVuln1 := policy.Vulnerability{Id: 41, Block: false}
	complianceContainerVuln2 := policy.Vulnerability{Id: 422, Block: true}
	complianceContainerCondition := policy.Condition{Vulnerabilities: []policy.Vulnerability{complianceContainerVuln1, complianceContainerVuln2}}
	complianceContainerRule := policy.Rule{Name: "example container compliance rule", Effect: "alert, block", Collections: []collection.Collection{complianceContainerColl}, Condition: complianceContainerCondition}
	complianceContainerRules := []policy.Rule{complianceContainerRule}
	complianceContainerPolicy := policy.Policy{PolicyType: "containerCompliance", Rules: complianceContainerRules}

	fmt.Printf("\nupdate container compliance policy\n")
	complianceContainerErr := policy.Update(*client, policy.ComplianceContainerEndpoint, complianceContainerPolicy)
	if complianceContainerErr != nil {
		fmt.Printf("\nFailed to update container compliance policy: %v\n", complianceContainerErr)
	}

	fmt.Printf("\nget container compliance policy:\n")
	retrievedPolicy, complianceContainerErr = policy.Get(*client, policy.ComplianceContainerEndpoint)
	if complianceContainerErr != nil {
		fmt.Printf("failed to get container compliance policy: %s\n", complianceContainerErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	fmt.Printf("\nupdate container compliance policy\n")
	complianceContainerRule.Name = "name change"
	complianceContainerPolicy = policy.Policy{PolicyType: "containerCompliance", Rules: []policy.Rule{complianceContainerRule}}
	complianceContainerErr = policy.Update(*client, policy.ComplianceContainerEndpoint, complianceContainerPolicy)
	if complianceContainerErr != nil {
		fmt.Printf("failed to update container compliance policy: %s\n", complianceContainerErr)
	}

	fmt.Printf("\nget container compliance policy:\n")
	retrievedPolicy, complianceContainerErr = policy.Get(*client, policy.ComplianceContainerEndpoint)
	if complianceContainerErr != nil {
		fmt.Printf("failed to get container compliance policy: %s\n", complianceContainerErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	/*
		HOST COMPLIANCE
	*/
	complianceHostColl := collection.Collection{Name: "All"}
	complianceHostVuln1 := policy.Vulnerability{Id: 11, Block: false}
	complianceHostVuln2 := policy.Vulnerability{Id: 112, Block: true}
	complianceHostCondition := policy.Condition{Vulnerabilities: []policy.Vulnerability{complianceHostVuln1, complianceHostVuln2}}
	complianceHostRule := policy.Rule{Name: "example host compliance rule", Effect: "alert, block", Collections: []collection.Collection{complianceHostColl}, Condition: complianceHostCondition}
	complianceHostRules := []policy.Rule{complianceHostRule}
	complianceHostPolicy := policy.Policy{PolicyType: "hostCompliance", Rules: complianceHostRules}

	fmt.Printf("\nupdate host compliance policy\n")
	complianceHostErr := policy.Update(*client, policy.ComplianceHostEndpoint, complianceHostPolicy)
	if complianceHostErr != nil {
		fmt.Printf("\nFailed to update host compliance policy: %v\n", complianceHostErr)
	}

	fmt.Printf("\nget host compliance policy:\n")
	retrievedPolicy, complianceHostErr = policy.Get(*client, policy.ComplianceHostEndpoint)
	if complianceHostErr != nil {
		fmt.Printf("failed to get host compliance policy: %s\n", complianceHostErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	fmt.Printf("\nupdate host compliance policy\n")
	complianceHostRule.Name = "name change"
	complianceHostPolicy = policy.Policy{PolicyType: "hostCompliance", Rules: []policy.Rule{complianceHostRule}}
	complianceHostErr = policy.Update(*client, policy.ComplianceHostEndpoint, complianceHostPolicy)
	if complianceHostErr != nil {
		fmt.Printf("failed to update host compliance policy: %s\n", complianceHostErr)
	}

	fmt.Printf("\nget host compliance policy:\n")
	retrievedPolicy, complianceHostErr = policy.Get(*client, policy.ComplianceHostEndpoint)
	if complianceHostErr != nil {
		fmt.Printf("failed to get host compliance policy: %s\n", complianceHostErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	/*
		CONTAINER RUNTIME
	*/
	runtimeContainerColl := collection.Collection{Name: "All"}
	runtimeContainerRule := policy.Rule{
		Name:               "example container runtime rule",
		Collections:        []collection.Collection{runtimeContainerColl},
		AdvancedProtection: true,
		Dns:                policy.Dns{Effect: "disable"},
		Filesystem:         policy.Filesystem{Effect: "alert"},
		Network:            policy.Network{Effect: "alert"},
		Processes:          policy.Processes{Effect: "alert"},
		WildFireAnalysis:   "alert",
	}
	runtimeContainerRules := []policy.Rule{runtimeContainerRule}
	runtimeContainerPolicy := policy.Policy{LearningDisabled: false, Rules: runtimeContainerRules}

	fmt.Printf("\nupdate container runtime policy\n")
	runtimeContainerErr := policy.Update(*client, policy.RuntimeContainerEndpoint, runtimeContainerPolicy)
	if runtimeContainerErr != nil {
		fmt.Printf("\nfailed to update container runtime policy: %v\n", runtimeContainerErr)
	}

	fmt.Printf("\nget container runtime policy:\n")
	retrievedPolicy, runtimeContainerErr = policy.Get(*client, policy.RuntimeContainerEndpoint)
	if runtimeContainerErr != nil {
		fmt.Printf("failed to get container runtime policy: %s\n", runtimeContainerErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	fmt.Printf("\nupdate container runtime policy\n")
	runtimeContainerRule.Name = "name change"
	runtimeContainerPolicy = policy.Policy{LearningDisabled: false, Rules: []policy.Rule{runtimeContainerRule}}
	runtimeContainerErr = policy.Update(*client, policy.RuntimeContainerEndpoint, runtimeContainerPolicy)
	if runtimeContainerErr != nil {
		fmt.Printf("failed to update container runtime policy: %s\n", runtimeContainerErr)
	}

	fmt.Printf("\nget container runtime policy:\n")
	retrievedPolicy, runtimeContainerErr = policy.Get(*client, policy.RuntimeContainerEndpoint)
	if runtimeContainerErr != nil {
		fmt.Printf("failed to get container runtime policy: %s\n", runtimeContainerErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	/*
		HOST RUNTIME
	*/
	runtimeHostColl := collection.Collection{Name: "All"}
	runtimeHostForensic := policy.Forensic{
		ActivitiesDisabled:       false,
		SshdEnabled:              false,
		SudoEnabled:              false,
		ServiceActivitiesEnabled: false,
		DockerEnabled:            false,
		ReadonlyDockerEnabled:    false,
	}
	runtimeHostNetwork := policy.Network{DenyListEffect: "alert", CustomFeed: "alert", IntelligenceFeed: "alert"}
	runtimeHostDNS := policy.Dns{DenyListEffect: "disable", IntelligenceFeed: "disable"}

	runtimeHostDeniedProcesses := policy.DeniedProcesses{Effect: "alert"}
	runtimeHostAntiMalware := policy.AntiMalware{
		DeniedProcesses:            runtimeHostDeniedProcesses,
		CryptoMiner:                "alert",
		ServiceUnknownOriginBinary: "alert",
		UserUnknownOriginBinary:    "alert",
		EncryptedBinaries:          "alert",
		SuspiciousElfHeaders:       "alert",
		TempFsProc:                 "alert",
		ReverseShell:               "alert",
		WebShell:                   "alert",
		ExecutionFlowHijack:        "alert",
		CustomFeed:                 "alert",
		IntelligenceFeed:           "alert",
		WildFireAnalysis:           "alert",
	}

	runtimeHostRule := policy.Rule{Name: "example host runtime rule", Collections: []collection.Collection{runtimeHostColl}, Forensic: runtimeHostForensic, Network: runtimeHostNetwork, Dns: runtimeHostDNS, AntiMalware: runtimeHostAntiMalware}
	runtimeHostRules := []policy.Rule{runtimeHostRule}
	runtimeHostPolicy := policy.Policy{Rules: runtimeHostRules}

	fmt.Printf("\nupdate host runtime policy\n")
	runtimeHostErr := policy.Update(*client, policy.RuntimeHostEndpoint, runtimeHostPolicy)
	if runtimeHostErr != nil {
		fmt.Printf("\nfailed to update host runtime policy: %v\n", runtimeHostErr)
	}

	fmt.Printf("\nget host runtime policy:\n")
	retrievedPolicy, runtimeHostErr = policy.Get(*client, policy.RuntimeHostEndpoint)
	if runtimeHostErr != nil {
		fmt.Printf("failed to get host runtime policy: %s\n", runtimeHostErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	fmt.Printf("\nupdate host runtime policy\n")
	runtimeHostRule.Name = "name change"
	runtimeHostPolicy = policy.Policy{Rules: []policy.Rule{runtimeHostRule}}
	runtimeHostErr = policy.Update(*client, policy.RuntimeHostEndpoint, runtimeHostPolicy)
	if runtimeHostErr != nil {
		fmt.Printf("failed to update host runtime policy: %s\n", runtimeHostErr)
	}

	fmt.Printf("\nget host runtime policy:\n")
	retrievedPolicy, runtimeHostErr = policy.Get(*client, policy.RuntimeHostEndpoint)
	if runtimeHostErr != nil {
		fmt.Printf("failed to get host runtime policy: %s\n", runtimeHostErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	/*
		CI IMAGE VULNERABILITY
	*/
	vulnerabilityCiImageColl := collection.Collection{Name: "All"}
	vulnerabilityCiImageRule := policy.Rule{Name: "example CI image vulnerability rule", Collections: []collection.Collection{vulnerabilityCiImageColl}, Effect: "alert"}
	vulnerabilityCiImagePolicy := policy.Policy{PolicyType: "ciImagesVulnerability", Rules: []policy.Rule{vulnerabilityCiImageRule}}
	fmt.Printf("\nupdate CI image vulnerability policy\n")
	vulnerabilityCiImageErr := policy.Update(*client, policy.VulnerabilityCiImagesEndpoint, vulnerabilityCiImagePolicy)
	if vulnerabilityCiImageErr != nil {
		fmt.Printf("failed to update CI image vulnerability policy: %s\n", vulnerabilityCiImageErr)
	}

	fmt.Printf("\nget CI image vulnerability policy:\n")
	retrievedPolicy, vulnerabilityCiImageErr = policy.Get(*client, policy.VulnerabilityCiImagesEndpoint)
	if vulnerabilityCiImageErr != nil {
		fmt.Printf("failed to get CI image vulnerability policy: %s\n", vulnerabilityCiImageErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	fmt.Printf("\nupdate CI image vulnerability policy\n")
	vulnerabilityCiImageRule.Name = "name change"
	vulnerabilityCiImagePolicy = policy.Policy{PolicyType: "ciImagesVulnerability", Rules: []policy.Rule{vulnerabilityCiImageRule}}
	vulnerabilityCiImageErr = policy.Update(*client, policy.VulnerabilityCiImagesEndpoint, vulnerabilityCiImagePolicy)
	if vulnerabilityCiImageErr != nil {
		fmt.Printf("failed to update CI image vulnerability policy: %s\n", vulnerabilityCiImageErr)
	}

	fmt.Printf("\nget CI image vulnerability policy:\n")
	retrievedPolicy, vulnerabilityCiImageErr = policy.Get(*client, policy.VulnerabilityCiImagesEndpoint)
	if vulnerabilityCiImageErr != nil {
		fmt.Printf("failed to get CI image vulnerability policy: %s\n", vulnerabilityCiImageErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	/*
		HOST VULNERABILITY
	*/
	vulnerabilityHostColl := collection.Collection{Name: "All"}
	vulnerabilityHostRule := policy.Rule{Name: "example host vulnerability rule", Collections: []collection.Collection{vulnerabilityHostColl}, Effect: "alert"}
	vulnerabilityHostPolicy := policy.Policy{PolicyType: "hostVulnerability", Rules: []policy.Rule{vulnerabilityHostRule}}
	fmt.Printf("\nupdate host vulnerability policy\n")
	vulnerabilityHostErr := policy.Update(*client, policy.VulnerabilityHostEndpoint, vulnerabilityHostPolicy)
	if vulnerabilityHostErr != nil {
		fmt.Printf("failed to update host vulnerability policy: %s\n", vulnerabilityHostErr)
	}

	fmt.Printf("\nget host vulnerability policy:\n")
	retrievedPolicy, vulnerabilityHostErr = policy.Get(*client, policy.VulnerabilityHostEndpoint)
	if vulnerabilityHostErr != nil {
		fmt.Printf("failed to get host vulnerability policy: %s\n", vulnerabilityHostErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	fmt.Printf("\nupdate host vulnerability policy\n")
	vulnerabilityHostRule.Name = "name change"
	vulnerabilityHostPolicy = policy.Policy{PolicyType: "hostVulnerability", Rules: []policy.Rule{vulnerabilityHostRule}}
	vulnerabilityHostErr = policy.Update(*client, policy.VulnerabilityHostEndpoint, vulnerabilityHostPolicy)
	if vulnerabilityHostErr != nil {
		fmt.Printf("failed to update host vulnerability policy: %s\n", vulnerabilityHostErr)
	}

	fmt.Printf("\nget host vulnerability policy:\n")
	retrievedPolicy, vulnerabilityHostErr = policy.Get(*client, policy.VulnerabilityHostEndpoint)
	if vulnerabilityHostErr != nil {
		fmt.Printf("failed to get host vulnerability policy: %s\n", vulnerabilityHostErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	/*
		IMAGE VULNERABILITY
	*/
	vulnerabilityimageColl := collection.Collection{Name: "All"}
	vulnerabilityimageRule := policy.Rule{Name: "example image vulnerability rule", Collections: []collection.Collection{vulnerabilityimageColl}, Effect: "alert"}
	vulnerabilityimagePolicy := policy.Policy{PolicyType: "containerVulnerability", Rules: []policy.Rule{vulnerabilityimageRule}}
	fmt.Printf("\nupdate image vulnerability policy\n")
	vulnerabilityimageErr := policy.Update(*client, policy.VulnerabilityImagesEndpoint, vulnerabilityimagePolicy)
	if vulnerabilityimageErr != nil {
		fmt.Printf("failed to update image vulnerability policy: %s\n", vulnerabilityimageErr)
	}

	fmt.Printf("\nget image vulnerability policy:\n")
	retrievedPolicy, vulnerabilityimageErr = policy.Get(*client, policy.VulnerabilityImagesEndpoint)
	if vulnerabilityimageErr != nil {
		fmt.Printf("failed to get image vulnerability policy: %s\n", vulnerabilityimageErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	fmt.Printf("\nupdate image vulnerability policy\n")
	vulnerabilityimageRule.Name = "name change"
	vulnerabilityimagePolicy = policy.Policy{PolicyType: "containerVulnerability", Rules: []policy.Rule{vulnerabilityimageRule}}
	vulnerabilityimageErr = policy.Update(*client, policy.VulnerabilityImagesEndpoint, vulnerabilityimagePolicy)
	if vulnerabilityimageErr != nil {
		fmt.Printf("failed to update image vulnerability policy: %s\n", vulnerabilityimageErr)
	}

	fmt.Printf("\nget image vulnerability policy:\n")
	retrievedPolicy, vulnerabilityimageErr = policy.Get(*client, policy.VulnerabilityImagesEndpoint)
	if vulnerabilityimageErr != nil {
		fmt.Printf("failed to get image vulnerability policy: %s\n", vulnerabilityimageErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	/*
		REGISTRY SETTINGS
	*/
	registrySpec := registry.Specification{
		Version:     "2",
		Registry:    "",
		Os:          "linux",
		Cap:         5,
		Scanners:    2,
		Repository:  "library/ubuntu",
		Tag:         "20.04",
		Collections: []string{"All"},
	}
	reg := registry.Registry{Specifications: []registry.Specification{registrySpec}}

	fmt.Printf("\ncreate registry settings:\n")
	registryErr := registry.Update(*client, reg)
	if registryErr != nil {
		fmt.Printf("failed to create registry settings: %s\n", registryErr)
	}

	fmt.Printf("\nget registry settings:\n")
	retrievedRegistry, registryErr := registry.Get(*client)
	if registryErr != nil {
		fmt.Printf("failed to get registry settings: %s\n", registryErr)
	}
	fmt.Printf("* %v\n", retrievedRegistry)

	fmt.Printf("\nupdate registry settings\n")
	registrySpec.Tag = "21.04"
	reg = registry.Registry{Specifications: []registry.Specification{registrySpec}}
	registryErr = registry.Update(*client, reg)
	if registryErr != nil {
		fmt.Printf("failed to update registry settings: %s\n", registryErr)
	}

	fmt.Printf("\nget registry settings:\n")
	retrievedRegistry, registryErr = registry.Get(*client)
	if registryErr != nil {
		fmt.Printf("failed to get registry settings: %s\n", registryErr)
	}
	fmt.Printf("* %v\n", retrievedRegistry)
}
