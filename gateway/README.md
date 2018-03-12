## Driver Gateway

A service for retrieving and updating the location data of a driver.


This service relies on the [Mix build tool](https://elixir-lang.org/getting-started/mix-otp/introduction-to-mix.html) for building and testing the application.

### Testing

To run the tests use:
```
$ mix test
```

### Building

To compile the `gateway` use:
```
$ mix deps.get
$ mix compile
```

Or with docker:
```
$ docker build -t gateway .
```

### Running

To run the service locally use:
```
$ mix run --no-halt
```

Or with docker (with the other services running compose cluster):

```
$ docker run --rm -it -p 3000:3000 \
                      -e PORT=3000 \
                      -e COOKIE=cookie \
                      -e NSQD_TOPIC=driver \
                      -e NSQD_HOST=docker.for.mac.host.internal:4150 \
                      -e ZOMBIE_HOST=docker.for.mac.host.internal:3002 \
                         gateway
```
