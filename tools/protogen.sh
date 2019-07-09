#!/usr/bin/env bash

if [ -z "$1" ]; then
    echo "No argument supplied"
    exit 1
fi

TARGET_PATH=$1

echo "Target: ${TARGET_PATH}"

# grpc proto
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  ${TARGET_PATH}

# proxy proto
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  ${TARGET_PATH}

echo "Compiled protocol buffer..."
