#!/usr/bin/env bash
# linux server build go-default-service images
set -x
set -e

imageTagVersion="v0.1"
registrieAddress="harbor.xxx.com/devops"
servicename=go-default-service
pkgname=go-default-service

export GOARCH=amd64
export GOOS=linux
export GCCGO=gc

go build -o build/${pkgname} main.go
chmod u+x build/${pkgname}

docker build -f build/dockerfile -t ${registrieAddress}/${servicename}:${imageTagVersion} .
docker push ${registrieAddress}/${servicename}:${imageTagVersion}
