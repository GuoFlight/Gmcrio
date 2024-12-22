package models

type Res struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	TraceId string      `json:"traceId,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
