package models

type APIResponse struct {
	Code     int         `json:"code"`
	Data     interface{} `json:"data"`
	HttpCode int         `json:"httpCode"`
}

type QueryParam map[string]interface{}

type ErrorDetails struct {
	ErrorCode int
	ErrorMsg  string
}
