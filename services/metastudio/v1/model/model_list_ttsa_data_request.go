package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTtsaDataRequest Request Object
type ListTtsaDataRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。格式为(YYYYMMDD'T'HHMMSS'Z')
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 偏移量，表示生成内容时间偏移，目前每次返回2S内容
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListTtsaDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTtsaDataRequest struct{}"
	}

	return strings.Join([]string{"ListTtsaDataRequest", string(data)}, " ")
}
