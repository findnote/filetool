#!/bin/sh
rm mwp3000

go build api/mwp3000.go

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" api/mwp3000.go
#http-server

# run upx
upx toproxy


