package main

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/skyline93/ctask"
	"github.com/skyline93/syncbyte/api"
	"github.com/skyline93/syncbyte/syncbyte"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = syncbyte.InitDB("host=localhost user=syncbyte password=123456 dbname=syncbyte port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	broker := ctask.NewRDBBroker(ctx, &redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})

	worker := syncbyte.NewWorker(ctx, broker, DB)

	go worker.Run()

	srv := api.NewServer(DB)
	err := srv.Run("0.0.0.0:8000")
	if err != nil {
		panic(err)
	}
}
