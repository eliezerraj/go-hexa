#!/bin/bash

mkdir -p ../generated/go
protoc rate/rate.proto --go_out=../generated/go