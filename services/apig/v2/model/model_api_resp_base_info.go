package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ApiRespBaseInfo struct {

	// API名称。  长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、下划线组成，且只能以英文或中文开头。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name"`

	// API类型 - 1：公有API - 2：私有API
	Type ApiRespBaseInfoType `json:"type"`

	// API的版本
	Version *string `json:"version,omitempty"`

	// API的请求协议 - HTTP - HTTPS - BOTH：同时支持HTTP和HTTPS
	ReqProtocol ApiRespBaseInfoReqProtocol `json:"req_protocol"`

	// API的请求方式
	ReqMethod ApiRespBaseInfoReqMethod `json:"req_method"`

	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512，且满足URI规范。  > 需要服从URI规范。
	ReqUri string `json:"req_uri"`

	// API的认证方式 - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证
	AuthType ApiRespBaseInfoAuthType `json:"auth_type"`

	AuthOpt *AuthOpt `json:"auth_opt,omitempty"`

	// 是否支持跨域 - TRUE：支持 - FALSE：不支持
	Cors *bool `json:"cors,omitempty"`

	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配） 默认：NORMAL
	MatchMode *ApiRespBaseInfoMatchMode `json:"match_mode,omitempty"`

	// 后端类型 - HTTP：web后端 - FUNCTION：函数工作流 - MOCK：模拟的后端
	BackendType ApiRespBaseInfoBackendType `json:"backend_type"`

	// API描述。字符长度不超过255 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// API所属的分组编号
	GroupId string `json:"group_id"`

	// API请求体描述，可以是请求体示例、媒体类型、参数等信息。字符长度不超过20480 > 中文字符必须为UTF-8或者unicode编码。
	BodyRemark *string `json:"body_remark,omitempty"`

	// 正常响应示例，描述API的正常返回信息。字符长度不超过20480 > 中文字符必须为UTF-8或者unicode编码。
	ResultNormalSample *string `json:"result_normal_sample,omitempty"`

	// 失败返回示例，描述API的异常返回信息。字符长度不超过20480 > 中文字符必须为UTF-8或者unicode编码。
	ResultFailureSample *string `json:"result_failure_sample,omitempty"`

	// 前端自定义认证对象的ID
	AuthorizerId *string `json:"authorizer_id,omitempty"`

	// 标签。  支持英文，数字，中文，特殊符号（-*#%.:_），且只能以中文或英文开头。  默认支持10个标签，如需扩大配额请联系技术工程师修改API_TAG_NUM_LIMIT配置。
	Tags *[]string `json:"tags,omitempty"`

	// 分组自定义响应ID
	ResponseId *string `json:"response_id,omitempty"`

	// 集成应用ID  暂不支持
	RomaAppId *string `json:"roma_app_id,omitempty"`

	// API绑定的自定义域名  暂不支持
	DomainName *string `json:"domain_name,omitempty"`

	// 标签  待废弃，优先使用tags字段
	Tag *string `json:"tag,omitempty"`

	// 请求内容格式类型：  application/json application/xml multipart/form-date text/plain  暂不支持
	ContentType *ApiRespBaseInfoContentType `json:"content_type,omitempty"`

	// API编号
	Id *string `json:"id,omitempty"`

	// API状态   - 1： 有效
	Status *int32 `json:"status,omitempty"`

	// 是否需要编排
	ArrangeNecessary *int32 `json:"arrange_necessary,omitempty"`

	// API注册时间
	RegisterTime *sdktime.SdkTime `json:"register_time,omitempty"`

	// API修改时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// API所属分组的名称
	GroupName *string `json:"group_name,omitempty"`

	// API所属分组的版本  默认V1，其他版本暂不支持
	GroupVersion *string `json:"group_version,omitempty"`

	// 发布的环境编号  存在多个发布记录时，环境编号之间用|隔开
	RunEnvId *string `json:"run_env_id,omitempty"`

	// 发布的环境名称  存在多个发布记录时，环境名称之间用|隔开
	RunEnvName *string `json:"run_env_name,omitempty"`

	// 发布记录编号  存在多个发布记录时，发布记录编号之间用|隔开
	PublishId *string `json:"publish_id,omitempty"`

	// 发布时间  存在多个发布记录时，发布时间之间用|隔开
	PublishTime *string `json:"publish_time,omitempty"`

	// API归属的集成应用名称  暂不支持
	RomaAppName *string `json:"roma_app_name,omitempty"`

	// 当API的后端为自定义后端时，对应的自定义后端API编号  暂不支持
	LdApiId *string `json:"ld_api_id,omitempty"`

	BackendApi *BackendApi `json:"backend_api,omitempty"`

	ApiGroupInfo *ApiGroupCommonInfo `json:"api_group_info,omitempty"`
}

func (o ApiRespBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiRespBaseInfo struct{}"
	}

	return strings.Join([]string{"ApiRespBaseInfo", string(data)}, " ")
}

type ApiRespBaseInfoType struct {
	value int32
}

type ApiRespBaseInfoTypeEnum struct {
	E_1 ApiRespBaseInfoType
	E_2 ApiRespBaseInfoType
}

func GetApiRespBaseInfoTypeEnum() ApiRespBaseInfoTypeEnum {
	return ApiRespBaseInfoTypeEnum{
		E_1: ApiRespBaseInfoType{
			value: 1,
		}, E_2: ApiRespBaseInfoType{
			value: 2,
		},
	}
}

func (c ApiRespBaseInfoType) Value() int32 {
	return c.value
}

func (c ApiRespBaseInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiRespBaseInfoType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ApiRespBaseInfoReqProtocol struct {
	value string
}

type ApiRespBaseInfoReqProtocolEnum struct {
	HTTP  ApiRespBaseInfoReqProtocol
	HTTPS ApiRespBaseInfoReqProtocol
	BOTH  ApiRespBaseInfoReqProtocol
}

func GetApiRespBaseInfoReqProtocolEnum() ApiRespBaseInfoReqProtocolEnum {
	return ApiRespBaseInfoReqProtocolEnum{
		HTTP: ApiRespBaseInfoReqProtocol{
			value: "HTTP",
		},
		HTTPS: ApiRespBaseInfoReqProtocol{
			value: "HTTPS",
		},
		BOTH: ApiRespBaseInfoReqProtocol{
			value: "BOTH",
		},
	}
}

func (c ApiRespBaseInfoReqProtocol) Value() string {
	return c.value
}

func (c ApiRespBaseInfoReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiRespBaseInfoReqProtocol) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ApiRespBaseInfoReqMethod struct {
	value string
}

type ApiRespBaseInfoReqMethodEnum struct {
	GET     ApiRespBaseInfoReqMethod
	POST    ApiRespBaseInfoReqMethod
	PUT     ApiRespBaseInfoReqMethod
	DELETE  ApiRespBaseInfoReqMethod
	HEAD    ApiRespBaseInfoReqMethod
	PATCH   ApiRespBaseInfoReqMethod
	OPTIONS ApiRespBaseInfoReqMethod
	ANY     ApiRespBaseInfoReqMethod
}

func GetApiRespBaseInfoReqMethodEnum() ApiRespBaseInfoReqMethodEnum {
	return ApiRespBaseInfoReqMethodEnum{
		GET: ApiRespBaseInfoReqMethod{
			value: "GET",
		},
		POST: ApiRespBaseInfoReqMethod{
			value: "POST",
		},
		PUT: ApiRespBaseInfoReqMethod{
			value: "PUT",
		},
		DELETE: ApiRespBaseInfoReqMethod{
			value: "DELETE",
		},
		HEAD: ApiRespBaseInfoReqMethod{
			value: "HEAD",
		},
		PATCH: ApiRespBaseInfoReqMethod{
			value: "PATCH",
		},
		OPTIONS: ApiRespBaseInfoReqMethod{
			value: "OPTIONS",
		},
		ANY: ApiRespBaseInfoReqMethod{
			value: "ANY",
		},
	}
}

func (c ApiRespBaseInfoReqMethod) Value() string {
	return c.value
}

func (c ApiRespBaseInfoReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiRespBaseInfoReqMethod) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ApiRespBaseInfoAuthType struct {
	value string
}

type ApiRespBaseInfoAuthTypeEnum struct {
	NONE       ApiRespBaseInfoAuthType
	APP        ApiRespBaseInfoAuthType
	IAM        ApiRespBaseInfoAuthType
	AUTHORIZER ApiRespBaseInfoAuthType
}

func GetApiRespBaseInfoAuthTypeEnum() ApiRespBaseInfoAuthTypeEnum {
	return ApiRespBaseInfoAuthTypeEnum{
		NONE: ApiRespBaseInfoAuthType{
			value: "NONE",
		},
		APP: ApiRespBaseInfoAuthType{
			value: "APP",
		},
		IAM: ApiRespBaseInfoAuthType{
			value: "IAM",
		},
		AUTHORIZER: ApiRespBaseInfoAuthType{
			value: "AUTHORIZER",
		},
	}
}

func (c ApiRespBaseInfoAuthType) Value() string {
	return c.value
}

func (c ApiRespBaseInfoAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiRespBaseInfoAuthType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ApiRespBaseInfoMatchMode struct {
	value string
}

type ApiRespBaseInfoMatchModeEnum struct {
	SWA    ApiRespBaseInfoMatchMode
	NORMAL ApiRespBaseInfoMatchMode
}

func GetApiRespBaseInfoMatchModeEnum() ApiRespBaseInfoMatchModeEnum {
	return ApiRespBaseInfoMatchModeEnum{
		SWA: ApiRespBaseInfoMatchMode{
			value: "SWA",
		},
		NORMAL: ApiRespBaseInfoMatchMode{
			value: "NORMAL",
		},
	}
}

func (c ApiRespBaseInfoMatchMode) Value() string {
	return c.value
}

func (c ApiRespBaseInfoMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiRespBaseInfoMatchMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ApiRespBaseInfoBackendType struct {
	value string
}

type ApiRespBaseInfoBackendTypeEnum struct {
	HTTP     ApiRespBaseInfoBackendType
	FUNCTION ApiRespBaseInfoBackendType
	MOCK     ApiRespBaseInfoBackendType
}

func GetApiRespBaseInfoBackendTypeEnum() ApiRespBaseInfoBackendTypeEnum {
	return ApiRespBaseInfoBackendTypeEnum{
		HTTP: ApiRespBaseInfoBackendType{
			value: "HTTP",
		},
		FUNCTION: ApiRespBaseInfoBackendType{
			value: "FUNCTION",
		},
		MOCK: ApiRespBaseInfoBackendType{
			value: "MOCK",
		},
	}
}

func (c ApiRespBaseInfoBackendType) Value() string {
	return c.value
}

func (c ApiRespBaseInfoBackendType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiRespBaseInfoBackendType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ApiRespBaseInfoContentType struct {
	value string
}

type ApiRespBaseInfoContentTypeEnum struct {
	APPLICATION_JSON    ApiRespBaseInfoContentType
	APPLICATION_XML     ApiRespBaseInfoContentType
	MULTIPART_FORM_DATE ApiRespBaseInfoContentType
	TEXT_PLAIN          ApiRespBaseInfoContentType
}

func GetApiRespBaseInfoContentTypeEnum() ApiRespBaseInfoContentTypeEnum {
	return ApiRespBaseInfoContentTypeEnum{
		APPLICATION_JSON: ApiRespBaseInfoContentType{
			value: "application/json",
		},
		APPLICATION_XML: ApiRespBaseInfoContentType{
			value: "application/xml",
		},
		MULTIPART_FORM_DATE: ApiRespBaseInfoContentType{
			value: "multipart/form-date",
		},
		TEXT_PLAIN: ApiRespBaseInfoContentType{
			value: "text/plain",
		},
	}
}

func (c ApiRespBaseInfoContentType) Value() string {
	return c.value
}

func (c ApiRespBaseInfoContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiRespBaseInfoContentType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
