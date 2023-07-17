package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpcAttachmentBody This is a auto create Body Object
type CreateVpcAttachmentBody struct {
	VpcAttachment *VpcAttachmentCreateRequest `json:"vpc_attachment"`
}

func (o CreateVpcAttachmentBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcAttachmentBody struct{}"
	}

	return strings.Join([]string{"CreateVpcAttachmentBody", string(data)}, " ")
}
