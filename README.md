# Driver

Coming soon:

- a `gateway` service, that either forward or transform requests to be processed synchronously or asynchronously
- a `driver location` service, that consumes location update events and store them
- a `zombie driver` service, that allows to check if a driver matches the zombie predicate or not

## Assumptions

This project has been built on OSX and the instructions assume a `unix` environment with some mac specific references.

## Prerequisites

You will need to have the following tools installed to run locally:

- [Docker](https://www.docker.com) (I've used [Docker for Mac](https://www.docker.com/docker-mac))
- [Docker Compose](https://docs.docker.com/compose/install/)

## Development

As some of these services are `go` packages, you'll need to have set your [`$GOPATH`](https://github.com/golang/go/wiki/GOPATH) this repo cloned to `$GOPATH/src/driver`.

### Running the cluster

The get [NSQ](http://nsq.io/) binaries up an running use:

```bash
$ docker-compose up -d
> ...
```

Which will build and run a cluster of containers in a private network.

To see the list of running containers user:

```bash
$ docker-compose ps
> ...
```

To see the container logs use:

```bash
$ docker-compose logs
> ...
```

And to bring the cluster down use:

```bash
$ docker-compose down
> ...
```

The NSQ admin page will be available locally via: [localhost:4171](http://localhost:4171/)
