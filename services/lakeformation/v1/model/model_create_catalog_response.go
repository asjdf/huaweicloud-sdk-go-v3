package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCatalogResponse Response Object
type CreateCatalogResponse struct {

	// catalog名字
	CatalogName *string `json:"catalog_name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 路径地址
	Location *string `json:"location,omitempty"`

	// database路径列表
	DatabaseLocationList *[]string `json:"database_location_list,omitempty"`
	HttpStatusCode       int       `json:"-"`
}

func (o CreateCatalogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCatalogResponse struct{}"
	}

	return strings.Join([]string{"CreateCatalogResponse", string(data)}, " ")
}
