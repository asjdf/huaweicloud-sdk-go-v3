/*
    * IAM
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// 
type CreateCredentialResult struct {
	// 创建访问密钥时间。
	CreateTime string `json:"create_time"`
	// 创建的AK。
	Access string `json:"access"`
	// 创建的SK。
	Secret string `json:"secret"`
	// 访问密钥状态。
	Status string `json:"status"`
	// IAM用户ID。
	UserId string `json:"user_id"`
	// 访问密钥描述信息。
	Description string `json:"description"`
}
