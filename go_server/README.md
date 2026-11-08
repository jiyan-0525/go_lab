# 🌐 goserver 🚀

run
```
go mod init github.com/USERNAME/goserver
```
after replacing USERNAME and goserver with your own values.

Build the server: 
```
go build
```
This will create a standalone binary file in your current directory (the server application).

Run the server: 
```
./goserver
```
where goserver is the name of the built binary; by default it's the same as your project's directory.

Make sure you can hit your server's endpoint(s) either in a browser or with curl. It should be available on localhost:8010.
```
http://localhost:8010
```

# 📦Dockerizing the Server
Now that you know how to run your server manually, let's run it in Docker! The steps are simple:

Build the Server
Create a Dockerfile
Build an image using the Dockerfile (which will copy in the built server)
Run the image in a container
Assignment

Create a Dockerfile in the root of your server's repo. Let's start with a simple lightweight Debian Linux OS:
Add a COPY command on the next line of your Dockerfile. In the case of a simple compiled Go server, all we need is the compiled program itself!

Replace the first "goserver" with the name of your server executable if it's different.

The ADD command would also work here, but COPY is fine because we don't need the extra functionality that ADD offers.

Add a CMD command as the last line in the Dockerfile. This automatically starts the server process in the container when we run it.
```
FROM debian:stable-slim

COPY goserver /bin/goserver

CMD ["/bin/goserver"]
```
Build your Dockerfile into an image:

```
docker build . -t goserver:latest
```

Start a new container from the image. Be sure to forward the ports to your host machine.

```
docker run -p 8010:8010 goserver
```
If you get an exec format error, it's probably because you built the go server for your local architecture, but you're trying to run it on a Linux OS! To fix it, rebuild the binary (and then the Dockerfile) with these environment variables:

```
GOOS=linux GOARCH=amd64 go build
```
You should be able to access your server from the browser just like before, but this time it's running inside of Docker!

```
http://localhost:8010
```
## Publishing to Docker Hub

Rebuild the Go binary:

```
GOOS=linux GOARCH=amd64 go build
```

Rebuild the image. You'll need to use a name that corresponds to your namespace on Docker Hub. Swap out USERNAME for your Docker Hub username.
```
docker build . -t USERNAME/goserver
```

```
docker login

docker tag goserver jiyanwang/goserver
docker push jiyanwang/goserver
```
```
docker run -p 8991:8991 jiyanwang/goserver

docker push jiyanwang/goserver
```

Remove your local USERNAME/goserver image:

```
docker image rm jiyanwawng/goserver
```
Pull it back down from Docker Hub:

```
docker pull jiyanwang/goserver
```
Run it to make sure it works:

```
docker run -p 8991:8991 jiyanwang/goserver
```
