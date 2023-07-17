package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlertResponse Response Object
type DeleteAlertResponse struct {

	// Id value
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	Data *DataResponse `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAlertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlertResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlertResponse", string(data)}, " ")
}
