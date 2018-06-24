
BUILD := ${PWD}/_build

build:
	set -x
	-rm -rf ${BUILD}

	mkdir -p ${BUILD}/interstellar/bin
	mkdir -p ${BUILD}/broker/cache/bin
	mkdir -p ${BUILD}/broker/compute/bin
	mkdir -p ${BUILD}/broker/database/bin
	mkdir -p ${BUILD}/broker/project/bin
	mkdir -p ${BUILD}/broker/environ/bin

	cd launcher/interstellar;        go build -o ${BUILD}/interstellar/bin/interstellar
	cd broker/provider/aws/cache;    go build -o ${BUILD}/broker/cache/bin/cache
	cd broker/provider/aws/compute;  go build -o ${BUILD}/broker/compute/bin/compute
	cd broker/provider/aws/database; go build -o ${BUILD}/broker/database/bin/database
	cd broker/provider/aws/environ;  go build -o ${BUILD}/broker/environ/bin/environ
	cd broker/provider/aws/project;  go build -o ${BUILD}/broker/project/bin/initialize

.PHONY:
