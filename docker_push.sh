#!/bin/bash

echo "$DOCKER_PASSWORD" | docker login --username "$DOCKER_USERNAME" --password-stdin
docker tag etiennenapoleone/tomochain etiennenapoleone/tomochain:latest
docker tag etiennenapoleone/tomochain etiennenapoleone/tomochain:$(git log --pretty=format:'%h' -n 1 | cat)
docker push etiennenapoleone/xdcchain:latest
docker push etiennenapoleone/xdcchain:$(git log --pretty=format:'%h' -n 1 | cat)
