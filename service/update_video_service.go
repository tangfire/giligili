package service

import (
	"singo/model"
	"singo/serializer"
)

type UpdateVideoService struct {
	Title string `json:"title" form:"title" binding:"required,min=2,max=30"`
	Info  string `json:"info" form:"info" binding:"max=300"`
}

func (service *UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}

	video.Title = service.Title
	video.Info = service.Info
	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  serializer.CodeDBError,
			Msg:   "视频保存失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}

}
