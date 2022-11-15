package response

import (
	"vendor.bptr.co.id/api-auth/model"
)

type GetAppVersionResponse struct {
	Version   string `json:"version"`
	ApkUrl    string `json:"url"`
	Changelog string `json:"changelog"`
}

// GetAppVersionResponse
func NewGetAppVersionFromSetting(appVersion model.AppVersion) GetAppVersionResponse {

	return GetAppVersionResponse{
		ApkUrl:    appVersion.ApkUrl,
		Changelog: appVersion.ChangeLog,
		Version:   appVersion.Version,
	}
}
