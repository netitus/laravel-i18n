#!/bin/bash

docker rm -v "go-builder" || true
docker build --tag="go-builder" .
docker create --rm --name "go-builder-bin" "go-builder"
docker cp "go-builder-bin":/app/bin/ ./
docker rm -v "go-builder-bin"
