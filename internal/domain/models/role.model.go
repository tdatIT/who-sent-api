package models

import (
	"time"
)

type Role struct {
	ID          int        `gorm:"id;primaryKey;autoIncrement;type:int;" json:"id"`
	Name        string     `gorm:"name;type:varchar(25);not null;unique;" json:"name"`
	Description string     `gorm:"description;type:varchar(100)" json:"description"`
	Status      int        `gorm:"status;type:smallint;" json:"status"`
	CreatedAt   time.Time  `gorm:"autoCreateTime;" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime;" json:"updated_at"`
	// Relationships
}

func (Role) TableName() string {
	return "roles"
}
