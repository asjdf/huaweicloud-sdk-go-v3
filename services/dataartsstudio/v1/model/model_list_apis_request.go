package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApisRequest Request Object
type ListApisRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 查询起始坐标, 即跳过前X条数据。仅支持0或limit的整数倍，不满足则向下取整
	Offset *int32 `json:"offset,omitempty"`

	// 查询条数, 即查询Y条数据
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListApisRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisRequest struct{}"
	}

	return strings.Join([]string{"ListApisRequest", string(data)}, " ")
}
