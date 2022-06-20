#!/usr/bin/env bash
set -euo pipefail

docker run -d -p 1080:1080 -p 1025:1025 --name dtc-maildev maildev/maildev:latest