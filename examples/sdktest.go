package main

import (
	"fmt"
	"github.com/paloaltonetworks/prisma-cloud-compute-go"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/collection"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy"
	//    "github.com/paloaltonetworks/prisma-cloud-compute-go/policy/policyRuntimeContainer"
	//    "github.com/paloaltonetworks/prisma-cloud-compute-go/policy/policyVulnerabilityImages"
	//    "github.com/paloaltonetworks/prisma-cloud-compute-go/policy/policyComplianceContainer"
	//    "github.com/paloaltonetworks/prisma-cloud-compute-go/policy/policyVulnerabilityCiImages"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/policy/policyComplianceCiImages"
)

func main() {
	client := &prismacloudcompute.Client{}

	if err := client.Initialize("creds.json"); err != nil {
		fmt.Printf("Failed to connect: %s\n", err)
	}

	/*
	   fmt.Printf("\nCREATE new Collection:\n")
	   coll := collection.Collection{Name: "My Collection"}
	   err := collection.Create(client, coll)
	   if err != nil {
	       fmt.Printf("Failed to get collections: %s\n", err)
	   }

	   fmt.Printf("GET Collections List:\n")
	   listing, err := collection.List(client)
	   if err != nil {
	       fmt.Printf("Failed to get collections: %s\n", err)
	   }

	   for _, elm := range listing {
	       fmt.Printf("* %s   %s\n", elm.Name, elm.Color)
	   }

	   fmt.Printf("\nGET single Collection:\n")
	   singleColl, err := collection.Get(client, "My Collection")
	   if err != nil {
	       fmt.Printf("Failed to get single collection: %s\n", err)
	   }
	   fmt.Printf("* %s   %s\n", singleColl.Name, singleColl.Color)

	   fmt.Printf("\nUPDATE new Collection:\n")
	   coll = collection.Collection{Name: "My Collection", Color: "#FFFFFF"}
	   err = collection.Update(client, coll)
	   if err != nil {
	       fmt.Printf("Failed to get collections: %s\n", err)
	   }

	   fmt.Printf("GET Collections List:\n")
	   listing, err = collection.List(client)
	   if err != nil {
	       fmt.Printf("Failed to get collections: %s\n", err)
	   }

	   for _, elm := range listing {
	       fmt.Printf("* %s   %s\n", elm.Name, elm.Color)
	   }

	   fmt.Printf("\nDELETE added Collection:\n")
	   err = collection.Delete(client, "My Collection")
	   if err != nil {
	       fmt.Printf("Failed to get collections: %s\n", err)
	   }

	   fmt.Printf("GET Collections List:\n")
	   listing, err = collection.List(client)
	   if err != nil {
	       fmt.Printf("Failed to get collections: %s\n", err)
	   }

	   for _, elm := range listing {
	       fmt.Printf("* %s\n", elm.Name)
	   }
	*/

	/*    fmt.Printf("\nGET Policy Runtime Container:\n")
	      singlePol, err := policyRuntimeContainer.Get(client)
	      if err != nil {
	          fmt.Printf("Failed to get single policy: %s\n", err)
	      }
	      fmt.Printf("* %s   %s\n", singlePol.PolicyId, singlePol.LearningDisabled)

	      fmt.Printf("\nUPDATE new Policy:\n")
	      pol := policyRuntimeContainer.Policy{PolicyId: "My Policy", LearningDisabled: false}
	      err = policyRuntimeContainer.Update(client, pol)
	      if err != nil {
	          fmt.Printf("Failed to get policies: %s\n", err)
	      }

	      fmt.Printf("\nGET Policy:\n")
	      singlePol, err = policyRuntimeContainer.Get(client)
	      if err != nil {
	          fmt.Printf("Failed to get single policy: %s\n", err)
	      }
	      fmt.Printf("* %s   %s\n", singlePol.PolicyId, singlePol.LearningDisabled)
	*/

	/*    fmt.Printf("\nGET Policy Vulnerability Images:\n")
	      singlePol, err := policyVulnerabilityImages.Get(client)
	      if err != nil {
	          fmt.Printf("Failed to get single policy: %s\n", err)
	      }
	      fmt.Printf("* %s   %s\n", singlePol.PolicyId, singlePol.PolicyType)

	      fmt.Printf("\nUPDATE new Policy:\n")
	      pol := policyVulnerabilityImages.Policy{PolicyId: "My Policy", PolicyType: policy.PolicyTypeContainerVulnerability}
	      err = policyVulnerabilityImages.Update(client, pol)
	      if err != nil {
	          fmt.Printf("Failed to get policies: %s\n", err)
	      }

	      fmt.Printf("\nGET Policy:\n")
	      singlePol, err = policyVulnerabilityImages.Get(client)
	      if err != nil {
	          fmt.Printf("Failed to get single policy: %s\n", err)
	      }
	      fmt.Printf("* %s   %s\n", singlePol.PolicyId, singlePol.PolicyType)

	  }*/

	/*    fmt.Printf("\nCREATE new Policy Compliance Container:\n")
	      coll := collection.Collection{Name: "All"}
	      vuln := policy.Vulnerability{Id: 531, Block: false}
	      cond := policy.Condition{Vulnerabilities: []policy.Vulnerability{vuln}}
	      rule := policy.Rule{Name: "my-rule", Effect: policy.EffectAlert, Collections: []collection.Collection{coll}, Condition: cond}
	      rules := []policy.Rule{rule}
	      pol := policyComplianceContainer.Policy{PolicyType: policy.PolicyTypeContainerCompliance, Rules: rules}
	      err := policyComplianceContainer.Create(client, pol)
	      if err != nil {
	          fmt.Printf("Failed to get policies: %s\n", err)
	      }

	      fmt.Printf("\nGET Policy Compliance Container:\n")
	      singlePol, err := policyComplianceContainer.Get(client)
	      if err != nil {
	          fmt.Printf("Failed to get single policy: %s\n", err)
	      }
	      fmt.Printf("* %s   %s\n", singlePol.PolicyId, singlePol.PolicyType)

	      fmt.Printf("\nUPDATE new Policy:\n")
	      pol.Rules[0].Condition.Vulnerabilities[0].Id = 41
	      err = policyComplianceContainer.Update(client, pol)
	      if err != nil {
	          fmt.Printf("Failed to get policies: %s\n", err)
	      }

	      fmt.Printf("\nGET Policy:\n")
	      singlePol, err = policyComplianceContainer.Get(client)
	      if err != nil {
	          fmt.Printf("Failed to get single policy: %s\n", err)
	      }
	      fmt.Printf("* %s   %s\n", singlePol.PolicyId, singlePol.PolicyType)
	*/
	/*
	   policyName := "example ci image vulnerability rule"
	   coll := collection.Collection{Name: "All"}
	   alert := policy.Threshold{Value: 1, Disabled: false}
	   block := policy.Threshold{Value: 0, Enabled: false}
	   rule := policy.Rule{Name: policyName, Effect: policy.EffectAlert, Collections: []collection.Collection{coll}, AlertThreshold: alert, BlockThreshold: block}
	   rules := []policy.Rule{rule}
	   pol := policyVulnerabilityCiImages.Policy{PolicyType: policy.PolicyTypeCiImagesVulnerability, Rules: rules}
	   err := policyVulnerabilityCiImages.Create(client, pol)
	   if err != nil {
	       fmt.Printf("Failed to get policies: %s\n", err)
	   }

	   fmt.Printf("\nGET Policy Vulnerability CI Images:\n")
	   singlePol, err := policyVulnerabilityCiImages.Get(client)
	   if err != nil {
	       fmt.Printf("Failed to get single policy: %s\n", err)
	   }
	   fmt.Printf("* %s   %s\n", singlePol.PolicyId, singlePol.Rules[0].Name)

	   fmt.Printf("\nUPDATE new Policy:\n")
	   rule.Name = "name change"
	   pol = policyVulnerabilityCiImages.Policy{PolicyType: policy.PolicyTypeCiImagesVulnerability, Rules: []policy.Rule{rule}}
	   err = policyVulnerabilityCiImages.Update(client, pol)
	   if err != nil {
	       fmt.Printf("Failed to get policies: %s\n", err)
	   }

	   fmt.Printf("\nGET Policy:\n")
	   singlePol, err = policyVulnerabilityCiImages.Get(client)
	   if err != nil {
	       fmt.Printf("Failed to get single policy: %s\n", err)
	   }
	   fmt.Printf("* %s   %s\n", singlePol.PolicyId, singlePol.Rules)
	*/

	coll := collection.Collection{Name: "All"}
	vuln1 := policy.Vulnerability{Id: 41, Block: false}
	vuln2 := policy.Vulnerability{Id: 422, Block: false}
	vuln3 := policy.Vulnerability{Id: 424, Block: false}
	vuln4 := policy.Vulnerability{Id: 425, Block: false}
	vuln5 := policy.Vulnerability{Id: 426, Block: false}
	vuln6 := policy.Vulnerability{Id: 448, Block: false}
	vuln7 := policy.Vulnerability{Id: 5041, Block: false}
	cond := policy.Condition{Vulnerabilities: []policy.Vulnerability{vuln1, vuln2, vuln3, vuln4, vuln5, vuln6, vuln7}}
	rule := policy.Rule{Name: "example ci image compliance rule", Effect: policy.EffectAlert, Collections: []collection.Collection{coll}, Condition: cond}
	rules := []policy.Rule{rule}
	pol := policyComplianceCiImages.Policy{PolicyType: policy.PolicyTypeCiImagesCompliance, Rules: rules}
	err := policyComplianceCiImages.Create(client, pol)
	if err != nil {
		fmt.Printf("Failed to get policies: %s\n", err)
	}

	fmt.Printf("\nGET Policy Compliance CI Images:\n")
	singlePol, err := policyComplianceCiImages.Get(client)
	if err != nil {
		fmt.Printf("Failed to get single policy: %s\n", err)
	}
	fmt.Printf("* %s   %s\n", singlePol.PolicyId, singlePol.Rules[0].Name)

	fmt.Printf("\nUPDATE new Policy:\n")
	rule.Name = "name change"
	pol = policyComplianceCiImages.Policy{PolicyType: policy.PolicyTypeCiImagesCompliance, Rules: []policy.Rule{rule}}
	err = policyComplianceCiImages.Update(client, pol)
	if err != nil {
		fmt.Printf("Failed to get policies: %s\n", err)
	}

	fmt.Printf("\nGET Policy:\n")
	singlePol, err = policyComplianceCiImages.Get(client)
	if err != nil {
		fmt.Printf("Failed to get single policy: %s\n", err)
	}
	fmt.Printf("* %s   %s\n", singlePol.PolicyId, singlePol.Rules)
}
