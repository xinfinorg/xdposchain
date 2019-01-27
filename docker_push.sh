#!/bin/bash

echo "$DOCKER_PASSWORD" | docker login --username "$DOCKER_USERNAME" --password-stdin
docker tag xdcchain/xdcchain xdcchain/xdcchain:latest
docker tag xdcchain/xdcchain xdcchain/xdcchain:$TRAVIS_BUILD_ID
docker push xdcchain/xdcchain:latest
docker push xdcchain/xdcchain:$TRAVIS_BUILD_ID
