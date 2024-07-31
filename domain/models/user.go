package models

import (
	"time"

	"gorm.io/gorm"
)

type Gender = string

const (
	Girl    Gender = "Girl"
	Boy     Gender = "Boy"
	Unknown Gender = "Unknown"
)

type UserStatus = string

const (
	Normal       UserStatus = "Normal"
	Ban          UserStatus = "Ban"
	NotActivated UserStatus = "NotActivated "
)

type SysUserModel struct {
	Uid      int64
	UserCode string
	Avatar   string
	Password string
	Email    string
	NickName string
	Gender
	Birth    time.Time
	Status   UserStatus
	CreateAt time.Time
}

func (*SysUserModel) TableName() string {
	return "sys_user"
}

type UserRepository interface {
	GetByEmail(email string) (SysUserModel, error)
}

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		database: db,
	}
}

func (r *userRepository) GetByEmail(email string) (SysUserModel, error) {
	var user SysUserModel
	result := r.database.First(&user, "email = ?", email)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
