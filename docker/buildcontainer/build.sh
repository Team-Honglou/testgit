#!/bin/bash

docker kill gfbuild
docker rm gfbuild

docker build --tag "logdisplayplatform/buildcontainer" docker/buildcontainer

docker run -i -t \
  -v /home/torkel/dev/go:/go \
  --name gfbuild logdisplayplatform/buildcontainer
