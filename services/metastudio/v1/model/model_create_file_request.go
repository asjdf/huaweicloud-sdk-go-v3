package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFileRequest Request Object
type CreateFileRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间.格式为(YYYYMMDD'T'HHMMSS'Z')。 取值为当前系统的GMT时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 开发者应用作为资产权属的可选字段。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	Body *FilesCreateReq `json:"body,omitempty"`
}

func (o CreateFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFileRequest struct{}"
	}

	return strings.Join([]string{"CreateFileRequest", string(data)}, " ")
}
