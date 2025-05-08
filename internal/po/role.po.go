package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID              int    `gorm:"column:id; primaryKey; type:int; not null; autoIncrement; required"`
	RoleName        string `gorm:"column:role_name; type:varchar(50); not null; unique"`
	RoleDescription string `gorm:"column:role_description; type:text"`
}

func (r *Role) TableName() string {
	return "go_db_role"
}
