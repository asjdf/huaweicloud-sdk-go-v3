package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopVideoMotionCaptureJobRequest Request Object
type StopVideoMotionCaptureJobRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。格式为(YYYYMMDD'T'HHMMSS'Z')
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 开发者应用作为资产权属的可选字段。
	XAppUserId *string `json:"X-App-UserId,omitempty"`
}

func (o StopVideoMotionCaptureJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopVideoMotionCaptureJobRequest struct{}"
	}

	return strings.Join([]string{"StopVideoMotionCaptureJobRequest", string(data)}, " ")
}
