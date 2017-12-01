#!/bin/bash

if [[ -z $DOCKER_GO_PRODUCER_TARGET_REPO ]]; then
	echo "DOCKER_GO_PRODUCER_TARGET_REPO not defined"
	echo "You need to supply a docker.io repo as a docker push target"
	exit 1
fi

docker build -t go-producer:latest . \
	&& docker tag go-producer:latest $DOCKER_GO_PRODUCER_TARGET_REPO \
	&& docker push $DOCKER_GO_PRODUCER_TARGET_REPO

