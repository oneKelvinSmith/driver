## Zombie Driver

A small service for finding out if we have any zombie drivers.

### Testing

This project uses [ginko](https://onsi.github.io/ginkgo) and [gomega](https://onsi.github.io/gomega/) testing libraries.

To run the tests use:
```
$ go test
```

### Building

To build the `zombie` binary use:
```
$ go build
```

To build the release image with docker run:
```
$ docker build -t zombie .
```

### Running

To run the service locally use:
```
$ ./zombie
```

Or with the release image:
```
$ docker run --rm -it -p 3000:3000 zombie
```
