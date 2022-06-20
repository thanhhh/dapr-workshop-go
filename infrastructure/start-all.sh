#!/usr/bin/env bash

set -euo pipefail

basedir=$(dirname $0)

pushd ${basedir}/rabbitmq
./start-rabbitmq.sh
popd

pushd ${basedir}/mosquitto
./start-mosquitto.sh
popd

pushd ${basedir}/maildev
./start-maildev.sh
popd

# arch -arm64 docker network create dapr-network
arch -arm64 docker run --name "dapr_zipkin" --restart always -d -p 9411:9411 openzipkin/zipkin 
arch -arm64 docker run --name "dapr_redis"  --restart always -d -p 6379:6379 redis
arch -arm64 dapr init --slim