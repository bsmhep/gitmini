#!/bin/sh

cd $(dirname "$0")

go build -ldflags '-s' gitmini.go && tar -czf gitmini.tar.gz gitmini
