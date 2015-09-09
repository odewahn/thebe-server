#


Use `--volumes-from swarm-data` to include the swarm credentials automatically.  The certs will be mounted to the `/etc/docker` directory

```
docker run -d  \
   -p 8888 \
   -P \
   --volumes-from swarm-data \
   --hostname sneaky.interlock.odewahn.com \
   ipython/scipystack    \
   /bin/sh -c 'ipython notebook --ip=0.0.0.0 --no-browser'
```
