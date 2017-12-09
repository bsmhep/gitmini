#!/bin/sh

cd $(dirname "$0")

go build gitmini.go && tar -czf gitmini.tar.gz gitmini
