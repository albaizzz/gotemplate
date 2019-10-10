package svc

//HTTPResponse base response data
type HTTPResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

//LegacyHTTPResponse structure
type LegacyHTTPResponse struct {
	APIStatusCode int         `json:"api_status_code"`
	APIStatus     string      `json:"api_status,omitempty"`
	Data          interface{} `json:"data,omitempty"`
}

//httpHandlerFunc abstraction for http handler
// type httpHandlerFunc func(request *http.Request, ui *context.UIContext, scenario ucase.ScenarioFunc) (interface{}, string, error)

//Server contract
type Server interface {
	Run() error
	Done()
}
