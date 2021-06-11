package main

import (
    "fmt"
    "github.com/paloaltonetworks/prisma-cloud-compute-go"
    "github.com/paloaltonetworks/prisma-cloud-compute-go/collection"
)

func main() {
    client := &prismacloudcompute.Client{}
    
    if err := client.Initialize("creds.json"); err != nil {
        fmt.Printf("Failed to connect: %s\n", err)
    }
    
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
}
