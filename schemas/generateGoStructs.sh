#!/bin/sh

echo "Starting proto to struct..."

protoc -I=. --go_out=../api/src *.proto
protoc -I=. --go_out=../worker/src *.proto

protoc -I=. --go_out=../app/src *.proto

echo "Completed proto to struct..."
