/*
    * cts
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// Response Object
type ListTrackersResponse struct {
	// 本次查询追踪器列表返回的追踪器数组。
	Trackers []TrackerResponseBody `json:"trackers,omitempty"`
}
