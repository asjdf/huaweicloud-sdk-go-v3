package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndpointInfoDetailsRequest Request Object
type ListEndpointInfoDetailsRequest struct {

	// 发送的实体的MIME类型。推荐用户默认使用application/json， 如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。
	ContentType *string `json:"Content-Type,omitempty"`

	// 终端节点的ID。
	VpcEndpointId string `json:"vpc_endpoint_id"`
}

func (o ListEndpointInfoDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointInfoDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListEndpointInfoDetailsRequest", string(data)}, " ")
}
