/*
    * IAM
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// 
type AgencyPolicyRoleOption struct {
	// 自定义策略展示名。
	DisplayName string `json:"display_name"`
	// 自定义策略的显示模式。   > - AX表示在domain层显示。   > - XA表示在project层显示。   > - 自定义策略的显示模式只能为AX或者XA，不能在domain层和project层都显示（AA），或者在domain层和project层都不显示（XX）。
	Type string `json:"type"`
	// 自定义策略的描述信息。
	Description string `json:"description"`
	// 自定义策略的中文描述信息。
	DescriptionCn string `json:"description_cn,omitempty"`
	Policy *AgencyPolicy `json:"policy"`
}
