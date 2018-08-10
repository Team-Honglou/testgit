#!/bin/sh
set -e

_logdisplayplatform_tag=$1

# If the tag starts with v, treat this as a official release
if echo "$_logdisplayplatform_tag" | grep -q "^v"; then
	_logdisplayplatform_version=$(echo "${_logdisplayplatform_tag}" | cut -d "v" -f 2)
	_docker_repo=${2:-logdisplayplatform/logdisplayplatform}
else
	_logdisplayplatform_version=$_logdisplayplatform_tag
	_docker_repo=${2:-logdisplayplatform/logdisplayplatform-dev}
fi

echo "pushing ${_docker_repo}:${_logdisplayplatform_version}"
docker push "${_docker_repo}:${_logdisplayplatform_version}"

if echo "$_logdisplayplatform_tag" | grep -q "^v"; then
	echo "pushing ${_docker_repo}:latest"
	docker push "${_docker_repo}:latest"
else
	echo "pushing logdisplayplatform/logdisplayplatform:master"
	docker push logdisplayplatform/logdisplayplatform:master
fi
