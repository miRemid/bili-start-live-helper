package model

type Common struct {
	Code    int         `json:"code,omitempty"`
	Status  bool        `json:"status,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Ts      int         `json:"ts,omitempty"`
	Message string      `json:"message,omitempty"`
}
