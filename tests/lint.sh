#!/bin/bash -xe

SUBDIRS=("./" "api" "cmd" "helpers" "tests")

for dir in "${SUBDIRS[@]}"; do
    go vet "./${dir}/..."
done
