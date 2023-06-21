package model

import (
	"time"
)

type Task struct {
	Id        int64      `gorm:"column:id;primary_key"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	Uid       int64      `gorm:"column:uid"`
	Title     string     `gorm:"column:title"`
	Status    int        `gorm:"column:status"`
	Content   string     `gorm:"column:content"`
	StartTime int64      `gorm:"column:start_time"`
	EndTime   int64      `gorm:"column:end_time"`
}

func (*Task) TableName() string {
	return "task"
}
