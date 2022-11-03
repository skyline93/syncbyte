package backup

import (
	"context"

	"github.com/skyline93/syncbyte-go/internal/engine/config"
	"github.com/skyline93/syncbyte-go/internal/engine/grpc"
	"github.com/skyline93/syncbyte-go/internal/engine/schema"
	"github.com/skyline93/syncbyte-go/pkg/logging"
	"github.com/skyline93/syncbyte-go/pkg/mongodb"
)

var logger = logging.GetSugaredLogger("backup")

func Backup(sourcePath, mountPoint string) error {
	mongoClient, err := mongodb.NewClient(config.Conf.MongodbUri)
	if err != nil {
		panic(err)
	}
	defer mongoClient.Close()

	fiChan := make(chan schema.FileInfo)

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	go grpc.Backup(fiChan, sourcePath, mountPoint, ctx)

	col := mongoClient.GetCollection("backup01")

	for fi := range fiChan {
		if _, err := col.InsertOne(context.TODO(), fi); err != nil {
			logger.Errorf("insert error, err: %v", err)
			continue
		}

		logger.Debugf("fi: %s", fi.String())
	}

	return nil
}
