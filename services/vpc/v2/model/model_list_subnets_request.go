/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Request Object
type ListSubnetsRequest struct {
	Limit int32 `json:"limit,omitempty"`
	Marker string `json:"marker,omitempty"`
	VpcId string `json:"vpc_id,omitempty"`
}
