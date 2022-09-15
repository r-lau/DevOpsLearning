## Day 03


## Linux

Brief knowldge of CentOS and YUM(particular for ansible)

Package manager in centOS is RPM(Red Hat Package Manager)

```systemctl start <service>``` is the prefer way to start a service
```systemctl enable <service>``` configured a service to start at startup

Configuring application as a service
Adding a .service file to ```/etc/systemd/system``` with the following content:
``` [Content] ExecStart=<command to invoke the application>```

Can then use ```systemctl daemon-reload``` to let the system know there's a new service configured followed by ```systemctl start <new service>```

Introduction to Vagrant, but didn't follow lab since it's involving Virtualbox
Compatible with VMware, Docker, and Hyper-V as well
When Vagrantfile has changes made, must run ```vagrant reload``` to have changes take effect

## Networking in Linux

Hosts file is in ```/etc/hosts```
Hosts file is the source of truth for that device
DNS Server takes care of the name resolution, so modify hosts file in the server instead and that's why hosts are pointed to the DNS server
Entry for nameserver ```/etc/resolv.conf``` will be under nameserver and the IP

With duplicate DNS names in two different hosts(host A and DNS server) the local file will take precedence

If you don't have an entry for something like www.facebook.com, you may add an additional entry to the resolv.conf file to have nameserver point to 8.8.8.8 because 8.8.8.8 will have the DNS resolution

Better to have DNS server forward any unknown host names to the public name server.

To make internal DNS entries; add ```search  yourcompany.com``` to the resolv.conf file the same line can contain a subdomain so it can be ```yourcompany.com prod.yourcompany.com``` 

```nslookup``` does not consider the entries in the local etc\hosts file
```dig``` is the same as well. This command provides more DNS information

To assign an IP to eth0 use ``` ip addr add <ipv4/cidr> dev eth0```

To configure a gateway use ```ip route add <ipv4/cidr> via <gateway IP>``` from host A on network A so that it can reach host C on network B

To reach out to the internet use ```ip route add default via <gateway IP>```
Sidenote: default can be replaced with 0.0.0.0 in the destination field when checking ```route```
A 0.0.0.0 gateway means it doesn't need a gateway to access hosts within the same network

Eth0 does not forward to eth1 by default for security reasons(Ex: Eth0 is private and Eth1 can be a public network)

```/proc/sys/net/ipv4/ip_forward``` is where the setting lies and the value is 0 by default
This setting does not persist through reboots, for that, must adjust ```net.ipv4.ip_forward``` in ```/etc/sysctl.conf```

