package repository

import (
	"time"

	"github.com/skyline93/syncbyte-go/internal/pkg/types"
	"gorm.io/gorm"
)

type DBResource struct {
	gorm.Model
	Name     string
	DBType   string
	Version  string
	Server   string
	Port     int
	User     string
	Password string
	DBName   string
	Args     string
}

type S3Backend struct {
	gorm.Model
	EndPoint    string
	AccessKey   string
	SecretKey   string
	Bucket      string
	StorageType string
	DataType    string
}

type BackupPolicy struct {
	gorm.Model
	ResourceID uint
	Retention  int
	IsCompress bool `gorm:"default:false"`
	AgentID    uint
	Status     string

	ScheduleType string
	Cron         string
	Frequency    int
	StartTime    types.LocalTime
	EndTime      types.LocalTime
	IsActive     bool `gorm:"default:true"`
}

type BackupJob struct {
	gorm.Model
	StartTime  time.Time
	EndTime    time.Time
	Status     string
	ResourceID uint
	BackendID  uint
	PolicyID   uint
}

type BackupSet struct {
	gorm.Model
	DataSetName string
	IsCompress  bool
	IsValid     bool `gorm:"default:false"`
	Size        int
	BackupJobID uint
	BackupTime  time.Time
	ResourceID  uint
	BackendID   uint
	Retention   int
}

type RestoreJob struct {
	gorm.Model
	StartTime   time.Time
	EndTime     time.Time
	Status      string
	BackupSetID uint
}

type RestoreDBResource struct {
	gorm.Model
	Name         string
	DBType       string
	Version      string
	Server       string
	Port         int
	User         string
	Password     string
	DBName       string
	Args         string
	RestoreJobID uint
	IsValid      bool `gorm:"default:false"`
	RestoreTime  time.Time
}

type Agent struct {
	gorm.Model
	IP       string
	Port     int
	HostName string
	HostType string
}

type ScheduledJob struct {
	gorm.Model

	JobType        string
	Status         string
	BackupPolicyID uint
}
