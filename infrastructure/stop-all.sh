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
podman stop dapr_zipkin
podman stop dapr_redis

podman rm dapr_zipkin
podman rm dapr_redis

# podman network rm dapr-network
