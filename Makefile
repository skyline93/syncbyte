build-all: swag build-go

build-go:
	go build -o syncbyte cmd/syncbyte/syncbyte.go
	go build -o syncbyte-agent cmd/agent/agent.go

swag:
	swag init -g api.go --dir ./internal/api/ -o ./internal/api/docs
	swag fmt -g ./internal/api/api.go

clean:
	rm -rf syncbyte syncbyte-agent* logs/*
