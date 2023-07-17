package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowForwardingConfigsRequest Request Object
type ShowForwardingConfigsRequest struct {

	// **参数说明**：实例ID。dris物理实例的唯一标识。获取方法参见[获取Instance-Id](https://support.huaweicloud.com/api-v2x/v2x_04_0030.html)。  **取值范围**：仅支持数字，小写字母和连接符（-）的组合，长度36。
	InstanceId *string `json:"Instance-Id,omitempty"`

	// **参数说明**：转发配置的类型。  **取值范围**：当前仅支持“kafka，mrskafka”。
	ForwardingType string `json:"forwarding_type"`

	// **参数说明**：查询配置列表的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// **参数说明**：查询时每页显示的记录数。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowForwardingConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowForwardingConfigsRequest struct{}"
	}

	return strings.Join([]string{"ShowForwardingConfigsRequest", string(data)}, " ")
}
