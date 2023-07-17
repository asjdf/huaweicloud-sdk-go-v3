package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUserResponse Response Object
type DeleteUserResponse struct {

	// DDM实例帐号名称。
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteUserResponse", string(data)}, " ")
}
