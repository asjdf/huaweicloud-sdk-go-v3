package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GaussDBforOpenDatabaseForCreation 数据库信息。
type GaussDBforOpenDatabaseForCreation struct {

	// 数据库名称。  数据库名称长度可在1～63个字符之间，由字母、数字、或下划线组成，不能包含其他特殊字符，不能以“pg”和数字开头，且不能和GaussDB 模板库重名。 GaussDB 模板库包括postgres， template0 ，template1。
	Name string `json:"name"`

	// 数据库字符集。默认C。
	CharacterSet *string `json:"character_set,omitempty"`

	// 数据库所属用户，缺省时默认是root，不能和系统用户重名，且必须是已存在的用户。  系统用户包括“rdsAdmin”,“ rdsMetric”, “rdsBackup”, “rdsRepl”。
	Owner *string `json:"owner,omitempty"`

	// 数据库模板名称，仅为template0。
	Template *string `json:"template,omitempty"`

	// 数据库排序集。默认默认C。  - 须知： 不同的排序规则下，相同字符串的比较其结果可能是不同的。 例如，在en_US.utf8下， select 'a'>'A';执行结果为false，但在'C'下，select 'a'>'A';结果为true。如果数据库从“O”迁移到GaussDB ，数据库排序集需使用'C'才能得到一致的预期。支持的排序规则可以查询系统表 pg_collation。
	LcCollate *string `json:"lc_collate,omitempty"`

	// 数据库分类集。默认C。
	LcCtype *string `json:"lc_ctype,omitempty"`
}

func (o GaussDBforOpenDatabaseForCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GaussDBforOpenDatabaseForCreation struct{}"
	}

	return strings.Join([]string{"GaussDBforOpenDatabaseForCreation", string(data)}, " ")
}
