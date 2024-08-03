package model

type AuthApp struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type AuthResp struct {
	Code              int    `json:"code"`
	Expire            int    `json:"expire"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
}
