package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConsumerListOrDetailsResponse Response Object
type ShowConsumerListOrDetailsResponse struct {

	// Topic列表（当查询topic消费“列表”时才显示此参数）。
	Topics *[]string `json:"topics,omitempty"`

	// Topic总数（当查询topic消费“列表”时才显示此参数）。
	Total *int32 `json:"total,omitempty"`

	// Topic关联代理（当查询topic消费“详情”才显示此参数）。
	Brokers        *[]Brokers `json:"brokers,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowConsumerListOrDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsumerListOrDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowConsumerListOrDetailsResponse", string(data)}, " ")
}
