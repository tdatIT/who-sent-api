package models

import (
	"time"
)

type Role struct {
	ID          int        `gorm:"id;primaryKey;autoIncrement;type:int;" json:"id"`
	Name        string     `gorm:"name;type:varchar(25);not null;unique;" json:"name"`
	Description string     `gorm:"description;type:varchar(100)" json:"description"`
	Status      int        `gorm:"status;type:tinyint;" json:"status"`
	CreatedAt   time.Time  `gorm:"created_at;type:datetime;autoCreateTime;" json:"createdAt"`
	UpdatedAt   *time.Time `gorm:"updated_at;type:datetime;autoUpdateTime;" json:"updatedAt"`
	// Relationships
}

func (*Role) TableName() string {
	return "role"
}
