FROM golang:1.18-alpine AS builder

WORKDIR /go/src/github.com/skyline93/syncbyte-go

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o syncbyte-agent cmd/agent/*.go \
    && CGO_ENABLED=0 go build -o syncbyte-engine cmd/engine/*.go

FROM alpine:latest AS restic

RUN apk add --update --no-cache ca-certificates fuse openssh-client tzdata \
    && mkdir -p /var/syncbyte/data /var/syncbyte/log

VOLUME [ "/var/syncbyte/data", "/var/syncbyte/log" ]

COPY --from=builder /go/src/github.com/skyline93/syncbyte-go/syncbyte-agent /usr/bin
COPY --from=builder /go/src/github.com/skyline93/syncbyte-go/syncbyte-engine /usr/bin
