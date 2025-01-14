// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameProjectComment = "project_comments"

// ProjectComment mapped from table <project_comments>
type ProjectComment struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	Text      string    `gorm:"column:text;not null" json:"text"`
	UserID    string    `gorm:"column:user_id;not null" json:"user_id"`
	InReplyTo int64     `gorm:"column:in_reply_to" json:"in_reply_to"`
	ProjectID string    `gorm:"column:project_id;not null" json:"project_id"`
}

// TableName ProjectComment's table name
func (*ProjectComment) TableName() string {
	return TableNameProjectComment
}
