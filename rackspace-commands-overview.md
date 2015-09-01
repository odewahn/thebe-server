## Setting up `swarm-manager` and `interlock` servers

### Download and configure `rack`

[rack](https://github.com/rackspace/rack) is the new Rackspace command-line tool for managing their clusters.  You'll need to download the executable, and once it's installd, run `rack configure` to set up your credentials.  You'll be asked for:

* Username
* API Key (get this in your account settings screen)
* Region (DFW, IAD, ORD, LON, SYD, HKG)

You default credentials are saved at `~/.rack/config`


### Get a list of available images

An image is basically the OS type you want.  You can get a list of the available image like this (I'm just showing Ubuntu becasue there are a *lot* of images)

```
$ rack servers image list | grep Ubuntu
973775ab-0653-4ef8-a571-7a2777787735	Ubuntu 12.04 LTS (Precise Pangolin) (PVHVM)			ACTIVE	20512
656e65f7-6441-46e8-978d-0d39beaaf559	Ubuntu 12.04 LTS (Precise Pangolin) (PV)			ACTIVE	20512
4315b2dc-23fc-4d81-9e73-aa620357e1d8	Ubuntu 15.04 (Vivid Vervet) (PVHVM)				ACTIVE	20512
458e08d7-d230-41a7-9aa2-335129d2e49c	OnMetal - Ubuntu 15.04 (Vivid Vervet)				ACTIVE	20512
fb508530-2500-4a30-a947-038491df2bb5	OnMetal - Ubuntu 14.04 LTS (Trusty Tahr)			ACTIVE	20512
eb6f98a3-5f5d-4153-a011-99823e076dd7	OnMetal - Ubuntu 12.04 LTS (Precise Pangolin)			ACTIVE	20512
09de0a66-3156-48b4-90a5-1cf25a905207	Ubuntu 14.04 LTS (Trusty Tahr) (PVHVM)				ACTIVE	20512
5ed162cc-b4eb-4371-b24a-a0ae73376c73	Ubuntu 14.04 LTS (Trusty Tahr) (PV)				ACTIVE	20512
```

### Get  list of available flavors

A `flavor` is basically the specs (RAM, disk, etc) that the new image will run on.

```
$ rack servers flavor list
ID			Name			RAM	Disk	Swap	VCPUs	RxTxFactor
compute1-15		15 GB Compute v1	15360	0	0	8	1250
compute1-30		30 GB Compute v1	30720	0	0	16	2500
compute1-4		3.75 GB Compute v1	3840	0	0	2	312.5
compute1-60		60 GB Compute v1	61440	0	0	32	5000
compute1-8		7.5 GB Compute v1	7680	0	0	4	625
general1-1		1 GB General Purpose v1	1024	20	0	1	200
general1-2		2 GB General Purpose v1	2048	40	0	2	400
general1-4		4 GB General Purpose v1	4096	80	0	4	800
general1-8		8 GB General Purpose v1	8192	160	0	8	1600
...
performance2-120	120 GB Performance	122880	40	0	32	10000
performance2-15		15 GB Performance	15360	40	0	4	1250
performance2-30		30 GB Performance	30720	40	0	8	2500
performance2-60		60 GB Performance	61440	40	0	16	5000
performance2-90		90 GB Performance	92160	40	0	24	7500
```

## Create and configure the two servers

First, I'll make two servers: `rackspace-swarm-test` and `interlock`.


The first server, `rackspace-swarm-test`, will be used to launch and control the cluster.  Here's the command to create it:

```
rack servers instance create \
    --name rackspace-swarm-test \
    --image-id 5ed162cc-b4eb-4371-b24a-a0ae73376c73 \
    --flavor-id general1-1
```

You'll get back a response like this:

```
ID		e3f8a765-8ced-457b-92ea-a16adbfd578d
AdminPass	************
```

Once you get the login, run the following command to get your IP-address:

```
$ rack servers instance list
ID					Name			Status	PublicIPv4	PrivateIPv4	Image		Flavor
9ca37f1c-1951-43a5-a478-9cfef4571609	rackspace-swarm-test	ACTIVE	104.130.169.111	10.176.194.230	5ed162cc-b4e
...
```

Then, login, like this:

```
$ ssh root@104.130.169.111
```

Finally, to make things simpler, add your public key to the `~/.ssh/authorized_keys` file and do whatever other setup you would normally do on a new user account (`apt-get update`, create a user, etc).

Use the password that you got from the command line tool.  You can also probably use a [keypair](keypair.md), although I've not figured out how to make that work.  *ALSO, THIS WOULD BE A GOOD PLACE TO HAVE SOMETHING LIKE ANSIBLE*.

Once you've set up that server, create another one to run [Interlock](https://github.com/ehazlett/interlock), the proxy server for swarm:

```
rack servers instance create \
    --name interlock \
    --image-id 5ed162cc-b4eb-4371-b24a-a0ae73376c73 \
    --flavor-id general1-1
```


# Set up your cluster

Once you've got the two servers set up, you're ready to start the swarm.  First, log into the [rackspace cluster manager](https://mycluster.rackspacecloud.com/).  
From the main page, create a new cluster (right now there's only 1 button to push).  I've named mine `thebe-test`:

![Cluster dashboard](images/rackspace-cluster.png)

Once you've got the cluster running, you'll need to download the credentials file by clicking the "Download credentials" link in the "Action" area.  (It looks like a circle with an arrow pointing down.)

When you unzip it, you'll see a new
directory whose name looks like a UUID or a swarm id.  

```
$ unzip thebe-test.zip
Archive:  thebe-test.zip
   creating: e1c3155e-a0f2-4e8b-acb8-de123e8dd3c6/
  inflating: e1c3155e-a0f2-4e8b-acb8-de123e8dd3c6/ca.pem  
  inflating: e1c3155e-a0f2-4e8b-acb8-de123e8dd3c6/README.md  
  inflating: e1c3155e-a0f2-4e8b-acb8-de123e8dd3c6/ca-key.pem  
  inflating: e1c3155e-a0f2-4e8b-acb8-de123e8dd3c6/docker.env  
  inflating: e1c3155e-a0f2-4e8b-acb8-de123e8dd3c6/cert.pem  
  inflating: e1c3155e-a0f2-4e8b-acb8-de123e8dd3c6/key.pem  
$ cd e1c3155e-a0f2-4e8b-acb8-de123e8dd3c6/
$ ls -la
total 48
drwxr-xr-x@  8 apple  staff   272 Aug 28 10:36 .
drwx------+ 56 apple  staff  1904 Sep  1 14:50 ..
-rw-r--r--@  1 apple  staff  3177 Aug 28 10:36 README.md
-rw-r--r--@  1 apple  staff  1766 Aug 28 10:35 ca-key.pem
-rw-r--r--@  1 apple  staff  1119 Aug 28 10:35 ca.pem
-rw-r--r--@  1 apple  staff  1086 Aug 28 10:35 cert.pem
-rw-r--r--@  1 apple  staff   153 Aug 28 10:36 docker.env
-rw-r--r--@  1 apple  staff  1679 Aug 28 10:35 key.pem
$
```

You'll see the following files inside this new directory:

* README.md - a set of instructions with how to use the credentials
* ca.pem - Certificate Authority, used by clients to validate servers
* cert.pem - Client Certificate, used by clients to identify themselves to servers
* key.pem - Client Private Key, used by clients to encrypt their requests
* ca-key.pem - Certificate Authority Key, private file used to generate more client certificates.
* docker.env - Shell environment config file


## Using `rackspace-swarm-test`

If you don't have Docker 1.6 installed on your local host machine, you can also run it from the `rackspace-swarm-test` server you created earlier.  

Once you've downloaded the credentials, just `scp` the file over, like this:

```
scp your-credential-file.zip root@rackspace-swarm-test
```

Once it's copied, login to the server, unzip it, and then you're good to go.



### Set up the docker environment

Run the `docker.env` file inside your new directory in order to set your credentials so that you can access the swarm:

```
$ cd e1c3155e-a0f2-4e8b-acb8-de123e8dd3c6/
$ source docker.env
```



You can use `docker info` to see the setup:

```
# docker info
Containers: 12
Images: 8
Storage Driver:
Role: primary
Strategy: spread
Filters: affinity, health, constraint, port, dependency
Nodes: 2
 e1c3155e-a0f2-4e8b-acb8-de123e8dd3c6-n1: 104.130.0.51:42376
  └ Containers: 6
  └ Reserved CPUs: 0 / 12
  └ Reserved Memory: 0 B / 2.1 GiB
  └ Labels: executiondriver=native-0.2, kernelversion=3.12.36-2-rackos, operatingsystem=Debian GNU/Linux 7 (wheezy) (containerized), storagedriver=aufs
 e1c3155e-a0f2-4e8b-acb8-de123e8dd3c6-n2: 104.130.0.52:42376
  └ Containers: 6
  └ Reserved CPUs: 0 / 12
  └ Reserved Memory: 0 B / 2.1 GiB
  └ Labels: executiondriver=native-0.2, kernelversion=3.12.36-2-rackos, operatingsystem=Debian GNU/Linux 7 (wheezy) (containerized), storagedriver=aufs
Execution Driver:
Kernel Version:
Operating System:
CPUs: 24
Total Memory: 4.2 GiB
Name: cbf73ed2e215
ID:
Http Proxy:
Https Proxy:
No Proxy:
```

# Using the Swarm

Once you get all the credentials set up, using the swarm looks a lot like using regular old docker, except now everything is running on a cluster.  

For example, here's how you start a notebook on the swarm

```
docker run -d            \
   -p 8888               \
   ipython/scipystack    \
   /bin/sh -c 'ipython notebook --ip=0.0.0.0 --no-browser'
```
