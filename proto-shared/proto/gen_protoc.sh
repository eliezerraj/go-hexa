#!/bin/bash

mkdir -p ../generated/go
protoc rate/rate.proto --go_out=plugins=grpc:../generated/go