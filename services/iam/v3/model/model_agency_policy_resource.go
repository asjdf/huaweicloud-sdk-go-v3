/*
    * IAM
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// 
type AgencyPolicyResource struct {
	// 委托资源的URI，长度不超过128。格式为：/iam/agencies/委托ID。例：   ``` \"uri\": [\"/iam/agencies/07805acaba800fdd4fbdc00b8f888c7c\"] ```
	Uri []string `json:"uri"`
}
