package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	pcc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collections"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policies"
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
	newColl := collections.Collection{Name: "My Collection"}
	err = collections.Create(*client, newColl)
	if err != nil {
		fmt.Printf("Failed to create collections: %s\n", err)
		return
	}

	fmt.Printf("\nlist collections:\n")
	colls, err := collections.List(*client)
	if err != nil {
		fmt.Printf("Failed to list collections: %s\n", err)
	}

	for _, v := range colls {
		fmt.Printf("* %s %s\n", v.Name, v.Color)
	}

	fmt.Printf("\nget collection:\n")
	coll, err := collections.Get(*client, "My Collection")
	if err != nil {
		fmt.Printf("Failed to get collection: %s\n", err)
	}
	fmt.Printf("* %s %s\n", coll.Name, coll.Color)

	fmt.Printf("\nupdate collection\n")
	existingColl := collections.Collection{Name: "My Collection", Color: "#FFFFFF"}
	err = collections.Update(*client, existingColl)
	if err != nil {
		fmt.Printf("Failed to update collection: %s\n", err)
	}

	fmt.Printf("\nlist collections:\n")
	colls, err = collections.List(*client)
	if err != nil {
		fmt.Printf("Failed to list collections: %s\n", err)
	}

	for _, v := range colls {
		fmt.Printf("* %s %s\n", v.Name, v.Color)
	}

	fmt.Printf("\ndelete collection\n")
	err = collections.Delete(*client, "My Collection")
	if err != nil {
		fmt.Printf("Failed to delete collection: %s\n", err)
	}

	fmt.Printf("\nlist collections:\n")
	colls, err = collections.List(*client)
	if err != nil {
		fmt.Printf("Failed to list collections: %s\n", err)
	}

	for _, v := range colls {
		fmt.Printf("* %s %s\n", v.Name, v.Color)
	}

	/*
		CI IMAGE COMPLIANCE
	*/
	complianceCiImageColl := collections.Collection{Name: "All"}
	complianceCiImageVuln1 := policies.Vulnerability{Id: 41, Block: false}
	complianceCiImageVuln2 := policies.Vulnerability{Id: 422, Block: true}
	complianceCiImageCondition := policies.Condition{Vulnerabilities: []policies.Vulnerability{complianceCiImageVuln1, complianceCiImageVuln2}}
	complianceCiImageRule := policies.Rule{Name: "example ci image compliance rule", Effect: "alert, block", Collections: []collections.Collection{complianceCiImageColl}, Condition: complianceCiImageCondition}
	complianceCiImageRules := []policies.Rule{complianceCiImageRule}
	complianceCiImagePolicy := policies.Policy{PolicyType: "ciImagesCompliance", Rules: complianceCiImageRules}

	fmt.Printf("\nupdate CI image compliance policy\n")
	complianceCiImageErr := policies.Update(*client, policies.ComplianceCiImagesEndpoint, complianceCiImagePolicy)
	if complianceCiImageErr != nil {
		fmt.Printf("\nFailed to update CI image compliance policy: %v\n", complianceCiImageErr)
	}

	fmt.Printf("\nget CI image compliance policy:\n")
	retrievedPolicy, complianceCiImageErr := policies.Get(*client, policies.ComplianceCiImagesEndpoint)
	if complianceCiImageErr != nil {
		fmt.Printf("failed to get CI image compliance policy: %s\n", complianceCiImageErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	fmt.Printf("\nupdate CI image compliance policy\n")
	complianceCiImageRule.Name = "name change"
	complianceCiImagePolicy = policies.Policy{PolicyType: "ciImagesCompliance", Rules: []policies.Rule{complianceCiImageRule}}
	complianceCiImageErr = policies.Update(*client, policies.ComplianceCiImagesEndpoint, complianceCiImagePolicy)
	if complianceCiImageErr != nil {
		fmt.Printf("failed to update CI image compliance policy: %s\n", complianceCiImageErr)
	}

	fmt.Printf("\nget CI image compliance policy:\n")
	retrievedPolicy, complianceCiImageErr = policies.Get(*client, policies.ComplianceCiImagesEndpoint)
	if complianceCiImageErr != nil {
		fmt.Printf("failed to get CI image compliance policy: %s\n", complianceCiImageErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	/*
		CONTAINER COMPLIANCE
	*/
	complianceContainerColl := collections.Collection{Name: "All"}
	complianceContainerVuln1 := policies.Vulnerability{Id: 41, Block: false}
	complianceContainerVuln2 := policies.Vulnerability{Id: 422, Block: true}
	complianceContainerCondition := policies.Condition{Vulnerabilities: []policies.Vulnerability{complianceContainerVuln1, complianceContainerVuln2}}
	complianceContainerRule := policies.Rule{Name: "example container compliance rule", Effect: "alert, block", Collections: []collections.Collection{complianceContainerColl}, Condition: complianceContainerCondition}
	complianceContainerRules := []policies.Rule{complianceContainerRule}
	complianceContainerPolicy := policies.Policy{PolicyType: "containerCompliance", Rules: complianceContainerRules}

	fmt.Printf("\nupdate container compliance policy\n")
	complianceContainerErr := policies.Update(*client, policies.ComplianceContainerEndpoint, complianceContainerPolicy)
	if complianceContainerErr != nil {
		fmt.Printf("\nFailed to update container compliance policy: %v\n", complianceContainerErr)
	}

	fmt.Printf("\nget container compliance policy:\n")
	retrievedPolicy, complianceContainerErr = policies.Get(*client, policies.ComplianceContainerEndpoint)
	if complianceContainerErr != nil {
		fmt.Printf("failed to get container compliance policy: %s\n", complianceContainerErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	fmt.Printf("\nupdate container compliance policy\n")
	complianceContainerRule.Name = "name change"
	complianceContainerPolicy = policies.Policy{PolicyType: "containerCompliance", Rules: []policies.Rule{complianceContainerRule}}
	complianceContainerErr = policies.Update(*client, policies.ComplianceContainerEndpoint, complianceContainerPolicy)
	if complianceContainerErr != nil {
		fmt.Printf("failed to update container compliance policy: %s\n", complianceContainerErr)
	}

	fmt.Printf("\nget container compliance policy:\n")
	retrievedPolicy, complianceContainerErr = policies.Get(*client, policies.ComplianceContainerEndpoint)
	if complianceContainerErr != nil {
		fmt.Printf("failed to get container compliance policy: %s\n", complianceContainerErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	/*
		HOST COMPLIANCE
	*/
	complianceHostColl := collections.Collection{Name: "All"}
	complianceHostVuln1 := policies.Vulnerability{Id: 11, Block: false}
	complianceHostVuln2 := policies.Vulnerability{Id: 112, Block: true}
	complianceHostCondition := policies.Condition{Vulnerabilities: []policies.Vulnerability{complianceHostVuln1, complianceHostVuln2}}
	complianceHostRule := policies.Rule{Name: "example host compliance rule", Effect: "alert, block", Collections: []collections.Collection{complianceHostColl}, Condition: complianceHostCondition}
	complianceHostRules := []policies.Rule{complianceHostRule}
	complianceHostPolicy := policies.Policy{PolicyType: "hostCompliance", Rules: complianceHostRules}

	fmt.Printf("\nupdate host compliance policy\n")
	complianceHostErr := policies.Update(*client, policies.ComplianceHostEndpoint, complianceHostPolicy)
	if complianceHostErr != nil {
		fmt.Printf("\nFailed to update host compliance policy: %v\n", complianceHostErr)
	}

	fmt.Printf("\nget host compliance policy:\n")
	retrievedPolicy, complianceHostErr = policies.Get(*client, policies.ComplianceHostEndpoint)
	if complianceHostErr != nil {
		fmt.Printf("failed to get host compliance policy: %s\n", complianceHostErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	fmt.Printf("\nupdate host compliance policy\n")
	complianceHostRule.Name = "name change"
	complianceHostPolicy = policies.Policy{PolicyType: "hostCompliance", Rules: []policies.Rule{complianceHostRule}}
	complianceHostErr = policies.Update(*client, policies.ComplianceHostEndpoint, complianceHostPolicy)
	if complianceHostErr != nil {
		fmt.Printf("failed to update host compliance policy: %s\n", complianceHostErr)
	}

	fmt.Printf("\nget host compliance policy:\n")
	retrievedPolicy, complianceHostErr = policies.Get(*client, policies.ComplianceHostEndpoint)
	if complianceHostErr != nil {
		fmt.Printf("failed to get host compliance policy: %s\n", complianceHostErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	/*
		CONTAINER RUNTIME
	*/
	runtimeContainerColl := collections.Collection{Name: "All"}
	runtimeContainerRule := policies.Rule{
		Name:               "example container runtime rule",
		Collections:        []collections.Collection{runtimeContainerColl},
		AdvancedProtection: true,
		Dns:                policies.Dns{Effect: "disable"},
		Filesystem:         policies.Filesystem{Effect: "alert"},
		Network:            policies.Network{Effect: "alert"},
		Processes:          policies.Processes{Effect: "alert"},
		WildFireAnalysis:   "alert",
	}
	runtimeContainerRules := []policies.Rule{runtimeContainerRule}
	runtimeContainerPolicy := policies.Policy{LearningDisabled: false, Rules: runtimeContainerRules}

	fmt.Printf("\nupdate container runtime policy\n")
	runtimeContainerErr := policies.Update(*client, policies.RuntimeContainerEndpoint, runtimeContainerPolicy)
	if runtimeContainerErr != nil {
		fmt.Printf("\nfailed to update container runtime policy: %v\n", runtimeContainerErr)
	}

	fmt.Printf("\nget container runtime policy:\n")
	retrievedPolicy, runtimeContainerErr = policies.Get(*client, policies.RuntimeContainerEndpoint)
	if runtimeContainerErr != nil {
		fmt.Printf("failed to get container runtime policy: %s\n", runtimeContainerErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	fmt.Printf("\nupdate container runtime policy\n")
	runtimeContainerRule.Name = "name change"
	runtimeContainerPolicy = policies.Policy{LearningDisabled: false, Rules: []policies.Rule{runtimeContainerRule}}
	runtimeContainerErr = policies.Update(*client, policies.RuntimeContainerEndpoint, runtimeContainerPolicy)
	if runtimeContainerErr != nil {
		fmt.Printf("failed to update container runtime policy: %s\n", runtimeContainerErr)
	}

	fmt.Printf("\nget container runtime policy:\n")
	retrievedPolicy, runtimeContainerErr = policies.Get(*client, policies.RuntimeContainerEndpoint)
	if runtimeContainerErr != nil {
		fmt.Printf("failed to get container runtime policy: %s\n", runtimeContainerErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	/*
		HOST RUNTIME
	*/
	runtimeHostColl := collections.Collection{Name: "All"}
	runtimeHostForensic := policies.Forensic{
		ActivitiesDisabled:       false,
		SshdEnabled:              false,
		SudoEnabled:              false,
		ServiceActivitiesEnabled: false,
		DockerEnabled:            false,
		ReadonlyDockerEnabled:    false,
	}
	runtimeHostNetwork := policies.Network{DenyListEffect: "alert", CustomFeed: "alert", IntelligenceFeed: "alert"}
	runtimeHostDNS := policies.Dns{DenyListEffect: "disable", IntelligenceFeed: "disable"}

	runtimeHostDeniedProcesses := policies.DeniedProcesses{Effect: "alert"}
	runtimeHostAntiMalware := policies.AntiMalware{
		DeniedProcesses:            runtimeHostDeniedProcesses,
		CryptoMiner:                "alert",
		ServiceUnknownOriginBinary: "alert",
		UserUnknownOriginBinary:    "alert",
		EncryptedBinaries:          "alert",
		SuspiciousELFHeaders:       "alert",
		TempFSProc:                 "alert",
		ReverseShell:               "alert",
		WebShell:                   "alert",
		ExecutionFlowHijack:        "alert",
		CustomFeed:                 "alert",
		IntelligenceFeed:           "alert",
		WildFireAnalysis:           "alert",
	}

	runtimeHostRule := policies.Rule{Name: "example host runtime rule", Collections: []collections.Collection{runtimeHostColl}, Forensic: runtimeHostForensic, Network: runtimeHostNetwork, Dns: runtimeHostDNS, AntiMalware: runtimeHostAntiMalware}
	runtimeHostRules := []policies.Rule{runtimeHostRule}
	runtimeHostPolicy := policies.Policy{Rules: runtimeHostRules}

	fmt.Printf("\nupdate host runtime policy\n")
	runtimeHostErr := policies.Update(*client, policies.RuntimeHostEndpoint, runtimeHostPolicy)
	if runtimeHostErr != nil {
		fmt.Printf("\nfailed to update host runtime policy: %v\n", runtimeHostErr)
	}

	fmt.Printf("\nget host runtime policy:\n")
	retrievedPolicy, runtimeHostErr = policies.Get(*client, policies.RuntimeHostEndpoint)
	if runtimeHostErr != nil {
		fmt.Printf("failed to get host runtime policy: %s\n", runtimeHostErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	fmt.Printf("\nupdate host runtime policy\n")
	runtimeHostRule.Name = "name change"
	runtimeHostPolicy = policies.Policy{Rules: []policies.Rule{runtimeHostRule}}
	runtimeHostErr = policies.Update(*client, policies.RuntimeHostEndpoint, runtimeHostPolicy)
	if runtimeHostErr != nil {
		fmt.Printf("failed to update host runtime policy: %s\n", runtimeHostErr)
	}

	fmt.Printf("\nget host runtime policy:\n")
	retrievedPolicy, runtimeHostErr = policies.Get(*client, policies.RuntimeHostEndpoint)
	if runtimeHostErr != nil {
		fmt.Printf("failed to get host runtime policy: %s\n", runtimeHostErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	/*
		CI IMAGE VULNERABILITY
	*/
	vulnerabilityCiImageColl := collections.Collection{Name: "All"}
	vulnerabilityCiImageRule := policies.Rule{Name: "example CI image vulnerability rule", Collections: []collections.Collection{vulnerabilityCiImageColl}, Effect: "alert"}
	vulnerabilityCiImagePolicy := policies.Policy{PolicyType: "ciImagesVulnerability", Rules: []policies.Rule{vulnerabilityCiImageRule}}
	fmt.Printf("\nupdate CI image vulnerability policy\n")
	vulnerabilityCiImageErr := policies.Update(*client, policies.VulnerabilityCiImagesEndpoint, vulnerabilityCiImagePolicy)
	if vulnerabilityCiImageErr != nil {
		fmt.Printf("failed to update CI image vulnerability policy: %s\n", vulnerabilityCiImageErr)
	}

	fmt.Printf("\nget CI image vulnerability policy:\n")
	retrievedPolicy, vulnerabilityCiImageErr = policies.Get(*client, policies.VulnerabilityCiImagesEndpoint)
	if vulnerabilityCiImageErr != nil {
		fmt.Printf("failed to get CI image vulnerability policy: %s\n", vulnerabilityCiImageErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	fmt.Printf("\nupdate CI image vulnerability policy\n")
	vulnerabilityCiImageRule.Name = "name change"
	vulnerabilityCiImagePolicy = policies.Policy{PolicyType: "ciImagesVulnerability", Rules: []policies.Rule{vulnerabilityCiImageRule}}
	vulnerabilityCiImageErr = policies.Update(*client, policies.VulnerabilityCiImagesEndpoint, vulnerabilityCiImagePolicy)
	if vulnerabilityCiImageErr != nil {
		fmt.Printf("failed to update CI image vulnerability policy: %s\n", vulnerabilityCiImageErr)
	}

	fmt.Printf("\nget CI image vulnerability policy:\n")
	retrievedPolicy, vulnerabilityCiImageErr = policies.Get(*client, policies.VulnerabilityCiImagesEndpoint)
	if vulnerabilityCiImageErr != nil {
		fmt.Printf("failed to get CI image vulnerability policy: %s\n", vulnerabilityCiImageErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	/*
		HOST VULNERABILITY
	*/
	vulnerabilityHostColl := collections.Collection{Name: "All"}
	vulnerabilityHostRule := policies.Rule{Name: "example host vulnerability rule", Collections: []collections.Collection{vulnerabilityHostColl}, Effect: "alert"}
	vulnerabilityHostPolicy := policies.Policy{PolicyType: "hostVulnerability", Rules: []policies.Rule{vulnerabilityHostRule}}
	fmt.Printf("\nupdate host vulnerability policy\n")
	vulnerabilityHostErr := policies.Update(*client, policies.VulnerabilityHostEndpoint, vulnerabilityHostPolicy)
	if vulnerabilityHostErr != nil {
		fmt.Printf("failed to update host vulnerability policy: %s\n", vulnerabilityHostErr)
	}

	fmt.Printf("\nget host vulnerability policy:\n")
	retrievedPolicy, vulnerabilityHostErr = policies.Get(*client, policies.VulnerabilityHostEndpoint)
	if vulnerabilityHostErr != nil {
		fmt.Printf("failed to get host vulnerability policy: %s\n", vulnerabilityHostErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	fmt.Printf("\nupdate host vulnerability policy\n")
	vulnerabilityHostRule.Name = "name change"
	vulnerabilityHostPolicy = policies.Policy{PolicyType: "hostVulnerability", Rules: []policies.Rule{vulnerabilityHostRule}}
	vulnerabilityHostErr = policies.Update(*client, policies.VulnerabilityHostEndpoint, vulnerabilityHostPolicy)
	if vulnerabilityHostErr != nil {
		fmt.Printf("failed to update host vulnerability policy: %s\n", vulnerabilityHostErr)
	}

	fmt.Printf("\nget host vulnerability policy:\n")
	retrievedPolicy, vulnerabilityHostErr = policies.Get(*client, policies.VulnerabilityHostEndpoint)
	if vulnerabilityHostErr != nil {
		fmt.Printf("failed to get host vulnerability policy: %s\n", vulnerabilityHostErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	/*
		IMAGE VULNERABILITY
	*/
	vulnerabilityimageColl := collections.Collection{Name: "All"}
	vulnerabilityimageRule := policies.Rule{Name: "example image vulnerability rule", Collections: []collections.Collection{vulnerabilityimageColl}, Effect: "alert"}
	vulnerabilityimagePolicy := policies.Policy{PolicyType: "containerVulnerability", Rules: []policies.Rule{vulnerabilityimageRule}}
	fmt.Printf("\nupdate image vulnerability policy\n")
	vulnerabilityimageErr := policies.Update(*client, policies.VulnerabilityImagesEndpoint, vulnerabilityimagePolicy)
	if vulnerabilityimageErr != nil {
		fmt.Printf("failed to update image vulnerability policy: %s\n", vulnerabilityimageErr)
	}

	fmt.Printf("\nget image vulnerability policy:\n")
	retrievedPolicy, vulnerabilityimageErr = policies.Get(*client, policies.VulnerabilityImagesEndpoint)
	if vulnerabilityimageErr != nil {
		fmt.Printf("failed to get image vulnerability policy: %s\n", vulnerabilityimageErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	fmt.Printf("\nupdate image vulnerability policy\n")
	vulnerabilityimageRule.Name = "name change"
	vulnerabilityimagePolicy = policies.Policy{PolicyType: "containerVulnerability", Rules: []policies.Rule{vulnerabilityimageRule}}
	vulnerabilityimageErr = policies.Update(*client, policies.VulnerabilityImagesEndpoint, vulnerabilityimagePolicy)
	if vulnerabilityimageErr != nil {
		fmt.Printf("failed to update image vulnerability policy: %s\n", vulnerabilityimageErr)
	}

	fmt.Printf("\nget image vulnerability policy:\n")
	retrievedPolicy, vulnerabilityimageErr = policies.Get(*client, policies.VulnerabilityImagesEndpoint)
	if vulnerabilityimageErr != nil {
		fmt.Printf("failed to get image vulnerability policy: %s\n", vulnerabilityimageErr)
	}
	fmt.Printf("* %s %v\n", retrievedPolicy.PolicyId, retrievedPolicy.Rules)

	/*
		REGISTRY SETTINGS
	*/
	fmt.Printf("\nupdate registry settings\n")
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
	registryErr := registry.Update(*client, reg)
	if registryErr != nil {
		fmt.Printf("failed to update registry settings: %s\n", registryErr)
	}

	fmt.Printf("\nget registry settings:\n")
	retrievedRegistry, registryErr := registry.Get(*client)
	if registryErr != nil {
		fmt.Printf("failed to get registry settings: %s\n", registryErr)
	}
	fmt.Printf("* %v\n", retrievedRegistry)

	fmt.Printf("\nupdate registry settings\n")
	registrySpec = registry.Specification{
		Version:     "2",
		Registry:    "",
		Os:          "linux",
		Cap:         5,
		Scanners:    2,
		Repository:  "library/ubuntu",
		Tag:         "21.04",
		Collections: []string{"All"},
	}
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
