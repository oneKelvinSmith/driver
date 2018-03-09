## Driver Location

A small service for retrieving the location data for a driver.

### Testing

This project uses [ginko](https://onsi.github.io/ginkgo) and [gomega](https://onsi.github.io/gomega/) testing libraries.

To run the tests use:
```
$ go test
```

### Building

To build the `location` binary use:
```
$ go build
```

To build the release image with docker run:
```
$ docker build -t location .
```

### Running

To run the service locally use:
```
$ ./location
```

Or with the release image:
```
$ docker run --rm -it -p 3000:3000 location
```
