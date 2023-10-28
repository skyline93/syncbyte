package syncbyte

import "time"

type BackupSet struct {
	ID         uint
	IsValid    bool
	Size       int64
	BackupTime *time.Time
	ResourceID uint
	Retention  int
}
