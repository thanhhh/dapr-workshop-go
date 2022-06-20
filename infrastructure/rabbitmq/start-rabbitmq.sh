#!/usr/bin/env bash
set -euo pipefail

arch -arm64 docker run -d -p 5672:5672 -p 15672:15672 --name dtc-rabbitmq rabbitmq:3-management-alpine