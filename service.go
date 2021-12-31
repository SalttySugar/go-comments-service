package service

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type GormCommentsService struct {
	connection *gorm.DB
}

func (service GormCommentsService) FindById(id string) Comment {
	var comment *Comment
	if err := service.connection.First(&comment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(ErrCommentNotFound(id));
		}
		panic(err);
	}

	return *comment
}

func (service GormCommentsService) FindAll() []Comment {
	var comments []Comment

	if err:=service.connection.Find(&comments).Error; err != nil {
		panic(err);
	}
	return comments
}

func (service GormCommentsService) DeleteById(id string) {
	if err := service.connection.Delete(service.FindById(id)).Error; err !=nil {
		panic(err);
	}
}

func (service GormCommentsService) Update(id string, message string) Comment {
	comment := service.FindById(id)
	comment.Message = message
	comment.UpdatedOn = &time.Time{}

	if err := service.connection.Save(&comment).Error; err != nil {
		panic(err);
	}

	return comment
}

func (service GormCommentsService) Create(payload CommentDto) Comment {
	comment := Comment{
		Message:  payload.Message,
		RecordId: payload.RecordId,
		PublisherId:   payload.PublisherId,
		PostendOn: &time.Time{},
	}

	if err := service.connection.Create(&comment).Error; err != nil {
		panic(err);
	}

	return comment
}


var gormCommentsServiceInstance CommentsService


func NewGormCommentsService(conn *gorm.DB) CommentsService {
	if gormCommentsServiceInstance != nil {
		return gormCommentsServiceInstance
	}
	gormCommentsServiceInstance = GormCommentsService{connection: conn}
	return gormCommentsServiceInstance
}