## Install HAProx

The interlock plugin requires HAProxy v 1.5+; if you're on Ubuntu 14.04, you can install it with the following commands:

```
echo deb http://archive.ubuntu.com/ubuntu trusty-backports main universe | \
      tee /etc/apt/sources.list.d/backports.list

apt-get update

apt-get install haproxy -t trusty-backports
```

You'll also need your cert file from the rackspace swarm.
