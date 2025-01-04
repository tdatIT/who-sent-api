package models

import (
	"fmt"
	"github.com/tdatIT/who-sent-api/pkgs/utils/genid"
	"gorm.io/gorm"
	"time"
)

const (
	UserActiveStatus = true
	UserInActive     = false
)

type User struct {
	ID                     int            `gorm:"primaryKey,type:bigint" json:"id"`
	Firstname              string         `gorm:"type:varchar(50)" json:"firstname"`
	Lastname               string         `gorm:"type:varchar(50)" json:"lastname"`
	Phone                  string         `gorm:"type:varchar(20)" json:"phone"`
	Gender                 int            `gorm:"type:int;default:0" json:"gender"`
	IsActivated            bool           `gorm:"type:bool;default:false" json:"is_activated"`
	IsVerified             bool           `gorm:"type:bool;default:false" json:"is_verified"`
	AvatarUrl              string         `gorm:"avatar_url;type:varchar(255)" json:"avatar_url"`
	Email                  string         `gorm:"type:varchar(150);not null;index" json:"email"`
	RequiredChangePassword bool           `gorm:"type:bool;default:false" json:"required_change_password"`
	Password               string         `gorm:"type:varchar(255);not null" json:"password"`
	Version                int            `gorm:"type:int;default:0" json:"version"`
	CreatedAt              time.Time      `gorm:"created_at;autoUpdateTime" json:"created_at"`
	UpdatedAt              time.Time      `gorm:"updated_at;autoUpdateTime;" json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"index"`
	// Relationships
	Roles []Role `gorm:"many2many:user_roles;" json:"-"`
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

func (i *User) GetUserIdStr() string {
	return fmt.Sprintf("%d", i.ID)
}

func (i *User) GetRolesStringSlice() []string {
	roles := make([]string, 0)
	for _, role := range i.Roles {
		roles = append(roles, role.Name)
	}
	return roles
}
