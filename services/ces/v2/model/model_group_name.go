package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GroupName 资源分组名称
type GroupName struct {
}

func (o GroupName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupName struct{}"
	}

	return strings.Join([]string{"GroupName", string(data)}, " ")
}
