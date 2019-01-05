## agents

```powershell
$ go run agents.go
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

## syscollector

```powershell
$ go run syscollector.go
[+] Get hardware info of an agent with id '001'
&{BoardSerial: RAM:{Usage:85 Total:5071792 Free:786204} CPU:{Cores:1 Mhz:2300 Name:Intel(R) Xeon(R) CPU @ 2.30GHz} Scan:{ID:1887178901 Time:2019/01/05 20:05:38}}

[+] Get network address info of an agent with id '001'
&[{Broadcast:10.142.0.3 ScanID:77802661 Proto:ipv4 Netmask:255.255.255.255 Address:10.142.0.3} {Broadcast: ScanID:77802661 Proto:ipv6 Netmask:ffff:ffff:ffff:ffff:: Address:fe80::4001:aff:fe8e:3}]

[+] Get network protocol info of an agent with id '001'
&[{Dhcp:enabled ScanID:77802661 Iface:ens4 Type:ipv4 Gateway:10.142.0.1} {Dhcp:enabled ScanID:77802661 Iface:ens4 Type:ipv6 Gateway:}]

[+] Get OS info of an agent with id '001'
&{Sysname:Linux Version:#42~16.04.1-Ubuntu SMP Wed Oct 24 17:09:54 UTC 2018 Architecture:x86_64 Scan:{ID:1452170967 Time:2019/01/05 20:05:38} Release:4.15.0-39-generic Hostname:backend-api Os:{Major:16 Name:Ubuntu Platform:ubuntu Version:16.04.5 LTS (Xenial Xerus) Codename:Xenial Xerus Minor:04}}

[+] Get packages info of an agent with id '001'
{Vendor:Ubuntu Core developers <ubuntu-devel-discuss@lists.ubuntu.com> Name:libquadmath0 Scan:{ID:335503100 Time:2019/01/05 20:05:38} Section:libs Format:deb Priority:optional Source:gcc-5 Version:5.4.0-6ubuntu1~16.04.10 Architecture:amd64 Multiarch:same Size:265 Description:GCC Quad-Precision Math Library}

[+] Get ports info of an agent with id '001'
&[{Remote:{IP:0.0.0.0 Port:0} Scan:{ID:1300173469 Time:2019/01/05 20:05:44} Inode:20913 State:listening TxQueue:0 Protocol:tcp RxQueue:0 Local:{IP:0.0.0.0 Port:22}} {Remote:{IP:0.0.0.0 Port:0} Scan:{ID:1300173469 Time:2019/01/05 20:05:44} Inode:22181 State:listening TxQueue:0 Protocol:tcp RxQueue:0 Local:{IP:127.0.0.1 Port:24224}} {Remote:{IP:0.0.0.0 Port:0} Scan:{ID:1300173469 Time:2019/01/05 20:05:44} Inode:22769 State:listening TxQueue:0 Protocol:tcp RxQueue:0 Local:{IP:127.0.0.1 Port:514}} {Remote:{IP:0.0.0.0 Port:0} Scan:{ID:1300173469 Time:2019/01/05 20:05:44} Inode:22174 State:listening TxQueue:0 Protocol:tcp RxQueue:0 Local:{IP:0.0.0.0 Port:24231}} {Remote:{IP:0.0.0.0 Port:0} Scan:{ID:1300173469 Time:2019/01/05 20:05:44} Inode:21807 State:listening TxQueue:0 Protocol:tcp RxQueue:0 Local:{IP:127.0.0.1 Port:3306}} {Remote:{IP:0.0.0.0 Port:0} Scan:{ID:1300173469 Time:2019/01/05 20:05:44} Inode:1701143 State:listening TxQueue:0 Protocol:tcp RxQueue:0 Local:{IP:127.0.0.1 Port:6379}} {Remote:{IP:0.0.0.0 Port:0} Scan:{ID:1300173469 Time:2019/01/05 20:05:44} Inode:925422 State:listening TxQueue:0 Protocol:tcp RxQueue:0 Local:{IP:0.0.0.0 Port:80}} {Remote:{IP::: Port:0} Scan:{ID:1300173469 Time:2019/01/05 20:05:44} Inode:20921 State:listening TxQueue:0 Protocol:tcp6 RxQueue:0 Local:{IP::: Port:22}} {Remote:{IP::: Port:0} Scan:{ID:1300173469 Time:2019/01/05 20:05:44} Inode:925423 State:listening TxQueue:0 Protocol:tcp6 RxQueue:0 Local:{IP::: Port:80}}]

[+] Get processes info of an agent with id '001'
{Euser:root Tty:0 Rgroup:root Sgroup:root Scan:{ID:1139544256 Time:2019/01/05 20:05:44} Resident:1450 Share:990 StartTime:5 Pid:1 Session:1 Stime:789 VMSize:185220 Size:46305 Ppid:0 Egroup:root Name:systemd Pgrp:1 Tgid:1 Utime:775 Cmd:/sbin/init Priority:20 Fgroup:root State:S Ruser:root Suser:root Nlwp:1 Processor:0 Nice:0 Argvs:}
```
