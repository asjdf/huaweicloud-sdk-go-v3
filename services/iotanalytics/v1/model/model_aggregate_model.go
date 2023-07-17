package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AggregateModel 聚合计算
type AggregateModel struct {

	// 输入参数，最多支持10个
	Inputs []InputModel `json:"inputs"`

	Expression *Expression `json:"expression"`

	Schedule *DtSchedule `json:"schedule"`

	// 输出属性名(分析任务单输出场景，配合expression的formula使用)
	OutputProperty *string `json:"output_property,omitempty"`

	// 输出属性，最多支持10个
	Outputs *[]OutputWithModel `json:"outputs,omitempty"`
}

func (o AggregateModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregateModel struct{}"
	}

	return strings.Join([]string{"AggregateModel", string(data)}, " ")
}
