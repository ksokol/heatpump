#!/usr/bin/env bash
set -e

export GO111MODULE=on

rm -rf bin/heatpump

CGO_ENABLED=0 GOGC=off go build -o bin/heatpump ./cmd/heatpump
