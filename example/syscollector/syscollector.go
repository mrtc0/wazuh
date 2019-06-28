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

	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}

	fmt.Println("[+] Get hardware info of an agent with id '001'")
	hw_info, _ := client.GetHardwareInfo("001")
	fmt.Printf("%+v\n", hw_info)

	fmt.Println("\n[+] Get network address info of an agent with id '001'")
	nwaddr_info, _ := client.GetNetworkAddrInfo("001")
	fmt.Printf("%+v\n", nwaddr_info)

	fmt.Println("\n[+] Get network protocol info of an agent with id '001'")
	nwproto_info, _ := client.GetNetworkProtocolInfo("001")
	fmt.Printf("%+v\n", nwproto_info)

	fmt.Println("\n[+] Get OS info of an agent with id '001'")
	os_info, _ := client.GetOSInfo("001")
	fmt.Printf("%+v\n", os_info)

	fmt.Println("\n[+] Get packages info of an agent with id '001'")
	packages_info, _ := client.GetPackagesInfo("001")
	fmt.Printf("%+v\n", (*packages_info)[0])

	fmt.Println("\n[+] Get ports info of an agent with id '001'")
	ports_info, _ := client.GetPortsInfo("001")
	fmt.Printf("%+v\n", ports_info)

	fmt.Println("\n[+] Get processes info of an agent with id '001'")
	processes_info, _ := client.GetProcessesInfo("001")
	fmt.Printf("%+v\n", (*processes_info)[0])
}
