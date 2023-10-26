FROM golang:1.20.2-bullseye

WORKDIR /go/src/github.com/skyline93/syncbyte-go

COPY . .

RUN GOPROXY=https://proxy.golang.com.cn,direct CGO_ENABLED=0 go build -o syncbyte-agent cmd/agent/*.go \
    && CGO_ENABLED=0 go build -o syncbyte-engine cmd/engine/*.go \
    && CGO_ENABLED=0 go build -o syncbytectl cmd/syncbytectl/*.go \
    && mv syncbyte-agent /usr/local/bin \
    && mv syncbyte-engine /usr/local/bin \
    && mv syncbytectl /usr/local/bin \
    && rm -rf * \
    && mkdir -p /var/syncbyte/data /var/syncbyte/log

CMD [ "/bin/bash" ]
