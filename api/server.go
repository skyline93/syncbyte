package api

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/skyline93/ctask"
	"github.com/skyline93/syncbyte/syncbyte"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	engine *gin.Engine
}

func NewServer(db *gorm.DB) *Server {
	engine := gin.Default()
	return &Server{db: db, engine: engine}
}

func (s *Server) Run(addr string) error {
	s.engine.Use(func(ctx *gin.Context) {
		ctx.Set("gormdb", s.db)
		ctx.Next()
	})

	s.engine.POST("/backup/policies", createBackupPolicy)
	s.engine.POST("/backup/jobs", createBackupJob)

	return s.engine.Run(addr)
}

type CreateBackupPolicyRequest struct {
	ResourceIndentifier string                `json:"resource_indentifier"`
	ResourceType        syncbyte.ResourceType `json:"resource_type"`
	ResourceAttr        datatypes.JSONMap     `json:"resource_attr"`
	Retention           int                   `json:"retention"`
}

func createBackupPolicy(c *gin.Context) {
	var req CreateBackupPolicyRequest
	gormDB, ok := c.Get("gormdb")
	if !ok {
		return
	}

	db := gormDB.(*gorm.DB)

	if err := c.BindJSON(&req); err != nil {
		return
	}

	pl, err := syncbyte.CreateBackupPolicy(db, req.ResourceIndentifier, req.ResourceType, req.ResourceAttr, req.Retention)
	if err != nil {
		return
	}

	log.Printf("pl_id: %d, resource: %s", pl.ID, pl.Resource.Identifier)
}

type CreateBackupJobRequest struct {
	BackupPolicyID uint `json:"backup_policy_id"`
}

func createBackupJob(c *gin.Context) {
	var req CreateBackupJobRequest
	gormDB, ok := c.Get("gormdb")
	if !ok {
		return
	}

	db := gormDB.(*gorm.DB)

	if err := c.BindJSON(&req); err != nil {
		return
	}

	pl, err := syncbyte.GetBackupPolicy(db, int(req.BackupPolicyID))
	if err != nil {
		return
	}

	bj, err := syncbyte.CreateBackupJob(db, pl)
	if err != nil {
		log.Printf("err: %v", err)
		return
	}

	task, err := syncbyte.NewBackupTask(bj.ID, pl.ResourceID, pl.Retention)
	if err != nil {
		return
	}

	sendTask(task)

	log.Printf("job_id: %d", bj.ID)
}

func sendTask(task *ctask.Task) {
	client := ctask.NewClient(ctask.NewRDBBroker(context.Background(), &redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	}))
	defer client.Close()

	info, err := client.Enqueue(context.TODO(), task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}

	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}
