/*
Package prismacloud is an SDK meant to assist in interacting with the Palo
Alto Networks Prisma Cloud API.

To connect:

Create a JSON for your connection information

{
	"url":"localhost",
	"username":"my-username",
	"password":"my-password",
	"port":8083,
	"skip_ssl_cert_verification":true
}

Create a client connetion with the desired params and then
initialize the connection:

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

In most cases the struct and types match what the Prisma Cloud API
specifies, so you may find it useful to refer to the Prisma Cloud API
for further information: https://prisma.pan.dev/api/cloud/
*/
package prismacloudcompute
