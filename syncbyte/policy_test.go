package syncbyte

import (
	"testing"

	"github.com/skyline93/syncbyte/syncbyte/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestBackupPolicy(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&Resource{}, &BackupPolicy{})

	identity := "testresource"
	resourceType := ResourceTypeDatabase
	retention := 7

	mysql := database.NewMySQL("MYDB", 3306)
	attr, err := mysql.GetAttr()
	if err != nil {
		t.FailNow()
	}

	pl, err := CreateBackupPolicy(db, identity, resourceType, attr, retention)
	assert.NoError(t, err)
	assert.Equal(t, retention, pl.Retention)
	assert.Equal(t, identity, pl.Resource.Identifier)
}
