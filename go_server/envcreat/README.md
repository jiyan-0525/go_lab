# Creating an Environment

You may be thinking, "What's the point of dockerizing this simple service"? Well, at the moment, there are only a couple of benefits:

Anyone with Docker can run your image, regardless of their OS
You can easily deploy containers of your image on any cloud service that uses images (most of them) or on an orchestration server like Kubernetes.
If your server were written in a language like Python or JavaScript, you could bundle the interpreter and dependencies inside the image so that you don't need to reconfigure them on the server.
That said, because our app is so simple, there's just not much environment required, and one of the best things about Docker is that it allows you to ship an entire environment.

So... let's make it more interesting!

Assignment
We're going to make the port that our server binds to configurable: it will be set by an environment variable.

In main.go, find the line that sets port to a hard-coded value of 8010 and update it so that it reads an environment variable called PORT. You can use 

[![os.Getenv Documentation](https://img.shields.io/badge/Standard_Lib-os.Getenv-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://pkg.go.dev/os#Getenv)
```
port := os.Getenv("PORT")
```
Make sure that the os package is imported:
```
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)
```
Change the port to 8999 by setting an environment variable in your shell:
```
export PORT="8999"
```
Rebuild and run your Go program and make sure it serves on port 8999. If you still have GOOS and GOARCH set from the last lesson, unset them first. Remember to replace goserver with the name of your binary.
bash
```
go build
./goserver
```
Add an ENV command to your Dockerfile to set the port within the image. You'll need to do it before the CMD command so that the environment variable is set before the server starts.
```
ENV PORT=8991
```
Rebuild your Docker image:
If you're not on Linux, rebuild the Go binary for Linux first: GOOS=linux GOARCH=amd64 
bash
```
go build

docker build . -t goserver:latest
```
Rerun your Docker container, be sure to expose the correct port:
bash
```
docker run -p 8991:8991 goserver
```