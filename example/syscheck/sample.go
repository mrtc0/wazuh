package main

import (
	"flag"
	"fmt"

	"github.com/mrtc0/wazuh"
)

func main() {
	flag.Parse()
	args := flag.Args()

	endpoint := args[0]
	user := args[1]
	pass := args[2]

	client, _ := wazuh.New(endpoint, wazuh.WithBasicAuth(user, pass))

	result, _ := client.GetSyscheckFiles("001")
	fmt.Printf("%#v\n", result)

	result2, _ := client.RunSyscheckAllAgents()
	fmt.Printf("%#v\n", *result2)

	result3, _ := client.RunSyscheckAgent("001")
	fmt.Printf("%#v\n", *result3)

	result4, _ := client.GetLastSyscheckScan("001")
	fmt.Printf("%#v\n", *result4)

	result5, _ := client.ClearSyscheckDatabase("001")
	fmt.Printf("%#v\n", *result5)
}
