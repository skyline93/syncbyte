package entity

import (
	"database/sql"
	"time"
)

// type User struct {
// 	ID   int    `gorm:"primary_key" json:"ID" yaml:"-"`
// 	Name string `gorm:"size:200;index;" json:"Name" yaml:"Name,omitempty"`
// }

// func (User) TableName() string {
// 	return "users"
// }

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // 一个常规字符串字段
	Email        *string        // 一个指向字符串的指针, allowing for null values
	Age          uint8          // 一个未签名的8位整数
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // 创建时间（由GORM自动管理）
	UpdatedAt    time.Time      // 最后一次更新时间（由GORM自动管理）
}
