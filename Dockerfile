FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/OctaneAL/swift-code-api
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/swift-code-api /go/src/github.com/OctaneAL/swift-code-api


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/swift-code-api /usr/local/bin/swift-code-api
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["swift-code-api"]
