# A quick tour of Rack

[rack](https://github.com/rackspace/rack) is a command-line tool for managing your Rackspace servers.  To get started, download the executable.  Once it's installed, run `rack configure` to set up your credentials.  You'll be asked for:

* Your Rackspace username
* API Key (get this in your account settings screen)
* Region (DFW, IAD, ORD, LON, SYD, HKG)

You can find this information in your web dashboard, but once you have them, you can start to manage your account strictly from the command line.

You default credentials are saved at `~/.rack/config`


## List available images

An image is the OS type you want.  You can get a list of the available image like this (I'm just showing Ubuntu becasue there are a *lot* of images)

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

## List available flavors

A `flavor` describes the specs (RAM, disk, etc) that the new image will run on.

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

## Manage keypairs

Like other cloud services, Rackspace uses keypairs to control access to the servers.  So, when you set up a server, you need to create or specify the keypairs for the machine's you'll want to have access to a server.

*NB: I'm not so sure any of this works the way I think it does!*


### Create a keypair

```
$ rack servers keypair generate --name odewahn-macbook-air
Name		odewahn-macbook-air
Fingerprint	01:de:f4:ee:73:15:46:9b:53:d4:06:9f:be:94:fb:c4
PublicKey	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDkAlOHDRIxij/A46o1MHEWVe4IPYJ5PRyahSh6gTq8gyVZJDltAVODQoudMjCDUb5FzWVKSwZjUr7WCGJvE0J/ClBxBfXfJ8jH+D5EeEQDbiwqHCAq+5/vtP9rf2bqZd1L8STBbcy+eYGSzMO4TO+RTe0XuIsrspL5wNkvnY82pG/CA9pCLsopUjcWG22iNFj29/bUGdgRToecJUud4WT4rfshiEEEXqTnflLxjPeRzf8mipZ+aIPyIPYU6Rf9/UBWgggrBgE0yOU5JLMqhyG8G/m63XgquJlCoQ9+99HzjyywM55Gi8RNJ53nMXLIaN9pjIbsDCBmVmtyl1UXkml/ Generated-by-Nova
PrivateKey	-----BEGIN RSA PRIVATE KEY-----
    Lots of stuff here

```

### Upload your existing keypair

If you have an existing kye on your system (and, if you use practically any service cloud-based service you most likely do!), you can also upload it with `rack`, like this:

```
$ rack servers keypair upload --name andrew-macbook-pro --file ~/.ssh/id_rsa.pub
Name		andrew-macbook-pro
Fingerprint	73:c0:21:55:28:f1:51:06:47:bc:58:fa:6d:91:28:8c
PublicKey	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCyl61Meicp9VNKN6dEQawsCAvg8SHNrdixMDheRs764OM/1JswVLaRuR+XWEZbvE1UHEXv8KEtBgJPZ2h3HBBZpK/QqdP1QjRwQOKAAObUIx6r25tSZ08X1rV3rlZ88jUGwQWY2EeceWY+MQlpmVy0mfT6M0RLDr2oKsyW8kSlOsvfBYt802Td28C2DfmqD+irmHhZbYfDpSOwri4vArWdz0V2Pvq/CrJclG/jq2/POyLoW5hv+HzAnULRfppLYzfv/OazqCZV9vlmfAXjqZDwkSFcd6tnEifije2A64STpadhc7gLcj1ndlVkeC/YZFiHdFNhtxb6F7nYnYyDAgAL apple@new-host-4.home

PrivateKey

```

Then, when you create a server (coming up in the next section!), you can add your key to the new instance to bootstap it.


### List existing keypairs

The folowing command shows the keypairs uploaded to your account:

```
$ rack servers keypair list
Name			Fingerprint
andrew-macbook-pro	73:c0:21:55:28:f1:51:06:47:bc:58:fa:6d:91:28:8c
odewahn-macbook-air	01:de:f4:ee:73:15:46:9b:53:d4:06:9f:be:94:fb:c4
odewahn-macbook-pro	7d:a5:c8:80:f2:dc:61:d3:cd:65:97:ef:66:fb:7f:7a
odewahn-orm-mac		9e:70:4f:55:1b:4f:0e:c9:d6:5f:39:1d:27:86:5c:0d

```

## Create a server

Use `rack servers create` to create a new server.  You'll need to provide the `image-id` and `flavor-id` you want (see the earlier section on how to find this information).  For example, the following command will create a 1GB RAM / 1 CPU server named `swarm-test` running Ubuntu 14.04 with my default keypair:


```
$ rack servers instance create \
    --name swarm-test \
    --image-id 5ed162cc-b4eb-4371-b24a-a0ae73376c73 \
    --flavor-id general1-1 \
    --keypair andrew-macbook-pro

ID		18a93442-9da1-4d48-b80d-923543e96a70
AdminPass	************

```

## List your servers

You can get a list of the servers you have (and their status) using the `rack servers list` command:

```
$ rack servers instance list
ID					Name		Status	PublicIPv4	PrivateIPv4	Image			Flavor
b2772795-a7ed-4f9d-801f-9eb3585b1f87	swarm-test	BUILD			10.176.192.65	5ed162cc-b4eb-4371-b24a-a0ae73376c73	general1-1
e3f8a765-8ced-457b-92ea-a16adbfd578d	interlock	ACTIVE	162.242.242.5	10.208.165.236	5ed162cc-b4eb-4371-b24a-a0ae73376c73	general1-1
95bdf7a8-bedb-4d8e-94d5-0733ab8df5ec	jupyter-kernel	ACTIVE	104.130.20.99	10.184.8.237	06bb130b-7607-46f9-85ae-124bae4d0f5b	onmetal-compute1
0c1e740e-fd06-408e-9d7b-b8f263351123	strata-lb	ACTIVE	23.253.156.165	10.208.196.91	e19a734c-c7e6-443a-8
```

Note that in this log, the `swarm-test` server is in the `BUILD` stage.  Once it's complete, it's status will change to `ACTIVE` and you can log in.

## Log in to a server

Once the server has been built, you can log into it by finding it's public IP address (the `PublicIPv4` field from the `list` command).  If you've uploaded a keypair, you should log right in without a password.  If not, you'll be prompted for the `AdminPass` returned when you created the server:

```
$ ssh root@104.130.13.36
The authenticity of host '104.130.13.36 (104.130.13.36)' can't be established.
RSA key fingerprint is 7e:8b:4c:a1:ee:ae:15:5d:f6:92:3f:fe:80:79:83:26.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '104.130.13.36' (RSA) to the list of known hosts.
Welcome to Ubuntu 14.04.2 LTS (GNU/Linux 3.13.0-58-generic x86_64)

 * Documentation:  https://help.ubuntu.com/

The programs included with the Ubuntu system are free software;
the exact distribution terms for each program are described in the
individual files in /usr/share/doc/*/copyright.

Ubuntu comes with ABSOLUTELY NO WARRANTY, to the extent permitted by
applicable law.

Last login: Thu Jan  1 00:00:10 1970
root@swarm-test:~#

```

## Delete a server

You can delete an unused server using the `rack servers instance delete` command.  You have to specifiy the

```
$ rack servers instance list
ID					Name		Status	PublicIPv4	PrivateIPv4	Image			Flavor
18a93442-9da1-4d48-b80d-923543e96a70	swarm-test	ACTIVE	104.130.13.36	10.176.192.231	5ed162cc-b4eb-4371-b24a-a0ae73376c73	general1-1

$ rack servers instance delete --id 18a93442-9da1-4d48-b80d-923543e96a70
Deleting instance [18a93442-9da1-4d48-b80d-923543e96a70]
```


```
