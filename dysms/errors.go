// Package dysms Copyright 2016 The GiterLab Authors. All rights reserved.
package dysms

import (
	"encoding/json"
)

// ErrorMessage 短信服务器返回的错误信息
type ErrorMessage struct {
	HTTPCode  int     `json:"-"`
	RequestID *string `json:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty"`
	Message   *string `json:"Message,omitempty"`
}

// SetHTTPCode 设置HTTP错误码
func (e *ErrorMessage) SetHTTPCode(code int) {
	e.HTTPCode = code
}

// GetHTTPCode 获取HTTP请求的错误码
func (e *ErrorMessage) GetHTTPCode() int {
	return e.HTTPCode
}

// GetRequestID 获取请求的ID序列
func (e *ErrorMessage) GetRequestID() string {
	if e != nil && e.RequestID != nil {
		return *e.RequestID
	}
	return ""
}

// GetCode 获取请求的错误码
func (e *ErrorMessage) GetCode() string {
	if e != nil && e.Code != nil {
		return *e.Code
	}
	return ""
}

// GetMessage 获取错误信息
func (e *ErrorMessage) GetMessage() string {
	if e != nil && e.Message != nil {
		return *e.Message
	}
	return ""
}

// Error 序列化成字符串
func (e *ErrorMessage) Error() string {
	body, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(body)
}
