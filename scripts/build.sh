#!/usr/bin/env bash
set -e

export GO111MODULE=on

rm -rf bin/heatpump*

CGO_ENABLED=0 GOGC=off go build -o bin/heatpump-x86 ./cmd/heatpump

export GOARCH=arm
export GOARM=6

CGO_ENABLED=0 GOGC=off go build -o bin/heatpump-arm ./cmd/heatpump
