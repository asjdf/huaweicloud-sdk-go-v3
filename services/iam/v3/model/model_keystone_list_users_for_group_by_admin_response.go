/*
    * IAM
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// Response Object
type KeystoneListUsersForGroupByAdminResponse struct {
	Links *Links `json:"links,omitempty"`
	// IAM用户信息列表。
	Users []KeystoneUserResult `json:"users,omitempty"`
}
