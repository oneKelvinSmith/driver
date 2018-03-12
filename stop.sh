#! /bin/sh -x

# A convenience script for stopping the running containers.

# This assumes you have started the cluster and gateway container using the start.sh script.
docker stop gateway
docker-compose down

