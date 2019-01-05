package main

import (
	"fmt"

	"github.com/mrtc0/wazuh"
)

func main() {
	wazuh := wazuh.New("https://wazuh.example.com:55000/")
	wazuh.SetBasicAuth("foo", "bar")

	fmt.Println("[+] Get hardware info of an agent with id '001'")
	hw_info, _ := wazuh.GetHardwareInfo("001")
	fmt.Printf("%+v\n", hw_info)

	fmt.Println("\n[+] Get network address info of an agent with id '001'")
	nwaddr_info, _ := wazuh.GetNetworkAddrInfo("001")
	fmt.Printf("%+v\n", nwaddr_info)

	fmt.Println("\n[+] Get network protocol info of an agent with id '001'")
	nwproto_info, _ := wazuh.GetNetworkProtocolInfo("001")
	fmt.Printf("%+v\n", nwproto_info)

	fmt.Println("\n[+] Get OS info of an agent with id '001'")
	os_info, _ := wazuh.GetOSInfo("001")
	fmt.Printf("%+v\n", os_info)

	fmt.Println("\n[+] Get packages info of an agent with id '001'")
	packages_info, _ := wazuh.GetPackagesInfo("001")
	fmt.Printf("%+v\n", (*packages_info)[0])

	fmt.Println("\n[+] Get ports info of an agent with id '001'")
	ports_info, _ := wazuh.GetPortsInfo("001")
	fmt.Printf("%+v\n", ports_info)

	fmt.Println("\n[+] Get processes info of an agent with id '001'")
	processes_info, _ := wazuh.GetProcessesInfo("001")
	fmt.Printf("%+v\n", (*processes_info)[0])
}
