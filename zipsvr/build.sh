#!/usr/bin/env bash
# Tell command line: chmod +x build.sh
set -e # If there is an error inany of these commands, stop everything
GOOS=linux go build
docker build -t leemeli/zipserver .
docker push leemeli/zipserver
go clean