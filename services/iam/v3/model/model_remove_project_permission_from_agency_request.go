/*
    * IAM
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// Request Object
type RemoveProjectPermissionFromAgencyRequest struct {
	ProjectId string `json:"project_id"`
	AgencyId string `json:"agency_id"`
	RoleId string `json:"role_id"`
}
