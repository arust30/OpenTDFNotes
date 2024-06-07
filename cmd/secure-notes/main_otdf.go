package main

import (
	"fmt"
	"github.com/opentdf/otdfctl/pkg/client"
)

func main() {
	// Create a new OpenTDF client
	c, err := client.NewClient()
	if err != nil {
		fmt.Println("Error creating client:", err)
		return
	}

	// Print the client version
	version, err := c.Version()
	if err != nil {
		fmt.Println("Error getting version:", err)
		return
	}
	fmt.Println("OpenTDF version:", version)
}

