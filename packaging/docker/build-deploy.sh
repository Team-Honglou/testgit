#!/bin/sh
set -e

_logdisplayplatform_version=$1
./build.sh "$_logdisplayplatform_version"
docker login -u "$DOCKER_USER" -p "$DOCKER_PASS"

./push_to_docker_hub.sh "$_logdisplayplatform_version"

if echo "$_logdisplayplatform_version" | grep -q "^master-"; then
  apk add --no-cache curl
  ./deploy_to_k8s.sh "logdisplayplatform/logdisplayplatform-dev:$_logdisplayplatform_version"
fi
