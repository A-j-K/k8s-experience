#!/bin/bash

if [[ -z $DOCKER_GO_CONSUMER_TARGET_REPO ]]; then
	echo "DOCKER_GO_CONSUMER_TARGET_REPO not defined"
	echo "You need to supply a docker.io repo as a docker push target"
	exit 1
fi

docker build -t go-consumer:latest . \
	&& docker tag go-consumer:latest $DOCKER_GO_CONSUMER_TARGET_REPO \
	&& docker push $DOCKER_GO_CONSUMER_TARGET_REPO

