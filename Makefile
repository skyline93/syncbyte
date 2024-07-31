swag:
	swag init -g api.go --dir ./internal/api/ -o ./internal/api/docs
	swag fmt -g ./internal/api/api.go
