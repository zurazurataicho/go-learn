# How to bulid and run
1. go mod init zura.org/docker-environment
1. vi main.go
1. docker build -t my-golang-app .
1. docker run -it --rm --name my-running-app my-golang-app
