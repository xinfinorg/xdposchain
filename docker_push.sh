#!/bin/bash

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push etiennenapoleone/xdcchain:latest
docker push etiennenapoleone/xdcchain:$(git log --pretty=format:'%h' -n 1 | cat)
