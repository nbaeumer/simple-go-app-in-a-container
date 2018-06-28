#!/bin/bash
VERSION=v3

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./go/simple-app ./go/main.go
docker build -t simpleapp-scratch:$VERSION .
docker images
rm simple-app


#Push Docker images
export _DOCKER_REPO="$(aws ecr get-authorization-token --profile nbaeumer-gameday --output text  --query 'authorizationData[].proxyEndpoint')"
aws ecr get-login --no-include-email --region eu-west-1 --profile nbaeumer-gameday | awk '{print $6}' | docker login -u AWS --password-stdin $_DOCKER_REPO
docker tag my_app_repo:latest 196987642544.dkr.ecr.eu-west-1.amazonaws.com/my_app_repo:latest
docker push 196987642544.dkr.ecr.eu-west-1.amazonaws.com/simple-app:latest
