# interstellar


## up

```
$ make docker-up

docker-compose up -d
Starting interstellar_broker.aws.database_1 ... done
Starting interstellar_broker.aws.cache_1    ... done
Starting interstellar_controller_1          ... done
Starting interstellar_broker.aws.project_1  ... done
Starting interstellar_broker.aws.environ_1  ... done
Starting interstellar_broker.aws.compute_1  ... done

docker ps
CONTAINER ID        IMAGE                 COMMAND             CREATED             STATUS                  PORTS                    NAMES
6496b2b82531        broker.aws.project    "./app"             16 hours ago        Up Less than a second   0.0.0.0:9081->8080/tcp   interstellar_broker.aws.project_1
4f4a2e899f11        controller            "./app"             16 hours ago        Up 1 second             0.0.0.0:9080->8080/tcp   interstellar_controller_1
e3a161c90d38        broker.aws.compute    "./app"             16 hours ago        Up 1 second             0.0.0.0:9085->8080/tcp   interstellar_broker.aws.compute_1
b9b4bb01dd30        broker.aws.database   "./app"             16 hours ago        Up 1 second             0.0.0.0:9083->8080/tcp   interstellar_broker.aws.database_1
ebd97426f6fc        broker.aws.cache      "./app"             16 hours ago        Up 1 second             0.0.0.0:9084->8080/tcp   interstellar_broker.aws.cache_1
80f2e90cbeac        broker.aws.environ    "./app"             16 hours ago        Up Less than a second   0.0.0.0:9082->8080/tcp   interstellar_broker.aws.environ_1
```

## register

```
$ make docker-register

curl -sX POST localhost:9080/v1/register -d '{"url": "http://interstellar_broker.aws.project_1:8080"}'  | jq .
{
  "status": 200,
  "service_id": "c7aa976d-79a3-11e8-8e2f-0242ac120004"
}

curl -sX POST localhost:9080/v1/register -d '{"url": "http://interstellar_broker.aws.environ_1:8080"}'  | jq .
{
  "status": 200,
  "service_id": "c7b22741-79a3-11e8-8e2f-0242ac120004"
}

curl -sX POST localhost:9080/v1/register -d '{"url": "http://interstellar_broker.aws.database_1:8080"}' | jq .
{
  "status": 200,
  "service_id": "c7b5f849-79a3-11e8-8e2f-0242ac120004"
}

curl -sX POST localhost:9080/v1/register -d '{"url": "http://interstellar_broker.aws.cache_1:8080"}'    | jq .
{
  "status": 200,
  "service_id": "c7bb9e9f-79a3-11e8-8e2f-0242ac120004"
}

curl -sX POST localhost:9080/v1/register -d '{"url": "http://interstellar_broker.aws.compute_1:8080"}'  | jq .
{
  "status": 200,
  "service_id": "c7bf8401-79a3-11e8-8e2f-0242ac120004"
}
```


## service

```
$ make service

curl -s localhost:9080/v1/service | jq .
{
  "status": 200,
  "service": [
    {
      "name": "aws_project",
      "service_id": "c7aa976d-79a3-11e8-8e2f-0242ac120004",
      "service_broker_url": "http://interstellar_broker.aws.project_1:8080"
    },
    {
      "name": "aws_environ",
      "service_id": "c7b22741-79a3-11e8-8e2f-0242ac120004",
      "service_broker_url": "http://interstellar_broker.aws.environ_1:8080"
    },
    {
      "name": "aws_database",
      "service_id": "c7b5f849-79a3-11e8-8e2f-0242ac120004",
      "service_broker_url": "http://interstellar_broker.aws.database_1:8080"
    },
    {
      "name": "aws_cache",
      "service_id": "c7bb9e9f-79a3-11e8-8e2f-0242ac120004",
      "service_broker_url": "http://interstellar_broker.aws.cache_1:8080"
    },
    {
      "name": "aws_compute",
      "service_id": "c7bf8401-79a3-11e8-8e2f-0242ac120004",
      "service_broker_url": "http://interstellar_broker.aws.compute_1:8080"
    }
  ]
}

curl -s localhost:9080/v1/service/c7aa976d-79a3-11e8-8e2f-0242ac120004 | jq .
{
  "status": 200,
  "service_id": "c7aa976d-79a3-11e8-8e2f-0242ac120004",
  "catalog": {
    "name": "aws_project",
    "display_name": "",
    "description": "",
    "tag": [
      "aws",
      "project"
    ],
    "require": [],
    "bindable": false,
    "parameter": [
      {
        "name": "aws_account_id",
        "required": true,
        "default_value": "",
        "allowed_value": null,
        "description": ""
      },
      {
        "name": "integration_role_arn",
        "required": false,
        "default_value": "",
        "allowed_value": null,
        "description": ""
      },
      {
        "name": "project_name",
        "required": true,
        "default_value": "",
        "allowed_value": null,
        "description": ""
      },
      {
        "name": "cidr",
        "required": true,
        "default_value": "",
        "allowed_value": null,
        "description": ""
      },
      {
        "name": "domain",
        "required": true,
        "default_value": "",
        "allowed_value": null,
        "description": ""
      }
    ]
  }
}
```


## create


```
$ make create

curl -s X POST  localhost:9080/v1/instance -d '{"service_id": "c7aa976d-79a3-11e8-8e2f-0242ac120004", "name": "develop01", "parameter": {"project_name": "myproject", "cidr": "10.1.0.0/16", "region": "ap-northeast-1"}' | jq .
{
  "status": 200,
  "instance": {
    "name": "develop01",
    "service_id": "c7aa976d-79a3-11e8-8e2f-0242ac120004",
    "instance_id": "f5fadf4f-79a3-11e8-8e2f-0242ac120004",
    "output": {
      "nameserver": "ns-1,ns-2,ns-3,ns-4"
    }
  }
}

$ make instance

curl -s localhost:9080/v1/instance| jq .
{
  "status": 200,
  "instance": [
    {
      "name": "develop01",
      "service_id": "c7aa976d-79a3-11e8-8e2f-0242ac120004",
      "instance_id": "f5fadf4f-79a3-11e8-8e2f-0242ac120004",
      "output": {
        "nameserver": "ns-1,ns-2,ns-3,ns-4"
      }
    }
  ]
}
```
