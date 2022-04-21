#!/bin/bash

protoc greet/greetpb/greet.proto --go-grpc_out=greet --go_out=greet

protoc calculator/calculatorpb/calculator.proto --go-grpc_out=calculator --go_out=calculator
