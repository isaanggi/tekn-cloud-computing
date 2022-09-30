# Mininet Walkthrough on Ubuntu
### Step 1: Install
Login, install net-tools, git clone and install Mininet on your Ubuntu system.
```bash
osboxes@isaanggi:~$ sudo apt-get install net-tools
osboxes@isaanggi:~$ git clone https://github.com/mininet/mininet
osboxes@isaanggi:~$ mininet/util/install.sh -w
osboxes@isaanggi:~$ mininet/util/install.sh -a
```
Wait until the installation is complete.
### Step 2: Minimal Topology
Start a minimal topology and enter the CLI:
![1](gambar/mininet-start-topology.jpg)
The default topology is the minimal topology, which includes one OpenFlow kernel switch connected to two hosts, plus the OpenFlow reference controller.
### Step 3: Interact with Hosts and Switches
Display Mininet CLI commands:
![2](gambar/mininet-cli-command.jpg)
Display nodes and links:
![3](gambar/mininet-nodes-links.jpg)
Dump information about all nodes:
![4](gambar/mininet-dump.jpg)
Run a command on a host process:
![5](gambar/mininet-h1-config.jpg)
![6](gambar/mininet-h2-config.jpg)
Print the process list from a host process:
![7](gambar/mininet-h1-ps.jpg)
![8](gambar/mininet-h2-ps.jpg)
Now, verify that you can ping from host 0 to host 1:
![9](gambar/mininet-h1-ping-h2.jpg)
An easier way to run this test is to use the Mininet CLI built-in pingall command, which does an all-pairs ping:
![10](gambar/mininet-pingall.jpg)
Exit the CLI:
![11](gambar/mininet-exit.jpg)
If Mininet crashes for some reason, clean it up:
![12](gambar/mininet-cleanup.jpg)












