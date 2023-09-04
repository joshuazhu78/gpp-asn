#!/bin/sh

proto_imports=".:${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate:${GOPATH}/src"

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api:${GOPATH}/src/github.com/joshuazhu78/gpp-asn/protos \
  --proto_path=pkg/nr \
  --go_out=./pkg/nr/ \
  nr-rrc-ies.proto
protoc-go-inject-tag -input=./pkg/nr/nr-rrc-ies/nr-rrc-ies.pb.go
