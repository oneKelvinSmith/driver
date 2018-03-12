### Build

FROM golang:alpine as builder

WORKDIR /go/src/zombie

RUN apk upgrade --no-cache && \
    apk add --no-cache bash git

RUN go get -u github.com/golang/dep/cmd/dep

COPY . .

RUN dep ensure

RUN go build


### Release

FROM alpine:3.7

WORKDIR /service

ENV PORT=3002

EXPOSE $PORT

COPY --from=builder /go/src/zombie/zombie .

CMD ["/service/zombie"]
