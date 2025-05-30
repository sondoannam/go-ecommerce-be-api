package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"column:uuid; primaryKey; type:varchar(50); not null; index:idx_uuid; unique"`
	Username string    `gorm:"column:user_name"`
	IsActive bool      `gorm:"column:is_active; type:boolean"`
	Roles    []Role    `gorm:"many2many:user_roles"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
