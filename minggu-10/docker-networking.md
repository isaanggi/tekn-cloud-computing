# Docker Networking Hands-on Lab

Laporan beserta gambar dibawah ini adalah hasil praktikum melalui [Docker Networking Hands-on Lab](https://training.play-with-docker.com/docker-networking-hol/), sehingga untuk materi dan penjelasan lebih detailnya dapat diakses melalui web tersebut.

### Task

- [Section #1 - Networking Basics](https://github.com/isaanggi/tekn-cloud-computing/edit/main/minggu-10/docker-networking.md#section-1---networking-basics)
- [Section #2 - Bridge Networking](https://github.com/isaanggi/tekn-cloud-computing/edit/main/minggu-10/docker-networking.md#section-2---bridge-networking)
- Section #3 - Overlay Networking
- Cleaning Up

## Section #1 - Networking Basics

### Step 1: The Docker Network Command
The ```docker network``` command is the main command for configuring and managing container networks. Run the ```docker network``` command from the first terminal.
<div><img src="gambar/SS1.png"></div>
<div><img src="gambar/SS2.png"></div>

### Step 2: List networks
Run a ```docker network ls``` command to view existing container networks on the current Docker host.
<div><img src="gambar/SS3.png"></div><br>

### Step 3: Inspect a network
The ```docker network inspect``` command is used to view network configuration details. Use ```docker network inspect <network>``` to view configuration details of the container networks on your Docker host. The command below shows the details of the network called ```bridge```.
<div><img src="gambar/SS4.png"></div><br>

### Step 4: List network driver plugins
The ```docker info``` command shows a lot of interesting information about a Docker installation. Run the ```docker info``` command and locate the list of network plugins.
<div><img src="gambar/SS5.png"></div><br>

## Section #2 - Bridge Networking
### Step 1: The Basics
Every clean installation of Docker comes with a pre-built network called bridge. Verify this with the docker ```network ls```.
<div><img src="gambar/SS6.png"></div><br>

Install the ```brctl``` command and use it to list the Linux bridges on your Docker host. You can do this by running ```sudo apt-get install bridge-utils```.
<div><img src="gambar/SS7.png"></div><br>

Then, list the bridges on your Docker host, by running ```brctl show```. You can also use the ```ip a``` command to view details of the docker0 bridge.

<div><img src="gambar/SS8.png"></div><br>

### Step 2: Connect a container
The bridge network is the default network for new containers. This means that unless you specify a different network, all new containers will be connected to the bridge network. Create a new container by running ```docker run -dt ubuntu sleep infinity```. This command will create a new container based on the ```ubuntu:latest``` image and will run the ```sleep``` command to keep the container running in the background. You can verify our example container is up by running ```docker ps```.
<div><img src="gambar/SS9.png"></div><br>

As no network was specified on the ```docker run``` command, the container will be added to the bridge network. Run the ```brctl show``` command again. Notice how the docker0 bridge now has an interface connected. This interface connects the docker0 bridge to the new container just created. You can inspect the bridge network again, by running ```docker network inspect bridge```, to see the new container attached to it.
<div><img src="gambar/SS10.png"></div><br>



COMMIT DULU :D

==============================================


<div><img src="gambar/SS11.png"></div><br>
<div><img src="gambar/SS12.png"></div><br>
<div><img src="gambar/SS13.png"></div><br>
<div><img src="gambar/SS14.png"></div><br>
<div><img src="gambar/SS15.png"></div><br>
<div><img src="gambar/SS16.png"></div><br>
<div><img src="gambar/SS17.png"></div><br>
<div><img src="gambar/SS18.png"></div><br>
<div><img src="gambar/SS19.png"></div><br>
<div><img src="gambar/SS20.png"></div><br>
<div><img src="gambar/SS21.png"></div><br>
<div><img src="gambar/SS22.png"></div><br>
<div><img src="gambar/SS23.png"></div><br>
<div><img src="gambar/SS24.png"></div><br>
<div><img src="gambar/SS25.png"></div><br>
<div><img src="gambar/SS26.png"></div><br>
<div><img src="gambar/SS27.png"></div><br>
<div><img src="gambar/SS28.png"></div><br>
<div><img src="gambar/SS29.png"></div><br>
