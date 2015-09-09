# thebe-server

Goal is to create a system that we can use to scale the thebe backends.  The idea is to:

* Mimic the tmpnb API for spawning an image, except that we can also spawn a specific container
* Use docker swarm as a backend; in this case, we'll also use Rackspace's hosting service
* Use interlock (https://github.com/ehazlett/interlock) to handle service discovery to the Docker containers

The following diagram summarizes the architecture:

![Architecture](public/images/architecture.png)


## Docker Swarm on Rackspace

All the docker containers are launched and managed by on a swarm cluster managed by Rackspace. We'll use the Rackspace offering for this as a host (https://mycluster.rackspacecloud.com/).  

[Creating a swarm cluster on Rackspace](public/rackspace-create-cluster.md) has the details.  If you're unfamiliar with some of the Rackspace tools, the [Rack overview](public/rack-overview.md) will give you a bit of background on the service.

## Interlock

Interlock (https://github.com/ehazlett/interlock), a Docker plugin service, provides a slick way to us HAProxy to handle service discovery for clusters on the swarm.  So, whenever we start a container on the swarm, Interlock will expose it using the `hostname` we supply in the `docker run` command.  When you kill a container, interlock will delete it from the host setup.

[Service discovery with interlock (HAProxy plugin)](public/service-discovery-with-interlock.md) has the details.

## Thebe Server

`thebe-server` is a Go-based application that provides an API for starting and stopping containers on the swarm.  Basically, it provide similar endpoints used by [tmpnb](https://github.com/jupyter/tmpnb), but it does not try to handle pooling or other services.  Instead, it's focus is on allowing you to start many different kinds of containers that you can connect to `thebe`

[Running `thebe-server`](public/thebe-server.md) has details on the API and on the development and deployment process.  The cool things about `thebe-server` is that it is also a containerized, and can run on the same swarm setup as the other containers, making it simple to manage.
