#!/usr/bin/env bash
# Tell command line: chmod +x build.sh
set -e # If there is an error inany of these commands, stop everything
GOOS=linux go build
docker build -t leemeli/testserver .
docker push leemeli/testserver
go clean