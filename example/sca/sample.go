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

	result, _ := client.GetSca("001")
	fmt.Println("%#v", result)

	policies, _ := client.GetScaChecks("001", "custom_policy")
	fmt.Println("%#v", policies)
}
