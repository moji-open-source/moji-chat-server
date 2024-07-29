package models

type SysUserModel struct {
	Email    string
	Name     string
	Password string
}

func (*SysUserModel) TableName() string {
	return "sys_user"
}

type UserRepository interface {
	GetByEmail(email string) SysUserModel
}
