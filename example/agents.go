package main

import (
	"fmt"

	"github.com/mrtc0/wazuh"
)

func main() {
	wazuh := wazuh.New("https://wazuh.example.com:55000/")
	wazuh.SetBasicAuth("foo", "bar")

	fmt.Println("[+] Get All Agents")
	agents, _ := wazuh.GetAllAgents()
	fmt.Printf("%-20v%-20v%-20v%-20v%-20v\n", "ID", "STATUS", "NAME", "IP", "OS")
	for _, agent := range *agents {
		fmt.Printf("%-20v%-20v%-20v%-20v%-20v\n", agent.ID, agent.Status, agent.Name, agent.IP, agent.Os.Name)
	}

	fmt.Println("\n[+] Get agent with id '001'")
	agent, _ := wazuh.GetAnAgent("001")
	fmt.Printf("%-20v%-20v%-20v%-20v%-20v\n", "ID", "STATUS", "NAME", "IP", "OS")
	fmt.Printf("%-20v%-20v%-20v%-20v%-20v\n", agent.ID, agent.Status, agent.Name, agent.IP, agent.Os.Name)
}
