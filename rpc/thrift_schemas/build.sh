#!/usr/bin/env bash

for f in $(ls *.thrift); do
    thrift --out .. --gen go:thrift_import=github.com/mshockwave/thrift-go/thrift,package_prefix=github.com/ShareSound/RPC-Server/rpc/ ${f}
done
