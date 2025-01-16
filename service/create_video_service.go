package service

import (
	"singo/model"
	"singo/serializer"
)

// CreateVideoService 视频投稿的服务
type CreateVideoService struct {
	Title string `json:"title" form:"title" binding:"required,min=2,max=30"`
	Info  string `json:"info" form:"info" binding:"max=300"`
}

// Create 创建视频
func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
	}
	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  serializer.CodeDBError,
			Msg:   "视频保存错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
