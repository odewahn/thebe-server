# Configure the server to control the cluster

We'll need two servers: `rackspace-swarm-test` and `interlock`.


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
