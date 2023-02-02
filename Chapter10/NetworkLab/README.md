# Network Lab

To launch a pre-build Linux host with the routers already running for network lab, just download it:

## Using pre-build image on VirtualBox
### Download image at
```
https://www.dropbox.com/s/h6mfz20ocu6i4g3/netlab.vdi.bz2?dl=0
```

### Add port forwarding to 22 to access SSH
Go to Virtualbox settings for the VM and configure NAT network with port forwarding 22, so you can access the network emulation via SSH.

### Username and password
- all devices including the host has the username: `netlab` and password: `netlab`.

- `sudo` is configured to be used by username `netlab` without password.


If you want to do yourself or use Qemu image instead, please visit my github instructions at:
https://github.com/brnuts/network-lab

# Testing code
Some examples of Go automation to be used in the network lab are included in the Go directory
