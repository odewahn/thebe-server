# The `thebe-server` API

`thebe-server` is a simple API that is used to start and stop containers on the swarm.  It provides the following endpoints:

* `/spawn` -- starts a container on the swarm
* ???

To run it (either locally or on production), be sure to create a `.env` file with the following variables (the value will, of course, depend on your environment):

* PORT=8000
* DOCKER_HOST=tcp://104.130.0.52:2376
* SWARM_CREDS_DIR=/Users/apple/Desktop/9fadfa89-0400-453a-a7eb-436aea737831
* THEBE_SERVER_BASE_URL="interlock.odewahn.com"


## Development

Use [gin](https://github.com/codegangsta/gin) to develop the `thebe-server` locally. (Don't forget to your `.env`).

## Building a production container

The `thebe-server` also runs on your rackspace cluster as a docker container.  Since we're using go, we can use [@ngauthier](https://twitter.com/ngauthier)'s great article [Building Minimal Docker Containers for Go Applications](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/) to create a super slim container image that has only our comiled binary and `.env` file.  (NB: it will probably also need to have docker 1.6 binary at some point!).  

Basically, we first build a static binary and add it to the super minimalist `scratch` image from Docker.  Here are the commands:

```
GOOS=linux go build -a -installsuffix cgo -o thebe-server .
docker build -t thebe-server .
```

As specified in the `Dockerfile`, the build process will put the `thebe-server` binary and a file called `.env.prod`, which holds the production environment variables, onto the image.

If you're not using Go 1.5, consult [Cross Compiling Go](http://stackoverflow.com/questions/27412601/cross-compiling-go) on Stackoverflow for information on how to compile a binary for earlier versions of Go.


## Starting `thebe-server` on the swarm

Once you've built the container, you're ready to start `thebe-server` on the swarm.  Since we'll be issuing docker commands against the swarm itself, we need to add the `--volumes-from swarm-data` option in the `docker run` command to include the swarm credentials automatically.  This will mount our swarm's credentials in the `/etc/docker` directory, so in your `.env.prod` file you'd set `SWARM_CREDS_DIR=/etc/docker`.

Finally, we're ready to start our service on the swarm:

```
docker run -d  \
   -p 8000 \
   -P \
   --volumes-from swarm-data \
   --hostname thebe-server.interlock.odewahn.com \
   thebe-server
```

Once it starts, you should see it pop up in the HAProxy stats page, just like the other containers you're running there.  However, in this instance, we're using  fixed hostname to serve as the endpoint for the API.
