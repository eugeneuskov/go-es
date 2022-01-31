#!/bin/bash

docker run -it --rm -v "$PWD/go":/go -v "$PWD":/app -w /app -p 11020:11020 golang:1.17.6 go run main.go