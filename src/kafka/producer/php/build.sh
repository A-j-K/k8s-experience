#!/bin/bash

if [[ -z $DOCKER_PHP_PRODUCER_TARGET_REPO ]]; then
	echo "DOCKER_PHP_PRODUCER_TARGET_REPO not defined"
	echo "You need to supply a docker.io repo as a docker push target"
	exit 1
fi

docker build -t php-producer:latest . \
	&& docker tag php-producer:latest $DOCKER_PHP_PRODUCER_TARGET_REPO \
	&& docker push $DOCKER_PHP_PRODUCER_TARGET_REPO

