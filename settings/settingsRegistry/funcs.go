package settingsRegistry

import (
	"fmt"
	"strings"

//	pc "github.com/paloaltonetworks/prisma-cloud-compute-go"
	pc "prisma-cloud-compute-go"
)

// List returns a list of all registries
/*func List(c pc.PrismaCloudClient) ([]Registry, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	var ans []Registry
	if _, err := c.Communicate("GET", Suffix, nil, nil, &ans); err != nil {
		return nil, err
	}

	return ans, nil
}*/

/*
// Identify returns the id for the given registry name.
func Identify(c pc.PrismaCloudClient, id string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s name:%s", singular, name)

	ans, err := List(c, map[string]string{"registry.name": name})
	if err != nil {
		return "", err
	}

	switch len(ans) {
	case 0:
		return "", pc.ObjectNotFoundError
	case 1:
		return ans[0].RegistryId, nil
	}

	return "", fmt.Errorf("Got %d results back not 1", len(ans))
}
*/

// Get returns the registry.
func Get(c pc.PrismaCloudClient) (Registry, error) {
	c.Log(pc.LogAction, "(get) %s name:%s", singular, name)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, name)

	var ans Registry
	_, err := c.Communicate("GET", path, nil, nil, &ans)

	return ans, err
}


// Create adds a new registry.
func Create(c pc.PrismaCloudClient, registry Registry) error {
	return createUpdate(false, c, registry)
}

// Update modifies the existing registry.
func Update(c pc.PrismaCloudClient, registry Registry) error {
	return createUpdate(true, c, registry)
}

// Delete removes a registry using its name.
func Delete(c pc.PrismaCloudClient, name string) error {
	c.Log(pc.LogAction, "(delete) %s name:%s", singular, name)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, name)
	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, registry Registry) error {
	var (
		logMsg strings.Builder
		method string
	)

	logMsg.Grow(30)
	logMsg.WriteString("(")
	if exists {
		logMsg.WriteString("update")
		method = "PUT"
	} else {
		logMsg.WriteString("create")
		method = "POST"
	}
	logMsg.WriteString(") ")

	logMsg.WriteString(singular)
	if exists {
		fmt.Fprintf(&logMsg, ":%s", registry.Name)
	}

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	if exists {
		path = append(path, registry.Name)
	}

	_, err := c.Communicate(method, path, nil, registry, nil)
	return err
}
