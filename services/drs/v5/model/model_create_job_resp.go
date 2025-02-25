package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateJobResp 创建任务响应体。
type CreateJobResp struct {

	// 任务ID。
	Id string `json:"id"`

	// 任务名称。
	Name string `json:"name"`

	// 任务状态。
	Status string `json:"status"`

	// 任务创建时间。
	CreateTime string `json:"create_time"`
}

func (o CreateJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobResp struct{}"
	}

	return strings.Join([]string{"CreateJobResp", string(data)}, " ")
}
