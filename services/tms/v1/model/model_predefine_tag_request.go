/*
    * TMS
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

//  标签详情。
type PredefineTagRequest struct {
	// 键。最大长度36个字符。字符集：A-Z，a-z ， 0-9，‘-’，‘_’，UNICODE字符（\\u4E00-\\u9FFF）。
	Key string `json:"key"`
	// 值。每个值最大长度43个字符，可以为空字符串。字符集：AZ，a-z ， 0-9，‘.’，‘-’，‘_’，UNICODE字符（\\u4E00-\\u9FFF）。
	Value string `json:"value"`
}
