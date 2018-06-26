BUILD := ${PWD}/_build

build:
	set -x

	cd controller/main;              docker build -t controller          .
	cd broker/provider/aws/project;  docker build -t broker.aws.project  .
	cd broker/provider/aws/environ;  docker build -t broker.aws.environ  .
	cd broker/provider/aws/database; docker build -t broker.aws.database .
	cd broker/provider/aws/cache;    docker build -t broker.aws.cache    .
	cd broker/provider/aws/compute;  docker build -t broker.aws.compute  .

	docker images

up:
	set -x

	docker-compose up -d
	docker ps

down:
	set -x

	docker-compose down
	docker ps

catalog:
	set -x

	curl -s localhost:9081/v1/catalog | jq .
	curl -s localhost:9082/v1/catalog | jq .
	curl -s localhost:9083/v1/catalog | jq .
	curl -s localhost:9084/v1/catalog | jq .
	curl -s localhost:9085/v1/catalog | jq .

test:
	set -x

	curl -sX POST localhost:9080/v1/register -d '{"url": "http://interstellar_broker.aws.project_1:8080"}'  | jq .
	curl -sX POST localhost:9080/v1/register -d '{"url": "http://interstellar_broker.aws.environ_1:8080"}'  | jq .
	curl -sX POST localhost:9080/v1/register -d '{"url": "http://interstellar_broker.aws.database_1:8080"}' | jq .
	curl -sX POST localhost:9080/v1/register -d '{"url": "http://interstellar_broker.aws.cache_1:8080"}'    | jq .
	curl -sX POST localhost:9080/v1/register -d '{"url": "http://interstellar_broker.aws.compute_1:8080"}'  | jq .

	curl -s localhost:9080/v1/service | jq .
	curl -s localhost:9080/v1/service/$(shell curl -s localhost:9080/v1/service | jq -r '.service[0].service_id') | jq .

package:
	set -x
	-rm -rf ${BUILD}

	mkdir -p ${BUILD}/controller/bin
	mkdir -p ${BUILD}/broker/{cache,compute,database,project,environ}/{bin,conf}

	cd controller/main;              go build -o ${BUILD}/controller/bin/controller
	cd broker/provider/aws/project;  go build -o ${BUILD}/broker/project/bin/project
	cd broker/provider/aws/environ;  go build -o ${BUILD}/broker/environ/bin/environ
	cd broker/provider/aws/database; go build -o ${BUILD}/broker/database/bin/database
	cd broker/provider/aws/cache;    go build -o ${BUILD}/broker/cache/bin/cache
	cd broker/provider/aws/compute;  go build -o ${BUILD}/broker/compute/bin/compute

	cp broker/provider/aws/project/template.yml  ${BUILD}/broker/project/conf
	cp broker/provider/aws/environ/template.yml  ${BUILD}/broker/environ/conf
	cp broker/provider/aws/database/template.yml ${BUILD}/broker/database/conf
	cp broker/provider/aws/cache/template.yml    ${BUILD}/broker/cache/conf
	cp broker/provider/aws/compute/template.yml  ${BUILD}/broker/compute/conf

prune:
	set -x

	docker image prune --force
	docker images

.PHONY:
