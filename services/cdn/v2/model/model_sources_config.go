package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SourcesConfig 源站配置。
type SourcesConfig struct {

	// 源站类型， ipaddr：源站IP，domain：源站域名，obs_bucket：OBS桶域名。
	OriginType string `json:"origin_type"`

	// 源站IP或者域名。
	OriginAddr string `json:"origin_addr"`

	// 源站优先级，70：主，30：备。
	Priority int32 `json:"priority"`

	// 权重，取值范围1-100。
	Weight *int32 `json:"weight,omitempty"`

	// 是否开启Obs静态网站托管，源站类型为obs_bucket时传递，off：关闭，on：开启。
	ObsWebHostingStatus *string `json:"obs_web_hosting_status,omitempty"`

	// HTTP端口，默认80,端口取值取值范围1-65535。
	HttpPort *int32 `json:"http_port,omitempty"`

	// HTTPS端口，默认443,端口取值取值范围1-65535。
	HttpsPort *int32 `json:"https_port,omitempty"`

	// 回源HOST，默认加速域名。
	HostName *string `json:"host_name,omitempty"`

	// obs桶源站类型 “private” 私有桶 “public” 公有桶。
	ObsBucketType *string `json:"obs_bucket_type,omitempty"`
}

func (o SourcesConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourcesConfig struct{}"
	}

	return strings.Join([]string{"SourcesConfig", string(data)}, " ")
}
