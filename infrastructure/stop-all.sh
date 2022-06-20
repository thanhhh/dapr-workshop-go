#!/usr/bin/env bash
set -euo pipefail

basedir=$(dirname $0)

pushd ${basedir}/rabbitmq
./stop-rabbitmq.sh
popd

pushd ${basedir}/mosquitto
./stop-mosquitto.sh
popd

pushd ${basedir}/maildev
./stop-maildev.sh
popd

# dapr uninstall --network dapr-network
docker stop dapr_zipkin
docker stop dapr_redis

docker rm dapr_zipkin
docker rm dapr_redis

dapr uninstall
# docker network rm dapr-network
