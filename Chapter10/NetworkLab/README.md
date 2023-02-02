# Network Lab

To launch a pre-build Linux host with the routers already running for network lab, just use the pre-build image below.

## Using pre-build image on VirtualBox
### Download image at
```
https://www.dropbox.com/s/h6mfz20ocu6i4g3/netlab.vdi.bz2?dl=0
```
### Uncompress `netlab.vdi.bz2`
### Create new VM on VirtualBox using the pre-built image downloaded `netlab.vdi`
### Add port forwarding to 22 to access SSH
Go to Virtualbox settings for the VM and configure NAT network with port forwarding 22, so you can access the network emulation via SSH.

### Accessing the devices via SSH and sudo
- all devices including the host have the username `netlab` with password `netlab`.
- `sudo` is configured to be used by username `netlab` without password.
- Access to `vtysh` on the routers via username `vtysh` with password `vtysh`.

If you want to do yourself or use Qemu image instead, please visit my github instructions at:
https://github.com/brnuts/network-lab

# Testing code
Some examples of Go automation to be used in the network lab are included in the Go directory
