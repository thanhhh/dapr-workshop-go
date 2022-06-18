arch -arm64 podman network create dapr-network
arch -arm64 podman run --name "dapr_zipkin" --network dapr-network --restart always -d -p 9411:9411 openzipkin/zipkin 
arch -arm64 podman run --name "dapr_redis" --network dapr-network --restart always -d -p 6379:6379 redis
arch -arm64 podman run --name "dtc-rabbitmq" --network dapr-network -d -p 5672:5672 -p 15672:15672  rabbitmq:management-alpine
arch -arm64 dapr init --slim --network dapr-network
$HOME/.dapr/bin/placement