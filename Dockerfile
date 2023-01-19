FROM golang:1.18-alpine AS builder

WORKDIR /go/src/github.com/skyline93/syncbyte-go

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o agent cmd/agent/*.go \
    && CGO_ENABLED=0 go build -o engine cmd/engine/*.go

FROM alpine:latest AS restic

RUN apk add --update --no-cache ca-certificates fuse openssh-client tzdata

COPY --from=builder /go/src/github.com/skyline93/syncbyte-go/agent /usr/bin
COPY --from=builder /go/src/github.com/skyline93/syncbyte-go/engine /usr/bin
