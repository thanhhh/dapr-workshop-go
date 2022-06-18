dapr uninstall --network dapr-network
podman stop dapr_zipkin
podman stop dapr_redis
podman stop dtc-rabbitmq

podman rm dapr_zipkin
podman rm dapr_redis
podman rm dtc-rabbitmq

podman network rm dapr-network
