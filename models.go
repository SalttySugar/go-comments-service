package service

import "time"

type Comment struct {
	ID             uint64     	`gorm:"primaryKey" json:"id"`
	Message   	   string     	`json:"message"`
	PublisherId    string     	`json:"publisher_id"`
	RecordId       string     	`json:"record_id"`
	PostendOn      *time.Time 	`json:"posted_on"`
	UpdatedOn      *time.Time 	`json:"updated_on"`
}

type CommentDto struct {
	PublisherId   string `json:"publisher_id"`
	RecordId      string `json:"record_id"`
	Message       string `json:"message"`
}
