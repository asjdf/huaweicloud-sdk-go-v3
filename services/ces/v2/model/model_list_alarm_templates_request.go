package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAlarmTemplatesRequest Request Object
type ListAlarmTemplatesRequest struct {

	// 分页查询时查询的起始位置，表示从第几条数据开始，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 查询结果条数的限制值，取值范围为[1,100]，默认值为100
	Limit *int32 `json:"limit,omitempty"`

	// 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”
	Namespace *string `json:"namespace,omitempty"`

	// 资源维度，必须以字母开头，多维度用\",\"分割，只能包含0-9/a-z/A-Z/_/-，每个维度的最大长度为32
	DimName *string `json:"dim_name,omitempty"`

	// 模板类型(system代表默认自定义模板，custom代表系统模板),不传自定义和系统均需返回
	TemplateType *ListAlarmTemplatesRequestTemplateType `json:"template_type,omitempty"`

	// 告警模板的名称，以字母或汉字开头，可包含字母、数字、汉字、_、-，长度范围[1,128]，支持模糊匹配
	TemplateName *string `json:"template_name,omitempty"`
}

func (o ListAlarmTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmTemplatesRequest", string(data)}, " ")
}

type ListAlarmTemplatesRequestTemplateType struct {
	value string
}

type ListAlarmTemplatesRequestTemplateTypeEnum struct {
	SYSTEM ListAlarmTemplatesRequestTemplateType
	CUSTOM ListAlarmTemplatesRequestTemplateType
}

func GetListAlarmTemplatesRequestTemplateTypeEnum() ListAlarmTemplatesRequestTemplateTypeEnum {
	return ListAlarmTemplatesRequestTemplateTypeEnum{
		SYSTEM: ListAlarmTemplatesRequestTemplateType{
			value: "system",
		},
		CUSTOM: ListAlarmTemplatesRequestTemplateType{
			value: "custom",
		},
	}
}

func (c ListAlarmTemplatesRequestTemplateType) Value() string {
	return c.value
}

func (c ListAlarmTemplatesRequestTemplateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmTemplatesRequestTemplateType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
