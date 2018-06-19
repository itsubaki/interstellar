
BUILD := ${PWD}/_build

build:
	set -x
	-rm -rf ${BUILD}

	mkdir -p ${BUILD}/bin
	mkdir -p ${BUILD}/broker/cache/bin
	mkdir -p ${BUILD}/broker/compute/bin
	mkdir -p ${BUILD}/broker/database/bin
	mkdir -p ${BUILD}/broker/network/bin

	cd launcher;                     go build -o ${BUILD}/bin/launcher
	cd broker/provider/aws/cache;    go build -o ${BUILD}/broker/cache/bin/cache
	cd broker/provider/aws/compute;  go build -o ${BUILD}/broker/compute/bin/compute
	cd broker/provider/aws/database; go build -o ${BUILD}/broker/database/bin/database
	cd broker/provider/aws/network;  go build -o ${BUILD}/broker/network/bin/network

.PHONY:
