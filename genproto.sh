#! /bin/bash
cd grpc
protoc  --micro_out=. --go_out=. *.proto