# My First Go App in a Container

### GO: SimpleApp to make requests to addition to numbers
Call http://*YOURSERVER*/additon?num1=1&num2=2 for addition.

Call http://*YOURSERVER*/addition?num1=1&num2=2&delay=100 for addition and a feedback delay of 100 Millisecond

### Build & Deploy & Run
[Micro Docker Images for Go Applications](https://blog.kloia.com/micro-docker-images-for-go-applications-8a8701130c01)

GOOS=linux GOARCH=amd64 go build -o simple-app ./main.go

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o simple-app ./main.go

docker build -t *IMAGE-NAME* . //build the docker image

docker run --rm -p 8080:8080 *IMAGE-NAME* //run the container

docker run --rm -it --entrypoint=/bin/bash *IMAGE-NAME* //run the container and look inside the container
