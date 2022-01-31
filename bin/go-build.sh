#!/bin/bash

docker run -it --rm -v "$PWD/go":/go -v "$PWD":/app -w /app golang:1.17.6 go build -o log-service