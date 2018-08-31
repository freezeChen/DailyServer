#! /bin/bash
cd rpc/proto
protoc  --micro_out=. --go_out=. user.proto