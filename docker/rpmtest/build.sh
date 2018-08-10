#!/bin/bash

cp Dockerfile ../../dist
cd ../../dist

docker build --tag "logdisplayplatform/rpmtest" .

rm Dockerfile

docker run -i -t logdisplayplatform/rpmtest /bin/bash
