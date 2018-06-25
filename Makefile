
BUILD := ${PWD}/_build

build:
	set -x

	cd broker/provider/aws/project;  docker build -t broker.aws.project .
	cd broker/provider/aws/environ;  docker build -t broker.aws.environ .
	cd broker/provider/aws/database; docker build -t broker.aws.database .
	cd broker/provider/aws/cache;    docker build -t broker.aws.cache .
	cd broker/provider/aws/compute;  docker build -t broker.aws.compute .

up:
	set -x

	docker run -d --rm -p 9080:8080 --name project  broker.aws.project
	docker run -d --rm -p 9081:8080 --name environ  broker.aws.environ
	docker run -d --rm -p 9082:8080 --name database broker.aws.database
	docker run -d --rm -p 9083:8080 --name cache    broker.aws.cache
	docker run -d --rm -p 9084:8080 --name compute  broker.aws.compute
	docker ps

down:
	set -x

	docker stop $(shell docker ps -a -q)
	docker ps

build-bin:
	set -x
	-rm -rf ${BUILD}

	mkdir -p ${BUILD}/interstellar/bin
	mkdir -p ${BUILD}/broker/cache/bin
	mkdir -p ${BUILD}/broker/compute/bin
	mkdir -p ${BUILD}/broker/database/bin
	mkdir -p ${BUILD}/broker/project/bin
	mkdir -p ${BUILD}/broker/environ/bin

	cd launcher/interstellar;        go build -o ${BUILD}/interstellar/bin/interstellar
	cd broker/provider/aws/project;  go build -o ${BUILD}/broker/project/bin/project
	cd broker/provider/aws/environ;  go build -o ${BUILD}/broker/environ/bin/environ
	cd broker/provider/aws/database; go build -o ${BUILD}/broker/database/bin/database
	cd broker/provider/aws/cache;    go build -o ${BUILD}/broker/cache/bin/cache
	cd broker/provider/aws/compute;  go build -o ${BUILD}/broker/compute/bin/compute

.PHONY:
