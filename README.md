# Driver

The services composing this application are found in subdirectories:

- `gateway` an elixir service, that either forward or transform requests to be processed synchronously or asynchronously.
- `location` a go service, that consumes location update events and store them.
- `zombie` a go service, that allows to check if a driver matches the zombie predicate or not.

Each service has it's own readme with instructions for getting them up and ruuning

## Assumptions

This project has been built on OSX and the instructions assume a `unix` environment with some mac specific references.

## Prerequisites

You will need to have the following tools installed to run locally:

- [Docker](https://www.docker.com) (I've used [Docker for Mac](https://www.docker.com/docker-mac))
- [Docker Compose](https://docs.docker.com/compose/install/)

## Development

As some of these services are `go` packages, you'll need to have set your [`$GOPATH`](https://github.com/golang/go/wiki/GOPATH) this repo cloned to `$GOPATH/src/driver`.

Please ensure you have the following language versions installed:

```bash
$ go version
go version go1.9.4 darwin/amd64
```

```bash
$ elixir -v
Erlang/OTP 20 [erts-9.1] [source] [64-bit] [smp:4:4] [ds:4:4:10] [async-threads:10] [hipe] [kernel-poll:false]

Elixir 1.6.1 (compiled with OTP 19)
```

### Running the cluster

To get the compose cluster up an running use:

```bash
$ ./start.sh
...
```

Too bring the cluster down use:

```bash
$ ./stop.sh
...
```

To see the list of running containers user:

```bash
$ docker-compose ps
...
```

To see the container logs use:

```bash
$ docker-compose logs
...
```

The `gateway` service will be available locally via: [localhost:3000](http://localhost:3000/)
The `location` service will be available locally via: [localhost:3001](http://localhost:3001/)
The `zombie` service will be available locally via: [localhost:3002](http://localhost:3002/)
The NSQ admin page will be available locally via: [localhost:4171](http://localhost:4171/)

### Verification

I've encluded a [postman](https://www.getpostman.com/) collection file `Driver.postman_collection.json` with some test requests set up. This can be used to test the endpoints.
