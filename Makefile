GO111MODULE=on
GONAME=$(shell basename $(CURDIR))
SRC=${CURDIR}/src/${GONAME}
GOFILES=$(GONAME).go

default: build

build:
	./scripts/build.sh

start: build
	bin/heatpump-x86 -db heatpump -source 192.168.84.45 -target http://localhost:8086 -username heatpump -password heatpump -verbose
