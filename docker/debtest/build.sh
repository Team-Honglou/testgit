#!/bin/bash

cp Dockerfile ../../dist
cd ../../dist

docker build --tag "logdisplayplatform/debtest" .

rm Dockerfile

docker run -i -t logdisplayplatform/debtest /bin/bash
