#! /bin/sh -x

# A convenience script for starting the cluster.

# This script assumes you're on OSX and using Docker for Mac to use the useful hostname.
# See https://docs.docker.com/docker-for-mac/networking/

# Build the container images.
docker build -t gateway ./gateway/
docker-compose build

# Run the cluster.
docker-compose up -d

# Unfortunatly the elixir container is not starting up properly with docker compose.
# So we run the container independently.
docker run --rm -id -p 3000:3000 \
                    -e PORT=3000 \
                    -e COOKIE=cookie \
                    -e NSQD_TOPIC=driver \
                    -e NSQD_HOST=docker.for.mac.host.internal:4150 \
                    -e ZOMBIE_HOST=docker.for.mac.host.internal:3002 \
                    --name gateway \
                    gateway
