#!/bin/env bash

buf generate
protoc-go-inject-tag -input=./src/gen/synapsis/v1/*.pb.go -remove_tag_comment
