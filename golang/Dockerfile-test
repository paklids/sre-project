## leverage a build container and copy binaries to final container
FROM golang:1.15-alpine AS build
## need git for pulling golang packages
RUN set -ex; \
    apk update; \
    apk add --no-cache git
WORKDIR /src
COPY src /src
RUN go get gopkg.in/redis.v4
RUN go build -o /out/go_binary .

## final container where binaries live
FROM alpine:latest AS bin
COPY --from=build /out/go_binary /
WORKDIR /