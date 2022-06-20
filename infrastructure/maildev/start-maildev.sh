#!/usr/bin/env bash
set -euo pipefail

arch -arm64 docker run -d -p 1080:1080 -p 1025:1025 --name dtc-maildev maildev/maildev:latest