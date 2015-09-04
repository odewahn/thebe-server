# A quick tour of Rack

[rack](https://github.com/rackspace/rack) is the new Rackspace command-line tool for managing their clusters.  You'll need to download the executable, and once it's installd, run `rack configure` to set up your credentials.  You'll be asked for:

* Username
* API Key (get this in your account settings screen)
* Region (DFW, IAD, ORD, LON, SYD, HKG)

You default credentials are saved at `~/.rack/config`


## List available images

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

### List available flavors

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

### List existing keypairs

```
$ rack servers keypair list
Name			Fingerprint
odewahn-macbook-air	01:de:f4:ee:73:15:46:9b:53:d4:06:9f:be:94:fb:c4
odewahn-orm-mac		9e:70:4f:55:1b:4f:0e:c9:d6:5f:39:1d:27:86:5c:0d
```

## SSH-ing into your machine with a keypair

Wha? Huh?  I would think you do something like this:

```
ssh -i ~/.ssh/my-rack-keypair ...
```
