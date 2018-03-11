### Build

FROM bitwalker/alpine-elixir:1.6.3 as builder

WORKDIR /gateway

ENV MIX_ENV=prod

COPY mix.exs mix.lock ./
COPY config/* config/

RUN mix do deps.get, deps.compile

COPY . .

RUN mix release --env=prod --verbose


### Release

FROM alpine:3.7

RUN apk upgrade --no-cache && \
    apk add --no-cache bash openssl-dev

WORKDIR /service

ENV PORT=3000 \
    MIX_ENV=prod \
    REPLACE_OS_VARS=true \
    SHELL=/bin/bash

EXPOSE $PORT

COPY --from=builder /gateway/_build/prod/rel/gateway/releases/latest/gateway.tar.gz .
RUN tar zxf gateway.tar.gz && rm gateway.tar.gz

RUN chown -R root ./releases

USER root

CMD ["/service/bin/gateway", "foreground"]
