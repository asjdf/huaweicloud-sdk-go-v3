package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateDedicatedHostTagsResponse Response Object
type BatchCreateDedicatedHostTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateDedicatedHostTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDedicatedHostTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateDedicatedHostTagsResponse", string(data)}, " ")
}
