build:
	CGO_ENABLED=0 go build -o syncbyte-agent cmd/agent/*.go
	CGO_ENABLED=0 go build -o syncbyte-engine cmd/engine/*.go
	CGO_ENABLED=0 go build -o syncbytectl cmd/syncbytectl/*.go

build-image:
	docker build -t syncbyte:latest .

compose:
	mkdir -p data/appdata data/applog data/pgdata data/mongodata data/sourcedata
	docker compose up -d

clean:
	rm -rf syncbyte-agent syncbyte-engine syncbytectl
