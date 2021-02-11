# Go Exercise

Simple go program to practice with go essentials

## Overview
Some producers produce messages that are consumed by consumers

## Run

    # go-routines version

    go run ./routines

    # grpc version
    # generate file
    cd ./grpc && ./gen.sh && cd -
    go run ./grpc/producer
    go run ./grpc/consumer

## GRPC

Instalation instructions [here](https://grpc.io/docs/protoc-installation/)
