package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNextflowJobLogRequest Request Object
type ShowNextflowJobLogRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`
}

func (o ShowNextflowJobLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNextflowJobLogRequest struct{}"
	}

	return strings.Join([]string{"ShowNextflowJobLogRequest", string(data)}, " ")
}
