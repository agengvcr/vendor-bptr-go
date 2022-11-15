package response

import "vendor.bptr.co.id/api-auth/model"

type AnnoucementResponse struct {
	Id    int    `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewGetAnnoucementFromSetting(setting model.Setting) AnnoucementResponse {
	return AnnoucementResponse{
		Id:    setting.Id,
		Key:   setting.Key,
		Value: setting.Value,
	}
}
