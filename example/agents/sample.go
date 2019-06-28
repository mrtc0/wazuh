package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mrtc0/wazuh"
)

func main() {
	flag.Parse()
	args := flag.Args()

	endpoint := args[0]
	user := args[1]
	pass := args[2]

	client, err := wazuh.New(endpoint, wazuh.WithBasicAuth(user, pass))

	agents, err := client.GetAllAgents()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%-20v%-20v\n", "ID", "NAME")
	for _, agent := range *agents {
		if agent.Status == "Active" {
			fmt.Printf("%-20v%-20v\n", agent.ID, agent.Name)
		}
	}
}
