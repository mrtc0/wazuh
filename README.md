# Wazuh API in Go 

[![Build Status](https://travis-ci.org/mrtc0/wazuh.svg?branch=master)](https://travis-ci.org/mrtc0/wazuh)

---

:warning: **WIP** :warning:

Wazuh REST API library

https://documentation.wazuh.com/current/user-manual/api/reference.html

##### Progress

- [ ] Active Response
- [ ] Agents
- [ ] Cache
- [ ] Ciscat
- [ ] Cluster
- [ ] Decoders
- [ ] Experimental
- [ ] Manager
- [ ] Rootcheck
- [x] Rules
- [ ] Syscheck
- [x] Syscollector

## Install

```powershell
$ go get -u github.com/mrtc0/wazuh
```

## Usage

### Getting agents information

```go
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
```

```powershell
[+] Get All Agents
ID                  STATUS              NAME                IP                  OS
000                 Active              wazuh-manager       127.0.0.1           Ubuntu
001                 Active              api                 10.142.0.3          Ubuntu
002                 Active              proxy               10.142.0.5          Ubuntu
003                 Active              frontend            10.142.0.2          Ubuntu
004                 Active              redis               10.142.0.4          Ubuntu

[+] Get agent with id '001'
ID                  STATUS              NAME                IP                  OS
001                 Active              api                 10.142.0.3          Ubuntu
```

### Get Syscollector

```go
package main

import (
	"fmt"

	"github.com/mrtc0/wazuh"
)

func main() {
	wazuh := wazuh.New("https://wazuh.example.com:55000/")
	wazuh.SetBasicAuth("foo", "bar")

	hw_info, _ := wazuh.GetHardwareInfo("001")
	fmt.Printf("%+v\n", hw_info)
}
```

```powershell
&{BoardSerial: RAM:{Usage:85 Total:5071792 Free:786204} CPU:{Cores:1 Mhz:2300 Name:Intel(R) Xeon(R) CPU @ 2.30GHz} Scan:{ID:1887178901 Time:2019/01/05 20:05:38}}
```

[more examples](./example)

## Contributing

You are more than welcome to contribute to this project.  
Fork and make a Pull Request, or create an Issue if you see any problem.


