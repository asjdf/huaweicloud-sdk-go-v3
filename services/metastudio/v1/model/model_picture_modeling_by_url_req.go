package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PictureModelingByUrlReq 任务请求
type PictureModelingByUrlReq struct {

	// 图片URL
	PictureUrl string `json:"picture_url"`

	// 风格ID
	StyleId string `json:"style_id"`

	// 风格ID
	Name *string `json:"name,omitempty"`

	// 照片建模任务结束的回调地址。
	NotifyUrl *string `json:"notify_url,omitempty"`
}

func (o PictureModelingByUrlReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PictureModelingByUrlReq struct{}"
	}

	return strings.Join([]string{"PictureModelingByUrlReq", string(data)}, " ")
}
