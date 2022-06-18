#!/usr/bin/env bash
set -euo pipefail

arch -arm64 podman run -d -p 4000:80 -p 4025:25 --name dtc-maildev maildev/maildev:latest