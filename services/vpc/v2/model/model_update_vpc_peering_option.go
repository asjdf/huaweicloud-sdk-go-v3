/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// 更新peering对象
type UpdateVpcPeeringOption struct {
	// 功能说明：对等连接名称 取值范围：支持1~64个字符
	Name string `json:"name,omitempty"`
	// 功能说明：对等连接描述 取值范围：0-255个字符，支持数字、字母、中文字符
	Description string `json:"description,omitempty"`
}
