#!/bin/sh

_logdisplayplatform_tag=$1

# If the tag starts with v, treat this as a official release
if echo "$_logdisplayplatform_tag" | grep -q "^v"; then
	_logdisplayplatform_version=$(echo "${_logdisplayplatform_tag}" | cut -d "v" -f 2)
	_docker_repo=${2:-logdisplayplatform/logdisplayplatform}
else
	_logdisplayplatform_version=$_logdisplayplatform_tag
	_docker_repo=${2:-logdisplayplatform/logdisplayplatform-dev}
fi

echo "Building ${_docker_repo}:${_logdisplayplatform_version}"

docker build \
	--tag "${_docker_repo}:${_logdisplayplatform_version}" \
	--no-cache=true .

# Tag as 'latest' for official release; otherwise tag as logdisplayplatform/logdisplayplatform:master
if echo "$_logdisplayplatform_tag" | grep -q "^v"; then
	docker tag "${_docker_repo}:${_logdisplayplatform_version}" "${_docker_repo}:latest"
else
	docker tag "${_docker_repo}:${_logdisplayplatform_version}" "logdisplayplatform/logdisplayplatform:master"
fi
