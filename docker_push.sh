#!/bin/bash

echo "$DOCKER_PASSWORD" | docker login --username "$DOCKER_USERNAME" --password-stdin
docker tag xdcchain/xdcchain xdcchain/xdcchain:latest
docker tag xdcchain/xdcchain xdcchain/xdcchain:$(git log --pretty=format:'%h' -n 1 | cat)
docker push xdcchain/xdcchain:latest
docker push xdcchain/xdcchain:$(git log --pretty=format:'%h' -n 1 | cat)
