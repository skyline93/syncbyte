package backup

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/skyline93/syncbyte-go/internal/engine/config"
	"github.com/skyline93/syncbyte-go/internal/engine/grpc"
	"github.com/skyline93/syncbyte-go/internal/engine/schema"
	"github.com/skyline93/syncbyte-go/pkg/mongodb"
)

func Backup(sourcePath string) error {
	mongoClient, err := mongodb.NewClient(config.Conf.Core.MongodbUri)
	if err != nil {
		panic(err)
	}
	defer mongoClient.Close()

	fiChan := make(chan schema.FileInfo)

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	go grpc.Backup(fiChan, sourcePath, ctx)

	col := mongoClient.GetCollection("backup01")

	for fi := range fiChan {
		if _, err := col.InsertOne(context.TODO(), fi); err != nil {
			log.Errorf("insert error, err: %v", err)
			continue
		}

		log.Debugf("fi: %s", fi.String())
	}

	return nil
}
