#!/bin/bash

SCRIPT_PATH=$(dirname "$(realpath -s "$0")")

go build -o plots-left "$SCRIPT_PATH/main.go"
env GOOS=linux GOARCH=amd64 go build -o plots-left-linux-amd64 "$SCRIPT_PATH/main.go"
env GOOS=darwin GOARCH=amd64 go build -o plots-left-darwin-amd64 "$SCRIPT_PATH/main.go"
env GOOS=linux GOARCH=arm64 go build -o plots-left-linux-arm64 "$SCRIPT_PATH/main.go"
