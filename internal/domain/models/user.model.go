package models

import (
	"github.com/tdatIT/who-sent-api/pkgs/utils/genid"
	"gorm.io/gorm"
	"time"
)

const (
	UserActiveStatus = true
	UserInActive     = false
)

type User struct {
	ID          int        `gorm:"primaryKey,type:bigint" json:"id"`
	Firstname   string     `gorm:"type:varchar(50)" json:"firstname"`
	Lastname    string     `gorm:"type:varchar(50)" json:"lastname"`
	Phone       string     `gorm:"type:varchar(20)" json:"phone"`
	Gender      int        `gorm:"type:int;default:0" json:"gender"`
	IsActivated bool       `gorm:"type:tinyint;default:1" json:"is_activated"`
	IsVerified  bool       `gorm:"type:tinyint;default:0" json:"is_verified"`
	AvatarUrl   string     `gorm:"avatar_url;type:varchar(255)" json:"avatar_url"`
	Email       string     `gorm:"type:varchar(150);not null;index" json:"email"`
	Password    string     `gorm:"type:varchar(255);not null" json:"password"`
	Version     int        `gorm:"type:int;default:0" json:"version"`
	CreatedAt   time.Time  `gorm:"created_at;type:datetime;" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"updated_at;type:datetime;" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"deleted_at;type:datetime;" json:"deleted_at"`
	// Relationships
	Roles []Role `gorm:"many2many:user_roles;"`
}

func (i *User) TableName() string {
	return "users"
}

func (i *User) BeforeCreate(tx *gorm.DB) error {
	i.ID = int(genid.GetSnowLakeIns().Generate().Int64())
	return nil
}

func (i *User) BeforeUpdate(tx *gorm.DB) (err error) {
	i.Version = i.Version + 1
	return nil
}
