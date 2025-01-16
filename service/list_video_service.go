package service

import (
	"singo/model"
	"singo/serializer"
)

type ListVideoService struct {
}

func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video

	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Code:  serializer.CodeDBConnectError,
			Msg:   "数据库连接失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}

}
