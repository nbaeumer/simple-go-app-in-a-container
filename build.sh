#!/bin/bash
VERSION=v3

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o simple-app ./main.go
docker build -t simpleapp-scratch:$VERSION .
docker images
rm simple-app