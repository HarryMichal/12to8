#!/bin/bash -xe

SUBDIRS=("api" "cmd" "helpers")

for dir in "${SUBDIRS[@]}"; do
    go test -v "./${dir}/..."
done
