#!/usr/bin/env bash
set -euo pipefail

arch -arm64 podman run -d -p 1883:1883 -p 9001:9001 --name dtc-mosquitto dapr-trafficcontrol/mosquitto:1.0