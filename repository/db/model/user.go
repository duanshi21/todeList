package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
	"todoList/consts"
)

// User 用户模型
type User struct {
	Id        int64      `gorm:"column:id;primary_key"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	Username  string     `gorm:"column:user_name;unique"`
	Password  string     `gorm:"column:password"`
}

func (*User) TableName() string {
	return "user"
}

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), consts.PasswordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
