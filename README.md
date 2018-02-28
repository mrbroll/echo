This is a basic go implementation of an http echo server. This is just a quick utility for learning by testing for things like service to service communication in a docker environment, as well as understanding how docker networking works through hands-on experimentation.

Build
=====

Prerequisites:
- Go (tested using 1.9, but any version _should_ work)
- Docker (duh)

In a shell:
```
go get github.com/mrbroll/echo
cd $GOPATH/src/github.com/mrbroll/echo
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build .
docker build -t echo .
```

Run
===

```
docker run --rm --name echo echo
```
or if you want to forward the port to your localhost:
```
docker run --rm --name echo -p 3000:80 echo
```
Then, you can make a request to the server:
```
curl localhost:3000 -d 'hello, world!'
```

Example Usage
=============

The reason I built this was to test out how the networking within a docker container related to the host machine, and I wanted to see if I could contact services bound to the host from within the container without using the host networking option for `docker run`. To re-create this, complete the build and run steps above. To avoid opening multiple terminals, you can run the echo container in the background:

```
docker run --rm --name echo -p 3000:80 -d echo
```
Then run our client container:
```
docker run -it --rm --name curl appropriate/curl /bin/sh
```
This will start the container and attach to its tty with a bourne shell running. From here, you can inspect the ip addresses associated with the container:
```
ip addr show
```
You should see an `eth0` interface with a CIDR of something like `172.17.0.3/16`. The least significant octet might be different depending on your environment. Using basic linux networking knowledge, we know that the gateway for this network is `172.17.0.1`. So we should be able to reach any services bound to the host machine through this:
```
curl 172.17.0.1:3000 -d 'hello, world!'
```
and you should see your input echoed back from the service.
