GO111MODULE=on
GONAME=$(shell basename $(CURDIR))
SRC=${CURDIR}/src/${GONAME}
GOFILES=$(GONAME).go

default: build

build:
	./scripts/build.sh

start: build
	bin/heatpump -verbose
