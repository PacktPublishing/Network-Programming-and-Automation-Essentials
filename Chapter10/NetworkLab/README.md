# Network Lab

To launch a pre-build Linux host with the routers already running for network lab, just use the pre-build image below.

## Using pre-build image on VirtualBox
- Download image at
```
https://www.dropbox.com/s/h6mfz20ocu6i4g3/netlab.vdi.bz2?dl=0
```
* Uncompress `netlab.vdi.bz3`
* Create new VM on VirtualBox using the pre-built image as the disk `netlab.vdi`
* Add port forwarding to 22 to access SSH by going to Virtualbox settings for the VM and configure NAT network with port forwarding 22, so you can access the network emulation via SSH. With that you can access the network lab by doing `ssh netlab@localhost`

<img src="https://github.com/brnuts/netlab/blob/main/Port-foward-example-Virtualbox.png" width="400"/>

* You might need to install Guest Additions on the Linux VM, use these steps:
   * Open VirtualBox.
   * Right-click the virtual machine, select the Start submenu and choose the Normal Start option.
   * Sign in to netlab (username: netlab password: netlab)
   * Click the Devices menu and select the Insert Guest Additions CD image option.
   * On Linux do `sudo mount /dev/cdrom`
   * Run the command `sudo sh /dev/cdrom0/VBoxLinuxAdditions.run`
   * Reboot your VM

- all devices including the host have the username `netlab` with password `netlab`.
- `sudo` is configured to be used by username `netlab` without password.
- Access to `vtysh` on the routers via username `vtysh` with password `vtysh`.

If you want to do yourself or use Qemu image instead, please visit my github instructions at:
https://github.com/brnuts/netlab

# Connecting the devices
To connect the devices you will just need to get the Go package that connect them for you located at https://github.com/brnuts/netlab/Go
