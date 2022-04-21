#!/bin/bash

protoc greet/greetpb/greet.proto --go-grpc_out=greet
