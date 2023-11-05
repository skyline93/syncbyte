run:
	mkdir -p .testdata/pgdata
	docker compose up -d

clean-testdata:
	rm -rf .testdata
