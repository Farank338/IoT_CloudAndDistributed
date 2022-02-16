#!/bin/sh

cd src &&\
go get -u
go mod tidy
go build
./hw2


