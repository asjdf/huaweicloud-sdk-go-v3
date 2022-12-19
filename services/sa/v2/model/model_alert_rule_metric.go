package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AlertRuleMetric struct {

	// category. GROUP_COUNT
	Category AlertRuleMetricCategory `json:"category"`

	// metric
	Metric map[string]float32 `json:"metric"`
}

func (o AlertRuleMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertRuleMetric struct{}"
	}

	return strings.Join([]string{"AlertRuleMetric", string(data)}, " ")
}

type AlertRuleMetricCategory struct {
	value string
}

type AlertRuleMetricCategoryEnum struct {
	GROUP_COUNT AlertRuleMetricCategory
}

func GetAlertRuleMetricCategoryEnum() AlertRuleMetricCategoryEnum {
	return AlertRuleMetricCategoryEnum{
		GROUP_COUNT: AlertRuleMetricCategory{
			value: "GROUP_COUNT",
		},
	}
}

func (c AlertRuleMetricCategory) Value() string {
	return c.value
}

func (c AlertRuleMetricCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertRuleMetricCategory) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
