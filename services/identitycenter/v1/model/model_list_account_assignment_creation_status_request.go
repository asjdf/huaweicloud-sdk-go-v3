package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAccountAssignmentCreationStatusRequest Request Object
type ListAccountAssignmentCreationStatusRequest struct {
	InstanceId string `json:"instance_id"`

	// Filters he operation status list based on the passed attribute value.
	Status *ListAccountAssignmentCreationStatusRequestStatus `json:"status,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`
}

func (o ListAccountAssignmentCreationStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccountAssignmentCreationStatusRequest struct{}"
	}

	return strings.Join([]string{"ListAccountAssignmentCreationStatusRequest", string(data)}, " ")
}

type ListAccountAssignmentCreationStatusRequestStatus struct {
	value string
}

type ListAccountAssignmentCreationStatusRequestStatusEnum struct {
	IN_PROGRESS ListAccountAssignmentCreationStatusRequestStatus
	SUCCEEDED   ListAccountAssignmentCreationStatusRequestStatus
	FAILED      ListAccountAssignmentCreationStatusRequestStatus
}

func GetListAccountAssignmentCreationStatusRequestStatusEnum() ListAccountAssignmentCreationStatusRequestStatusEnum {
	return ListAccountAssignmentCreationStatusRequestStatusEnum{
		IN_PROGRESS: ListAccountAssignmentCreationStatusRequestStatus{
			value: "IN_PROGRESS",
		},
		SUCCEEDED: ListAccountAssignmentCreationStatusRequestStatus{
			value: "SUCCEEDED",
		},
		FAILED: ListAccountAssignmentCreationStatusRequestStatus{
			value: "FAILED",
		},
	}
}

func (c ListAccountAssignmentCreationStatusRequestStatus) Value() string {
	return c.value
}

func (c ListAccountAssignmentCreationStatusRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAccountAssignmentCreationStatusRequestStatus) UnmarshalJSON(b []byte) error {
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
