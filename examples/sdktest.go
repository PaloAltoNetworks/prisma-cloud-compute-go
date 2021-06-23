package main

import (
    "fmt"
    "github.com/paloaltonetworks/prisma-cloud-compute-go"
//    "github.com/paloaltonetworks/prisma-cloud-compute-go/collection"
//    "github.com/paloaltonetworks/prisma-cloud-compute-go/policy/policyRuntimeContainer"
    "github.com/paloaltonetworks/prisma-cloud-compute-go/policy/policyVulnerabilityImages"
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


    fmt.Printf("\nGET Policy Vulnerability Images:\n")
    singlePol, err := policyVulnerabilityImages.Get(client)
    if err != nil {
        fmt.Printf("Failed to get single policy: %s\n", err)
    }
    fmt.Printf("* %s   %s\n", singlePol.PolicyId, singlePol.PolicyType)
    
    fmt.Printf("\nUPDATE new Policy:\n")
    pol := policyVulnerabilityImages.Policy{PolicyId: "My Policy", PolicyType: policyVulnerabilityImages.PolicyTypeContainerVulnerability}	
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

}
