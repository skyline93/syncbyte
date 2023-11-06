package syncbyte

import (
	"context"

	"github.com/skyline93/ctask"
	"gorm.io/gorm"
)

type Worker struct {
	ctx    context.Context
	broker ctask.Broker
	db     *gorm.DB
}

func NewWorker(ctx context.Context, broker ctask.Broker, db *gorm.DB) *Worker {
	return &Worker{
		ctx:    ctx,
		broker: broker,
		db:     db,
	}
}

func (w *Worker) Run() {
	srv := ctask.NewServer(w.broker, ctask.Config{
		Concurrency: 100,
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
	})

	srv.Handle(TypeBackup, NewBackupHandler(w.db))

	srv.Run(w.ctx)
}
