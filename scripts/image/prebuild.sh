#!/bin/bash

ROOTDIR=$(dirname $0)/../..

if [ -d "build" ]; then
        rm -rf build
fi
mkdir -p build

cp ${ROOTDIR}/bin/metadatad build/
