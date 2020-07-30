/*
    * IAM
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// 
type RegionLocales struct {
	// 区域的中文名称。
	ZhCn string `json:"zh-cn"`
	// 区域的英文名称。
	EnUs string `json:"en-us"`
	// 区域的葡萄牙语名称。
	PtBr string `json:"pt-br,omitempty"`
	// 区域的美国西班牙语名称。
	EsUs string `json:"es-us,omitempty"`
	// 区域的西班牙语名称。
	EsEs string `json:"es-es,omitempty"`
}
