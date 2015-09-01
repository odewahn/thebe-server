# Rackspace keypairs

Like other cloud services, Rackspace uses keypairs to control access to the servers.  So, when you set up a server, you need to create or specify the keypairs for the machine's you'll want to have access to a server.

*NB: I'm not so sure any of this works the way I think it does!*


## Create a keypair

```
$ rack servers keypair generate --name odewahn-macbook-air
Name		odewahn-macbook-air
Fingerprint	01:de:f4:ee:73:15:46:9b:53:d4:06:9f:be:94:fb:c4
PublicKey	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDkAlOHDRIxij/A46o1MHEWVe4IPYJ5PRyahSh6gTq8gyVZJDltAVODQoudMjCDUb5FzWVKSwZjUr7WCGJvE0J/ClBxBfXfJ8jH+D5EeEQDbiwqHCAq+5/vtP9rf2bqZd1L8STBbcy+eYGSzMO4TO+RTe0XuIsrspL5wNkvnY82pG/CA9pCLsopUjcWG22iNFj29/bUGdgRToecJUud4WT4rfshiEEEXqTnflLxjPeRzf8mipZ+aIPyIPYU6Rf9/UBWgggrBgE0yOU5JLMqhyG8G/m63XgquJlCoQ9+99HzjyywM55Gi8RNJ53nMXLIaN9pjIbsDCBmVmtyl1UXkml/ Generated-by-Nova
PrivateKey	-----BEGIN RSA PRIVATE KEY-----
    Lots of stuff here

```

## List existing keypairs

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
