#!/usr/bin/env bash
# linux server build go-default-service images
# author: weiqiang; date: 2022-12
set -x
set -e

registrieAddress="harbor.xxx.com/devops"
version=$1
if [ -z $version ]; then
    version=v0.2
fi

imageTagVersion=${version}
servicename=go-default-service
pkgname=go-default-service

export GOARCH=amd64
export GOOS=linux
export GCCGO=gc

go build -o build/${pkgname} main.go
chmod u+x build/${pkgname}
tar -zcvf ${pkgname}-linux-amd64-${version}.tar.gz \
  build/${pkgname} configs/config.yaml configs/${pkgname}.service README.md web/

docker build -f build/dockerfile -t ${registrieAddress}/${servicename}:${imageTagVersion} .
docker push ${registrieAddress}/${servicename}:${imageTagVersion}
