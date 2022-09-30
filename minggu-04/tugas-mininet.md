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
Start a minimal topology and enter the CLI:<br>
![1](gambar/mininet-start-topology.jpg)<br>
The default topology is the minimal topology, which includes one OpenFlow kernel switch connected to two hosts, plus the OpenFlow reference controller.
### Step 3: Interact with Hosts and Switches
Display Mininet CLI commands:<br>
![2](gambar/mininet-cli-command.jpg)<br>
Display nodes and links:<br>
![3](gambar/mininet-nodes-links.jpg)<br>
Dump information about all nodes:<br>
![4](gambar/mininet-dump.jpg)<br>
Run a command on a host process:<br>
![5](gambar/mininet-h1-config.jpg)<br>
![6](gambar/mininet-h2-config.jpg)<br>
Print the process list from a host process:<br>
![7](gambar/mininet-h1-ps.jpg)<br>
![8](gambar/mininet-h2-ps.jpg)<br>
Now, verify that you can ping from host 0 to host 1:<br>
![9](gambar/mininet-h1-ping-h2.jpg)<br>
An easier way to run this test is to use the Mininet CLI built-in pingall command, which does an all-pairs ping:<br>
![10](gambar/mininet-pingall.jpg)<br>
Exit the CLI:<br>
![11](gambar/mininet-exit.jpg)<br>
If Mininet crashes for some reason, clean it up:<br>
![12](gambar/mininet-cleanup.jpg)<br>












